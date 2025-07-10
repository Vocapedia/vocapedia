package lemonsqueezy

import (
	"net/http"
	"time"

	"github.com/akifkadioglu/vocapedia/pkg/config"
)

// CheckoutRequest represents a checkout creation request
type CheckoutRequest struct {
	UserID     string
	Email      string
	TokenCount int
	PriceCents int
	ProductID  string
	VariantID  string
	CustomData map[string]interface{}
}

// CheckoutResult represents the result of checkout creation
type CheckoutResult struct {
	CheckoutURL string
	PaymentID   string
	ProviderID  string
}

// WebhookResult represents the result of webhook processing
type WebhookResult struct {
	EventType  string
	PaymentID  string
	ProviderID string
	Status     string
	UserEmail  string
	CustomData map[string]interface{}
}

// LemonSqueezy specific structures
type LemonSqueezyCheckoutPayload struct {
	Data struct {
		Type       string `json:"type"`
		Attributes struct {
			ProductID      string `json:"product_id"`
			VariantID      string `json:"variant_id"`
			CustomQuantity int    `json:"custom_quantity,omitempty"`
			CustomPrice    int    `json:"custom_price,omitempty"`
			BillingEmail   string `json:"billing_email,omitempty"`
			CheckoutData   struct {
				Custom map[string]interface{} `json:"custom,omitempty"`
			} `json:"checkout_data,omitempty"`
		} `json:"attributes"`
	} `json:"data"`
}

type LemonSqueezyCheckoutResponse struct {
	Data struct {
		Type       string `json:"type"`
		ID         string `json:"id"`
		Attributes struct {
			URL       string `json:"url"`
			CreatedAt string `json:"created_at"`
			UpdatedAt string `json:"updated_at"`
		} `json:"attributes"`
	} `json:"data"`
}

type LemonSqueezyWebhookPayload struct {
	Meta struct {
		EventName  string                 `json:"event_name"`
		CustomData map[string]interface{} `json:"custom_data"`
	} `json:"meta"`
	Data struct {
		Type       string `json:"type"`
		ID         string `json:"id"`
		Attributes struct {
			Status    string `json:"status"`
			UserEmail string `json:"user_email"`
			CreatedAt string `json:"created_at"`
			UpdatedAt string `json:"updated_at"`
		} `json:"attributes"`
	} `json:"data"`
}

// LemonSqueezyProvider handles LemonSqueezy payments

// NewLemonSqueezyProvider creates a new LemonSqueezy provider
func NewLemonSqueezyProvider() *LemonSqueezyProvider {
	cfg := config.ReadValue()
	return &LemonSqueezyProvider{
		apiKey: cfg.Payments.Providers.LemonSqueezy.APIKey,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}
