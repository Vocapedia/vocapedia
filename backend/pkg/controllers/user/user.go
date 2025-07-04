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
	"github.com/akifkadioglu/vocapedia/pkg/i18n"
	"github.com/akifkadioglu/vocapedia/pkg/mail"
	"github.com/akifkadioglu/vocapedia/pkg/search"
	"github.com/akifkadioglu/vocapedia/pkg/snowflake"
	"github.com/akifkadioglu/vocapedia/pkg/token"
	"github.com/akifkadioglu/vocapedia/pkg/utils"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/meilisearch/meilisearch-go"
	"github.com/redis/go-redis/v9"
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

	if streak.LastDate == today {
	} else if streak.LastDate == yesterday {
		streak.Count++
		if streak.Count > 7 {
			streak.Count = 1
			streak.Rewarded = false
		}

		streak.LastDate = today
	} else {
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

	// Calculate price and discount
	price := calculateTokenPrice(req.Tokens)
	discount := getDiscountPercentage(req.Tokens)

	// Create a transaction record
	transactionID := fmt.Sprintf("%d", snowflake.GenerateID())

	// Create Gumroad payment URL with custom pricing and transaction ID
	gumroadURL := fmt.Sprintf("https://vocabloom.gumroad.com/l/vocatoken?wanted=true&price=%d&transaction_id=%s",
		int(price), // Price in cents
		transactionID,
	)

	// Store transaction details in cache for later confirmation
	transactionData := map[string]interface{}{
		"user_id":        userID,
		"tokens":         req.Tokens,
		"price":          price,
		"discount":       discount,
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
		Discount:      discount,
		PaymentURL:    gumroadURL,
		TransactionID: transactionID,
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

// GumroadWebhook handles webhook from Gumroad when payment is completed
func GumroadWebhook(w http.ResponseWriter, r *http.Request) {
	// Parse webhook data from Gumroad
	var webhookData map[string]interface{}
	if err := render.DecodeJSON(r.Body, &webhookData); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{
			"error": "Invalid webhook data",
		})
		return
	}

	// Extract transaction details from webhook
	// Gumroad webhook structure might vary, adjust according to their documentation
	transactionID := ""
	if customFields, ok := webhookData["custom_fields"].(map[string]interface{}); ok {
		if tid, exists := customFields["transaction_id"].(string); exists {
			transactionID = tid
		}
	}

	// If no transaction ID in custom fields, try to extract from URL parameters
	if transactionID == "" {
		transactionID = r.URL.Query().Get("transaction_id")
	}

	if transactionID == "" {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{
			"error": "Transaction ID not found in webhook",
		})
		return
	}

	// Process the payment confirmation
	db := database.Manager()

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
	transactionData["webhook_data"] = webhookData
	transactionJSON, _ := json.Marshal(transactionData)
	cache.Redis().Set(r.Context(), fmt.Sprintf("token_purchase:%s", transactionID), string(transactionJSON), 30*24*time.Hour) // Keep for 30 days

	render.JSON(w, r, map[string]interface{}{
		"status":  "success",
		"tokens":  tokens,
		"message": "Payment processed successfully",
	})
}
