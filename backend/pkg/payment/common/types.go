package common

import "errors"

// PaymentRequest represents a payment request
type PaymentRequest struct {
	Amount        float64           `json:"amount"`
	Currency      string            `json:"currency"`
	Description   string            `json:"description"`
	UserID        string            `json:"user_id"`
	UserEmail     string            `json:"user_email"`
	UserName      string            `json:"user_name"`
	TransactionID string            `json:"transaction_id"`
	CallbackURL   string            `json:"callback_url"`
	WebhookURL    string            `json:"webhook_url"`
	Metadata      map[string]string `json:"metadata"`
}

// PaymentResponse represents the response from creating a payment
type PaymentResponse struct {
	PaymentID  string `json:"payment_id"`
	PaymentURL string `json:"payment_url"`
	ProviderID string `json:"provider_id"`
	Status     string `json:"status"`
	ExpiresAt  string `json:"expires_at,omitempty"`
}

// PaymentVerification represents the result of payment verification
type PaymentVerification struct {
	IsValid       bool                   `json:"is_valid"`
	PaymentID     string                 `json:"payment_id"`
	TransactionID string                 `json:"transaction_id"`
	Amount        float64                `json:"amount"`
	Currency      string                 `json:"currency"`
	Status        string                 `json:"status"`
	ProviderData  map[string]interface{} `json:"provider_data"`
}

// PaymentStatus represents the status of a payment
type PaymentStatus struct {
	PaymentID     string                 `json:"payment_id"`
	TransactionID string                 `json:"transaction_id"`
	Status        string                 `json:"status"`
	Amount        float64                `json:"amount"`
	Currency      string                 `json:"currency"`
	CreatedAt     string                 `json:"created_at"`
	UpdatedAt     string                 `json:"updated_at"`
	ProviderData  map[string]interface{} `json:"provider_data"`
}

// Common payment statuses
const (
	StatusPending   = "pending"
	StatusCompleted = "completed"
	StatusFailed    = "failed"
	StatusCancelled = "cancelled"
	StatusExpired   = "expired"
)

// Common errors
var (
	ErrProviderNotFound    = errors.New("payment provider not found")
	ErrInvalidPayment      = errors.New("invalid payment data")
	ErrPaymentNotFound     = errors.New("payment not found")
	ErrPaymentVerification = errors.New("payment verification failed")
)
