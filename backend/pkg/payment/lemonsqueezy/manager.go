package lemonsqueezy

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/akifkadioglu/vocapedia/pkg/config"
	"github.com/akifkadioglu/vocapedia/pkg/payment"
)

// CreateCheckout implements PaymentProvider.CreateCheckout
func (p *LemonSqueezyProvider) CreateCheckout(ctx context.Context, req *payment.CheckoutRequest) (*payment.CheckoutResult, error) {
	payload := payment.LemonSqueezyCheckoutPayload{}
	payload.Data.Type = "checkouts"
	payload.Data.Attributes.ProductID = req.ProductID
	payload.Data.Attributes.VariantID = req.VariantID
	payload.Data.Attributes.CustomQuantity = req.TokenCount
	payload.Data.Attributes.CustomPrice = req.PriceCents
	payload.Data.Attributes.BillingEmail = req.Email

	// Add custom data for webhook processing
	if req.CustomData == nil {
		req.CustomData = make(map[string]interface{})
	}
	req.CustomData["user_id"] = req.UserID
	req.CustomData["token_count"] = req.TokenCount
	payload.Data.Attributes.CheckoutData.Custom = req.CustomData

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal payload: %w", err)
	}

	httpReq, err := http.NewRequestWithContext(ctx, "POST", config.ReadValue().Payments.Providers.LemonSqueezy.BaseURL+"/checkouts", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	httpReq.Header.Set("Authorization", "Bearer "+p.apiKey)
	httpReq.Header.Set("Content-Type", "application/vnd.api+json")
	httpReq.Header.Set("Accept", "application/vnd.api+json")

	resp, err := p.httpClient.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("LemonSqueezy API error (status %d): %s", resp.StatusCode, string(body))
	}

	var result payment.LemonSqueezyCheckoutResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return &payment.CheckoutResult{
		CheckoutURL: result.Data.Attributes.URL,
		ProviderID:  result.Data.ID,
	}, nil
}

// ProcessWebhook implements PaymentProvider.ProcessWebhook
func (p *LemonSqueezyProvider) ProcessWebhook(ctx context.Context, payload []byte, signature string) (*payment.WebhookResult, error) {
	// Verify webhook signature (if you have webhook secret configured)
	// This is optional but recommended for security

	var webhookPayload payment.LemonSqueezyWebhookPayload
	if err := json.Unmarshal(payload, &webhookPayload); err != nil {
		return nil, fmt.Errorf("failed to unmarshal webhook payload: %w", err)
	}

	// Map LemonSqueezy status to our internal status
	status := mapLemonSqueezyStatus(webhookPayload.Data.Attributes.Status)

	return &payment.WebhookResult{
		EventType:  webhookPayload.Meta.EventName,
		ProviderID: webhookPayload.Data.ID,
		Status:     status,
		UserEmail:  webhookPayload.Data.Attributes.UserEmail,
		CustomData: webhookPayload.Meta.CustomData,
	}, nil
}

// GetPaymentStatus implements PaymentProvider.GetPaymentStatus
func (p *LemonSqueezyProvider) GetPaymentStatus(ctx context.Context, paymentID string) (string, error) {
	httpReq, err := http.NewRequestWithContext(ctx, "GET", config.ReadValue().Payments.Providers.LemonSqueezy.BaseURL+"/orders/"+paymentID, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	httpReq.Header.Set("Authorization", "Bearer "+p.apiKey)
	httpReq.Header.Set("Accept", "application/vnd.api+json")

	resp, err := p.httpClient.Do(httpReq)
	if err != nil {
		return "", fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return "", fmt.Errorf("LemonSqueezy API error (status %d)", resp.StatusCode)
	}

	var result struct {
		Data struct {
			Attributes struct {
				Status string `json:"status"`
			} `json:"attributes"`
		} `json:"data"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", fmt.Errorf("failed to decode response: %w", err)
	}

	return mapLemonSqueezyStatus(result.Data.Attributes.Status), nil
}

// mapLemonSqueezyStatus maps LemonSqueezy status to our internal status
func mapLemonSqueezyStatus(lemonStatus string) string {
	switch lemonStatus {
	case "paid":
		return payment.PaymentStatusCompleted
	case "pending":
		return payment.PaymentStatusPending
	case "cancelled", "refunded":
		return payment.PaymentStatusCancelled
	case "failed":
		return payment.PaymentStatusFailed
	default:
		return payment.PaymentStatusPending
	}
}

// verifyWebhookSignature verifies the webhook signature (if you have webhook secret)
func (p *LemonSqueezyProvider) verifyWebhookSignature(payload []byte, signature, secret string) bool {
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write(payload)
	expectedSignature := hex.EncodeToString(mac.Sum(nil))
	return hmac.Equal([]byte(signature), []byte(expectedSignature))
}
