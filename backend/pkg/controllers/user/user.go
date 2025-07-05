package user

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"strings"
	"time"

	"github.com/akifkadioglu/vocapedia/pkg/cache"
	"github.com/akifkadioglu/vocapedia/pkg/database"
	"github.com/akifkadioglu/vocapedia/pkg/entities"
	"github.com/akifkadioglu/vocapedia/pkg/geolocation"
	"github.com/akifkadioglu/vocapedia/pkg/i18n"
	"github.com/akifkadioglu/vocapedia/pkg/mail"
	"github.com/akifkadioglu/vocapedia/pkg/payment"
	"github.com/akifkadioglu/vocapedia/pkg/payment/common"
	"github.com/akifkadioglu/vocapedia/pkg/search"
	"github.com/akifkadioglu/vocapedia/pkg/snowflake"
	"github.com/akifkadioglu/vocapedia/pkg/token"
	"github.com/akifkadioglu/vocapedia/pkg/utils"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/meilisearch/meilisearch-go"
	"github.com/redis/go-redis/v9"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

func GetByUsername(w http.ResponseWriter, r *http.Request) {
	db := database.Manager()

	username := r.URL.Query().Get("username")

	if username == "" {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{
			"error": "Username query parameter is required",
		})
		return
	}

	var user entities.User
	err := db.Where("LOWER(username) = ?", strings.ToLower(username)).First(&user).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			render.Status(r, http.StatusNotFound)
			render.JSON(w, r, map[string]string{
				"error": "User not found",
			})
		} else {
			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, map[string]string{
				"error": "Database error",
			})
		}
		return
	}

	render.JSON(w, r, user)
}

func SearchUsers(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	client := search.Meili()
	index := client.Index("users")

	searchRequest := &meilisearch.SearchRequest{
		Query: query,
		Limit: 10,
	}

	searchRes, err := index.Search(query, searchRequest)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{
			"error": "Search error: " + err.Error(),
		})
		return
	}

	render.JSON(w, r, map[string]any{
		"list": searchRes.Hits,
	})
}

func EditUser(w http.ResponseWriter, r *http.Request) {
	db := database.Manager()

	userID := token.User(r).UserID
	if userID == "" {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{
			"error": i18n.Localizer(r, "error.user_id_required"),
		})
		return
	}
	var user entities.User
	err := db.Where("id = ?", userID).First(&user).Error
	if err != nil {
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, map[string]string{
			"error": i18n.Localizer(r, "error.user_not_found"),
		})
		return
	}
	var updatedUser _updateUser
	if err := render.Decode(r, &updatedUser); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{
			"error": i18n.Localizer(r, "error.invalid_request_body"),
		})
		return
	}
	device, err := utils.StructToMap(updatedUser.Device)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{
			"error": i18n.Localizer(r, "error.something_went_wrong"),
		})
		return
	}

	deviceString := ""
	for _, k := range device {
		deviceString = deviceString + " | " + fmt.Sprintf("%v", k)
	}
	if err := db.Model(&user).Updates(map[string]any{
		"name":      updatedUser.Name,
		"username":  updatedUser.Username,
		"biography": updatedUser.Biography,
	}).Where("id = ?", userID).Error; err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{
			"error": i18n.Localizer(r, "error.failed_to_update_user"),
		})
		return
	}

	tmpl, err := template.ParseFiles("pkg/mail/templates/edit.user.html")
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{
			"error": i18n.Localizer(r, "error.something_went_wrong"),
		})
		return
	}
	var body bytes.Buffer

	err = tmpl.Execute(&body, _emailData{
		Header:      i18n.Localizer(r, "mail.edit.user.header"),
		Description: i18n.Localizer(r, "mail.edit.user.description"),
		Device:      deviceString,
	})
	if err != nil {
		render.Status(r, http.StatusInternalServerError)

		render.JSON(w, r, map[string]string{
			"error": i18n.Localizer(r, "error.something_went_wrong"),
		})
		return
	}
	subject := i18n.Localizer(r, "mail.edit.user.header")

	isSent, err := mail.Send(r, user.Email, subject, body.String())
	if err != nil && isSent {
		render.Status(r, http.StatusBadRequest)

		render.JSON(w, r, map[string]string{
			"error": i18n.Localizer(r, "error.something_went_wrong"),
		})
		return
	}
	jwtModel := entities.JwtModel{
		UserID:   userID,
		Username: user.Username,
		Device:   device,
	}

	tokenString, err := token.GenerateToken(jwtModel)
	if err != nil {
		render.Status(r, http.StatusBadRequest)

		render.JSON(w, r, map[string]string{
			"error": i18n.Localizer(r, "error.something_went_wrong"),
		})
		return
	}
	render.JSON(w, r, map[string]string{
		"token": tokenString,
	})
}

func Tokens(w http.ResponseWriter, r *http.Request) {
	db := database.Manager()
	var tokenEntity []entities.Token
	tokenStr := strings.Split(r.Header.Get("Authorization"), " ")[1]
	err := db.Where("token = ?", tokenStr).First(&tokenEntity).Error
	if err != nil {
		render.JSON(w, r, map[string][]entities.Token{
			"tokens": {},
		})
		return
	}
	userID := token.User(r).UserID

	err = db.Where("user_id = ?", userID).Find(&tokenEntity).Error
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{
			"error": i18n.Localizer(r, "error.tokens_fetch_failed"),
		})
		return
	}

	render.JSON(w, r, map[string][]entities.Token{
		"tokens": tokenEntity,
	})

}

func DeleteToken(w http.ResponseWriter, r *http.Request) {
	db := database.Manager()
	tokenIDStr := chi.URLParam(r, "id")

	userID := token.User(r).UserID
	var tokenEntity entities.Token
	if err := db.Where("id = ? AND user_id = ?", tokenIDStr, userID).First(&tokenEntity).Error; err != nil {
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, map[string]string{
			"error": i18n.Localizer(r, "error.token_not_found"),
		})
		return
	}

	if err := db.Delete(&tokenEntity).Error; err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{
			"error": i18n.Localizer(r, "error.something_went_wrong"),
		})
		return
	}

	render.JSON(w, r, map[string]string{
		"message": i18n.Localizer(r, "success.token_deleted"),
	})
}

func Check(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, map[string]string{})
}
func UpdateVocaToken(w http.ResponseWriter, r *http.Request) {
	db := database.Manager()
	userID := token.User(r).UserID
	vocatoken, err := utils.GenerateVocaToken(10)
	if err != nil {
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, map[string]string{
			"error": i18n.Localizer(r, "error.something_went_wrong"),
		})
		return
	}
	err = db.Model(&entities.User{}).
		Where("id = ?", userID).
		Update("vocatoken", vocatoken).Error
	if err != nil {
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, map[string]string{
			"error": i18n.Localizer(r, "error.something_went_wrong"),
		})
		return
	}
	render.JSON(w, r, map[string]string{
		"vocatoken": vocatoken,
	})
}

func GetVocaToken(w http.ResponseWriter, r *http.Request) {
	db := database.Manager()
	userID := token.User(r).UserID
	var user entities.User
	err := db.Where("id = ?", userID).First(&user).Error
	if err != nil {
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, map[string]string{
			"error": i18n.Localizer(r, "error.something_went_wrong"),
		})
		return
	}
	render.JSON(w, r, map[string]any{
		"vocatoken": user.Vocatoken,
	})
}

func DailyStreak(w http.ResponseWriter, r *http.Request) {
	rdb := cache.Redis()
	userID := token.User(r).UserID
	db := database.Manager()

	data, err := rdb.Get(r.Context(), fmt.Sprintf("streak:%s", userID)).Result()
	var streak _streak
	if err == redis.Nil {
		streak = _streak{Count: 0, LastDate: "", Rewarded: false}
	} else if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{
			"error": i18n.Localizer(r, "error.something_went_wrong"),
		})
		return
	} else {
		json.Unmarshal([]byte(data), &streak)
	}
	today := time.Now().Format("2006-01-02")
	yesterday := time.Now().AddDate(0, 0, -1).Format("2006-01-02")

	switch streak.LastDate {
	case today:
	case yesterday:
		streak.Count++
		if streak.Count > 7 {
			streak.Count = 1
			streak.Rewarded = false
		}

		streak.LastDate = today
	default:
		streak.Count = 1
		streak.LastDate = today
		streak.Rewarded = false
	}

	if streak.Count == 7 && !streak.Rewarded {
		streak.Rewarded = true
		if err := db.Model(&entities.User{}).Where("id = ?", userID).Update("vocatoken_val", gorm.Expr("vocatoken_val + ?", 20)).Error; err != nil {
			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, map[string]string{
				"error": i18n.Localizer(r, "error.something_went_wrong"),
			})
			return
		}
	}

	streakJSON, _ := json.Marshal(streak)
	rdb.Set(r.Context(), fmt.Sprintf("streak:%s", userID), string(streakJSON), 0)

	render.JSON(w, r, map[string]any{
		"streak": streak,
	})
}

// InitiateTokenPurchase handles the token purchase request
func InitiateTokenPurchase(w http.ResponseWriter, r *http.Request) {
	userID := token.User(r).UserID
	if userID == "" {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{
			"error": i18n.Localizer(r, "error.user_id_required"),
		})
		return
	}

	// Get user info for payment
	db := database.Manager()
	var user entities.User
	if err := db.Where("id = ?", userID).First(&user).Error; err != nil {
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, map[string]string{
			"error": "User not found",
		})
		return
	}

	var req _tokenPurchaseRequest
	if err := render.DecodeJSON(r.Body, &req); err != nil {
		fmt.Println("Error decoding JSON:", err)
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{
			"error": i18n.Localizer(r, "error.invalid_request"),
		})
		return
	}

	// Validate token count
	if req.Tokens < 1 {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{
			"error": "Token count must be greater than 1",
		})
		return
	}

	// Create a transaction record
	transactionID := fmt.Sprintf("%d", snowflake.GenerateID())

	// Smart payment provider detection with multiple factors
	userIP := geolocation.GetRealIP(r)
	geoService := geolocation.NewGeoLocationService()

	// Get multiple detection factors
	acceptLanguage := r.Header.Get("Accept-Language")
	timezone := r.Header.Get("X-Timezone") // If frontend sends this

	// Get country code from IP with fallback
	countryCode, err := geoService.GetCountryFromIP(r.Context(), userIP)
	if err != nil {
		// Log error but continue with fallback
		fmt.Printf("Geolocation error for IP %s: %v - using fallback\n", userIP, err)
	}

	// Get smart payment provider configuration
	providerConfig := geolocation.GetPaymentProviderForCountry(countryCode, acceptLanguage, timezone)

	// Log detection details for debugging
	fmt.Printf("🔍 Payment Detection - IP: %s, Country: %s, Lang: %s, Provider: %s (%s), Confidence: %d%%\n",
		userIP, countryCode, acceptLanguage, providerConfig.Provider, providerConfig.Reason, providerConfig.Confidence)

	// Use detected provider
	providerName := providerConfig.Provider
	currency := providerConfig.Currency

	// Calculate price and discount
	price := calculateTokenPrice(req.Tokens, currency)
	discount := getDiscountPercentage(req.Tokens)

	// Get payment manager and create payment
	paymentManager := payment.GetManager()

	// Prepare payment request
	paymentRequest := common.PaymentRequest{
		Amount:        price,
		Currency:      currency,
		Description:   fmt.Sprintf("%d VocaPedia Tokens", req.Tokens),
		UserID:        userID,
		UserEmail:     user.Email,
		UserName:      user.Name,
		TransactionID: transactionID,
		CallbackURL:   "https://vocapedia.space/payment/callback",
		WebhookURL:    "https://api.vocapedia.space/v1/public/webhooks/payment",
		Metadata: map[string]string{
			"tokens":   fmt.Sprintf("%d", req.Tokens),
			"discount": fmt.Sprintf("%d", discount),
		},
	}

	paymentResponse, err := paymentManager.CreatePayment(r.Context(), providerName, paymentRequest)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{
			"error": "Failed to create payment: " + err.Error(),
		})
		return
	}

	// Store transaction details in cache for later confirmation
	transactionData := map[string]interface{}{
		"user_id":        userID,
		"tokens":         req.Tokens,
		"price":          price,
		"currency":       currency,
		"discount":       discount,
		"provider":       providerName,
		"provider_id":    paymentResponse.ProviderID,
		"status":         "pending",
		"created_at":     time.Now(),
		"transaction_id": transactionID,
	}

	transactionJSON, _ := json.Marshal(transactionData)
	cache.Redis().Set(r.Context(), fmt.Sprintf("token_purchase:%s", transactionID), string(transactionJSON), 24*time.Hour)

	response := _tokenPurchaseResponse{
		Success:       true,
		Tokens:        req.Tokens,
		Price:         price,
		Currency:      currency,
		Discount:      discount,
		PaymentURL:    paymentResponse.PaymentURL,
		TransactionID: transactionID,
		Provider:      providerName,
		Country:       countryCode,
	}

	render.JSON(w, r, response)
}

// ConfirmTokenPurchase handles the confirmation of token purchase (webhook from Gumroad)
func ConfirmTokenPurchase(w http.ResponseWriter, r *http.Request) {
	db := database.Manager()

	transactionID := chi.URLParam(r, "transaction_id")
	if transactionID == "" {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{
			"error": "Transaction ID is required",
		})
		return
	}

	// Get transaction details from cache
	data, err := cache.Redis().Get(r.Context(), fmt.Sprintf("token_purchase:%s", transactionID)).Result()
	if err != nil {
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, map[string]string{
			"error": "Transaction not found",
		})
		return
	}

	var transactionData map[string]interface{}
	if err := json.Unmarshal([]byte(data), &transactionData); err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{
			"error": "Invalid transaction data",
		})
		return
	}

	userID := transactionData["user_id"].(string)
	tokens := int(transactionData["tokens"].(float64))

	// Add tokens to user account
	if err := db.Model(&entities.User{}).Where("id = ?", userID).Update("vocatoken_val", gorm.Expr("vocatoken_val + ?", tokens)).Error; err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{
			"error": "Failed to update user tokens",
		})
		return
	}

	// Update transaction status
	transactionData["status"] = "completed"
	transactionData["completed_at"] = time.Now()
	transactionJSON, _ := json.Marshal(transactionData)
	cache.Redis().Set(r.Context(), fmt.Sprintf("token_purchase:%s", transactionID), string(transactionJSON), 7*24*time.Hour)

	render.JSON(w, r, map[string]interface{}{
		"success": true,
		"tokens":  tokens,
		"message": "Tokens successfully added to your account",
	})
}

// GetTokenPurchaseStatus checks the status of a token purchase
func GetTokenPurchaseStatus(w http.ResponseWriter, r *http.Request) {
	transactionID := chi.URLParam(r, "transaction_id")
	if transactionID == "" {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{
			"error": "Transaction ID is required",
		})
		return
	}

	// Get transaction details from cache
	data, err := cache.Redis().Get(r.Context(), fmt.Sprintf("token_purchase:%s", transactionID)).Result()
	if err != nil {
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, map[string]string{
			"error": "Transaction not found",
		})
		return
	}

	var transactionData map[string]interface{}
	if err := json.Unmarshal([]byte(data), &transactionData); err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{
			"error": "Invalid transaction data",
		})
		return
	}

	render.JSON(w, r, transactionData)
}

// PaymentWebhook handles webhook from payment providers
func PaymentWebhook(w http.ResponseWriter, r *http.Request) {
	// Get provider from URL parameter or header
	providerName := r.URL.Query().Get("provider")
	if providerName == "" {
		providerName = r.Header.Get("X-Payment-Provider")
	}
	if providerName == "" {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{
			"error": "Payment provider not specified",
		})
		return
	}

	// Parse webhook data
	var webhookData map[string]interface{}
	if err := render.DecodeJSON(r.Body, &webhookData); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{
			"error": "Invalid webhook data",
		})
		return
	}

	// Get payment manager and verify payment
	paymentManager := payment.GetManager()
	verification, err := paymentManager.VerifyPayment(r.Context(), providerName, webhookData)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{
			"error": "Payment verification failed: " + err.Error(),
		})
		return
	}

	if !verification.IsValid {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{
			"error": "Invalid payment",
		})
		return
	}

	// Process the payment confirmation
	db := database.Manager()

	// Get transaction details from cache
	data, err := cache.Redis().Get(r.Context(), fmt.Sprintf("token_purchase:%s", verification.TransactionID)).Result()
	if err != nil {
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, map[string]string{
			"error": "Transaction not found",
		})
		return
	}

	var transactionData map[string]interface{}
	if err := json.Unmarshal([]byte(data), &transactionData); err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{
			"error": "Invalid transaction data",
		})
		return
	}

	// Check if already processed
	if transactionData["status"] == "completed" {
		render.JSON(w, r, map[string]string{
			"status": "already_processed",
		})
		return
	}

	userID := transactionData["user_id"].(string)
	tokens := int(transactionData["tokens"].(float64))

	// Add tokens to user account
	if err := db.Model(&entities.User{}).Where("id = ?", userID).Update("vocatoken_val", gorm.Expr("vocatoken_val + ?", tokens)).Error; err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{
			"error": "Failed to update user tokens",
		})
		return
	}

	// Update transaction status
	transactionData["status"] = "completed"
	transactionData["completed_at"] = time.Now()
	transactionData["verification_data"] = verification
	transactionJSON, _ := json.Marshal(transactionData)
	cache.Redis().Set(r.Context(), fmt.Sprintf("token_purchase:%s", verification.TransactionID), string(transactionJSON), 30*24*time.Hour) // Keep for 30 days

	render.JSON(w, r, map[string]interface{}{
		"status":  "success",
		"tokens":  tokens,
		"message": "Payment processed successfully",
	})
}

// GetAvailablePaymentProviders returns available payment providers for user's location
func GetAvailablePaymentProviders(w http.ResponseWriter, r *http.Request) {
	// Get detection factors
	userIP := geolocation.GetRealIP(r)
	acceptLanguage := r.Header.Get("Accept-Language")
	timezone := r.Header.Get("X-Timezone")

	geoService := geolocation.NewGeoLocationService()
	countryCode, err := geoService.GetCountryFromIP(r.Context(), userIP)
	if err != nil {
		fmt.Printf("Geolocation error for IP %s: %v\n", userIP, err)
	}

	// Get smart provider config
	primaryConfig := geolocation.GetPaymentProviderForCountry(countryCode, acceptLanguage, timezone)

	// Prepare response with both primary and alternative options
	providers := []map[string]interface{}{
		{
			"provider":    primaryConfig.Provider,
			"currency":    primaryConfig.Currency,
			"recommended": true,
			"reason":      primaryConfig.Reason,
			"confidence":  primaryConfig.Confidence,
		},
		{
			"provider":    primaryConfig.AlternativeProvider,
			"currency":    primaryConfig.AlternativeCurrency,
			"recommended": false,
			"reason":      "Alternative option",
			"confidence":  30,
		},
	}

	response := map[string]interface{}{
		"success":           true,
		"detected_country":  countryCode,
		"detected_language": acceptLanguage,
		"providers":         providers,
		"default_provider":  primaryConfig.Provider,
		"default_currency":  primaryConfig.Currency,
	}

	render.JSON(w, r, response)
}

// RequestTeacher handles teacher role requests
func RequestTeacher(w http.ResponseWriter, r *http.Request) {
	userID := token.User(r).UserID

	if userID == "" {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{
			"error": i18n.Localizer(r, "error.user_id_required"),
		})
		return
	}

	var reqBody struct {
		Description string `json:"description"`
	}

	if err := render.DecodeJSON(r.Body, &reqBody); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{
			"error": i18n.Localizer(r, "error.invalid_request_body"),
		})
		return
	}

	if len(reqBody.Description) < 50 {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{
			"error": i18n.Localizer(r, "error.teacher_description_too_short"),
		})
		return
	}

	// Store teacher request in Redis
	redisClient := cache.Redis()
	requestKey := fmt.Sprintf("teacher_request:%s", userID)

	// Prepare request data
	requestData := map[string]interface{}{
		"user_id":      userID,
		"description":  reqBody.Description,
		"requested_at": time.Now().Unix(),
		"status":       "pending",
	}

	requestJSON, _ := json.Marshal(requestData)

	// Set teacher request with expiration (e.g., 30 days)
	err := redisClient.Set(r.Context(), requestKey, requestJSON, 30*24*time.Hour).Err()

	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{
			"error": i18n.Localizer(r, "error.something_went_wrong"),
		})
		return
	}

	render.JSON(w, r, map[string]string{
		"message": i18n.Localizer(r, "success.teacher_request_submitted"),
	})
}

// UpdateLanguagePreferences handles updating user language preferences
func UpdateLanguagePreferences(w http.ResponseWriter, r *http.Request) {
	db := database.Manager()
	userID := token.User(r).UserID

	var reqBody struct {
		KnownLanguages  []string `json:"known_languages"`
		TargetLanguages []string `json:"target_languages"`
	}

	if err := render.DecodeJSON(r.Body, &reqBody); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{
			"error": i18n.Localizer(r, "error.invalid_request_body"),
		})
		return
	}

	var user entities.User
	if err := db.Where("id = ?", userID).First(&user).Error; err != nil {
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, map[string]string{
			"error": i18n.Localizer(r, "error.user_not_found"),
		})
		return
	}
	kl, _ := json.Marshal(reqBody.KnownLanguages)
	tl, _ := json.Marshal(reqBody.TargetLanguages)

	// Update language preferences
	updates := map[string]any{
		"known_languages":  datatypes.JSON(kl),
		"target_languages": datatypes.JSON(tl),
	}

	if err := db.Model(&user).Updates(updates).Error; err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{
			"error": i18n.Localizer(r, "error.something_went_wrong"),
		})
		return
	}

	render.JSON(w, r, map[string]string{
		"message": i18n.Localizer(r, "success.language_preferences_updated"),
	})
}

// GetLanguagePreferences handles getting user language preferences
func GetLanguagePreferences(w http.ResponseWriter, r *http.Request) {
	db := database.Manager()
	userID := token.User(r).UserID

	if userID == "" {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{
			"error": i18n.Localizer(r, "error.user_id_required"),
		})
		return
	}

	var user entities.User
	if err := db.Where("id = ?", userID).First(&user).Error; err != nil {
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, map[string]string{
			"error": i18n.Localizer(r, "error.user_not_found"),
		})
		return
	}

	render.JSON(w, r, map[string]interface{}{
		"known_languages":  user.KnownLanguages,
		"target_languages": user.TargetLanguages,
	})
}
