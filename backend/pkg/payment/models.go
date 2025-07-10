package payment

// Payment statuses
const (
	PaymentStatusPending   = "pending"
	PaymentStatusCompleted = "completed"
	PaymentStatusFailed    = "failed"
	PaymentStatusCancelled = "cancelled"
)

// Payment providers
const (
	ProviderLemonSqueezy = "lemonsqueezy"
)

// Token pricing tiers
const (
	TokenPriceTier1 = 100 // < 50 tokens: 1.00$ per token
	TokenPriceTier2 = 90  // 50-99 tokens: 0.90$ per token
	TokenPriceTier3 = 80  // 100+ tokens: 0.80$ per token
)

// Payment request structures
type PurchaseRequest struct {
	TokenCount int `json:"token_count" validate:"required,min=1,max=1000"`
}

type CheckoutResponse struct {
	CheckoutURL string `json:"checkout_url"`
	PaymentID   string `json:"payment_id,omitempty"`
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
