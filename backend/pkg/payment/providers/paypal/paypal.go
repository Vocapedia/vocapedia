package paypal

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/akifkadioglu/vocapedia/pkg/payment/common"
)

// PayPalProvider implements the PaymentProvider interface for PayPal
type PayPalProvider struct {
	ClientID     string
	ClientSecret string
	BaseURL      string
	IsLive       bool
}

// NewPayPalProvider creates a new PayPal payment provider
func NewPayPalProvider(clientID, clientSecret string, isLive bool) *PayPalProvider {
	baseURL := "https://api.sandbox.paypal.com"
	if isLive {
		baseURL = "https://api.paypal.com"
	}

	return &PayPalProvider{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		BaseURL:      baseURL,
		IsLive:       isLive,
	}
}

// GetProviderName returns the provider name
func (p *PayPalProvider) GetProviderName() string {
	return "paypal"
}

// CreatePayment creates a payment session with PayPal
func (p *PayPalProvider) CreatePayment(ctx context.Context, request common.PaymentRequest) (*common.PaymentResponse, error) {
	// Get access token first
	accessToken, err := p.getAccessToken(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get PayPal access token: %w", err)
	}

	// Create payment request
	paymentRequest := map[string]interface{}{
		"intent": "CAPTURE",
		"purchase_units": []map[string]interface{}{
			{
				"reference_id": request.TransactionID,
				"amount": map[string]interface{}{
					"currency_code": request.Currency,
					"value":         fmt.Sprintf("%.2f", request.Amount),
				},
				"description": request.Description,
			},
		},
		"application_context": map[string]interface{}{
			"return_url": request.CallbackURL,
			"cancel_url": request.CallbackURL + "?cancelled=true",
		},
	}

	jsonData, err := json.Marshal(paymentRequest)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal payment request: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", p.BaseURL+"/v2/checkout/orders", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+accessToken)

	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	if resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("paypal API error: %s", string(body))
	}

	var paypalResp map[string]interface{}
	if err := json.Unmarshal(body, &paypalResp); err != nil {
		return nil, fmt.Errorf("failed to parse PayPal response: %w", err)
	}

	orderID, ok := paypalResp["id"].(string)
	if !ok {
		return nil, fmt.Errorf("missing order ID in PayPal response")
	}

	// Extract approval URL
	var approvalURL string
	if links, ok := paypalResp["links"].([]interface{}); ok {
		for _, link := range links {
			if linkMap, ok := link.(map[string]interface{}); ok {
				if rel, ok := linkMap["rel"].(string); ok && rel == "approve" {
					if href, ok := linkMap["href"].(string); ok {
						approvalURL = href
						break
					}
				}
			}
		}
	}

	if approvalURL == "" {
		return nil, fmt.Errorf("missing approval URL in PayPal response")
	}

	return &common.PaymentResponse{
		PaymentID:  orderID,
		PaymentURL: approvalURL,
		ProviderID: orderID,
		Status:     common.StatusPending,
	}, nil
}

// VerifyPayment verifies a payment callback from PayPal
func (p *PayPalProvider) VerifyPayment(ctx context.Context, data map[string]interface{}) (*common.PaymentVerification, error) {
	orderID, ok := data["orderID"].(string)
	if !ok {
		return nil, fmt.Errorf("missing orderID in callback data")
	}

	// Get access token
	accessToken, err := p.getAccessToken(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get PayPal access token: %w", err)
	}

	// Get order details
	req, err := http.NewRequestWithContext(ctx, "GET", p.BaseURL+"/v2/checkout/orders/"+orderID, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+accessToken)

	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("paypal API error: %s", string(body))
	}

	var orderResp map[string]interface{}
	if err := json.Unmarshal(body, &orderResp); err != nil {
		return nil, fmt.Errorf("failed to parse PayPal response: %w", err)
	}

	status, ok := orderResp["status"].(string)
	if !ok {
		return nil, fmt.Errorf("missing status in PayPal response")
	}

	isValid := status == "COMPLETED"
	paymentStatus := common.StatusFailed
	if isValid {
		paymentStatus = common.StatusCompleted
	} else if status == "APPROVED" {
		paymentStatus = common.StatusPending
	}

	// Extract amount and transaction ID
	var amount float64
	var transactionID string
	if purchaseUnits, ok := orderResp["purchase_units"].([]interface{}); ok && len(purchaseUnits) > 0 {
		if unit, ok := purchaseUnits[0].(map[string]interface{}); ok {
			if refID, ok := unit["reference_id"].(string); ok {
				transactionID = refID
			}
			if amountObj, ok := unit["amount"].(map[string]interface{}); ok {
				if value, ok := amountObj["value"].(string); ok {
					fmt.Sscanf(value, "%f", &amount)
				}
			}
		}
	}

	return &common.PaymentVerification{
		IsValid:       isValid,
		PaymentID:     orderID,
		TransactionID: transactionID,
		Amount:        amount,
		Currency:      "USD", // PayPal typically uses USD
		Status:        paymentStatus,
		ProviderData:  orderResp,
	}, nil
}

// GetPaymentStatus checks the status of a payment
func (p *PayPalProvider) GetPaymentStatus(ctx context.Context, paymentID string) (*common.PaymentStatus, error) {
	// Get access token
	accessToken, err := p.getAccessToken(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get PayPal access token: %w", err)
	}

	// Get order details
	req, err := http.NewRequestWithContext(ctx, "GET", p.BaseURL+"/v2/checkout/orders/"+paymentID, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+accessToken)

	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("paypal API error: %s", string(body))
	}

	var orderResp map[string]interface{}
	if err := json.Unmarshal(body, &orderResp); err != nil {
		return nil, fmt.Errorf("failed to parse PayPal response: %w", err)
	}

	status, _ := orderResp["status"].(string)
	var paymentStatus string
	switch status {
	case "COMPLETED":
		paymentStatus = common.StatusCompleted
	case "APPROVED":
		paymentStatus = common.StatusPending
	default:
		paymentStatus = common.StatusFailed
	}

	return &common.PaymentStatus{
		PaymentID:    paymentID,
		Status:       paymentStatus,
		ProviderData: orderResp,
	}, nil
}

// getAccessToken gets an access token from PayPal
func (p *PayPalProvider) getAccessToken(ctx context.Context) (string, error) {
	data := "grant_type=client_credentials"
	req, err := http.NewRequestWithContext(ctx, "POST", p.BaseURL+"/v1/oauth2/token", bytes.NewBufferString(data))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.SetBasicAuth(p.ClientID, p.ClientSecret)

	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("paypal auth error: %s", string(body))
	}

	var tokenResp map[string]interface{}
	if err := json.Unmarshal(body, &tokenResp); err != nil {
		return "", fmt.Errorf("failed to parse token response: %w", err)
	}

	accessToken, ok := tokenResp["access_token"].(string)
	if !ok {
		return "", fmt.Errorf("missing access token in response")
	}

	return accessToken, nil
}
