package payment

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/akifkadioglu/vocapedia/pkg/config"
	"github.com/akifkadioglu/vocapedia/pkg/database"
	"github.com/akifkadioglu/vocapedia/pkg/entities"
	paymentpkg "github.com/akifkadioglu/vocapedia/pkg/payment"
	"github.com/akifkadioglu/vocapedia/pkg/payment/lemonsqueezy"
	"github.com/akifkadioglu/vocapedia/pkg/token"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

var lemonProvider *lemonsqueezy.LemonSqueezyProvider

// PurchaseTokens handles token purchase requests
func PurchaseTokens(w http.ResponseWriter, r *http.Request) {
	var req paymentpkg.PurchaseRequest
	var user entities.User

	if err := render.DecodeJSON(r.Body, &req); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{
			"error": "Invalid request body",
		})
		return
	}

	// Validate token count
	if err := paymentpkg.ValidateTokenCount(req.TokenCount); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{
			"error": err.Error(),
		})
		return
	}

	userID := token.User(r).UserID
	db := database.Manager()

	if err := db.First(&user, userID).Error; err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{
			"error": "Failed to retrieve user",
		})
		return
	}

	price := paymentpkg.CalculateTokenPrice(req.TokenCount)
	cfg := config.ReadValue()

	// Create checkout request for LemonSqueezy
	checkoutReq := &lemonsqueezy.CheckoutRequest{
		UserID:     userID,
		Email:      user.Email,
		TokenCount: req.TokenCount,
		PriceCents: price,
		ProductID:  cfg.Payments.Providers.LemonSqueezy.ProductID,
		VariantID:  cfg.Payments.Providers.LemonSqueezy.VariantID,
		CustomData: map[string]interface{}{
			"user_id":     userID,
			"token_count": req.TokenCount,
		},
	}

	result, err := lemonProvider.CreateCheckout(context.Background(), (*paymentpkg.CheckoutRequest)(checkoutReq))
	if err != nil {
		log.Printf("Failed to create checkout: %v", err)
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{
			"error": "Failed to create checkout session",
		})
		return
	}

	// Save payment record to database
	payment := &entities.PaymentRecord{
		UserID:      userID,
		Provider:    paymentpkg.ProviderLemonSqueezy,
		ProviderID:  result.ProviderID,
		TokenCount:  req.TokenCount,
		PriceCents:  price,
		Status:      paymentpkg.PaymentStatusPending,
		CheckoutURL: result.CheckoutURL,
	}

	repo := paymentpkg.NewGormPaymentRepository(db)
	if err := repo.Create(context.Background(), payment); err != nil {
		log.Printf("Failed to save payment record: %v", err)
		// Continue anyway since checkout was created
	}

	log.Printf("Created checkout for user %s: %s", userID, result.CheckoutURL)

	render.JSON(w, r, paymentpkg.CheckoutResponse{
		CheckoutURL: result.CheckoutURL,
		PaymentID:   fmt.Sprintf("%v", payment.ID),
	})
}

// GetPaymentHistory returns user's payment history
func GetPaymentHistory(w http.ResponseWriter, r *http.Request) {
	userID := token.User(r).UserID

	repo := paymentpkg.NewGormPaymentRepository(database.Manager())
	payments, err := repo.GetUserPayments(context.Background(), userID)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{
			"error": "Failed to retrieve payment history",
		})
		return
	}

	render.JSON(w, r, map[string]interface{}{
		"payments": payments,
	})
}

// GetPaymentStats returns user's payment statistics
func GetPaymentStats(w http.ResponseWriter, r *http.Request) {
	userID := token.User(r).UserID

	repo := paymentpkg.NewGormPaymentRepository(database.Manager())
	stats, err := repo.GetPaymentStats(context.Background(), userID)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{
			"error": "Failed to retrieve payment statistics",
		})
		return
	}

	render.JSON(w, r, stats)
}

// HandleLemonSqueezyWebhook handles LemonSqueezy webhooks
func HandleLemonSqueezyWebhook(w http.ResponseWriter, r *http.Request) {
	payload, err := io.ReadAll(r.Body)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{
			"error": "Failed to read request body",
		})
		return
	}

	signature := r.Header.Get("X-Signature")

	result, err := lemonProvider.ProcessWebhook(context.Background(), payload, signature)
	if err != nil {
		log.Printf("Failed to process webhook: %v", err)
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{
			"error": "Failed to process webhook",
		})
		return
	}

	// Find payment by provider ID and update status
	repo := paymentpkg.NewGormPaymentRepository(database.Manager())
	payment, err := repo.GetByProviderID(context.Background(), result.ProviderID)
	if err != nil {
		log.Printf("Payment not found for provider ID %s: %v", result.ProviderID, err)
		w.WriteHeader(http.StatusOK) // Still return 200 to acknowledge webhook
		return
	}

	// Update payment status
	if result.Status != payment.Status {
		if err := repo.UpdateStatus(context.Background(), fmt.Sprintf("%v", payment.ID), result.Status); err != nil {
			log.Printf("Failed to update payment status: %v", err)
		}
	}

	w.WriteHeader(http.StatusOK)
}

// GetPaymentStatus returns the status of a specific payment
func GetPaymentStatus(w http.ResponseWriter, r *http.Request) {
	paymentID := chi.URLParam(r, "paymentID")
	if paymentID == "" {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{
			"error": "Payment ID is required",
		})
		return
	}

	repo := paymentpkg.NewGormPaymentRepository(database.Manager())
	payment, err := repo.GetByID(context.Background(), paymentID)
	if err != nil {
		if err == paymentpkg.ErrPaymentNotFound {
			render.Status(r, http.StatusNotFound)
			render.JSON(w, r, map[string]string{
				"error": "Payment not found",
			})
			return
		}
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{
			"error": "Failed to retrieve payment",
		})
		return
	}

	// Verify user owns this payment
	userID := token.User(r).UserID
	if payment.UserID != userID {
		render.Status(r, http.StatusForbidden)
		render.JSON(w, r, map[string]string{
			"error": "Access denied",
		})
		return
	}

	render.JSON(w, r, payment)
}
