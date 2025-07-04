package iyzico

import (
	"bytes"
	"context"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/akifkadioglu/vocapedia/pkg/payment/common"
)

// IyzicoProvider implements the PaymentProvider interface for Iyzico
type IyzicoProvider struct {
	APIKey    string
	SecretKey string
	BaseURL   string
	IsLive    bool
}

// NewIyzicoProvider creates a new Iyzico payment provider
func NewIyzicoProvider(apiKey, secretKey string, isLive bool) *IyzicoProvider {
	baseURL := "https://sandbox-api.iyzipay.com"
	if isLive {
		baseURL = "https://api.iyzipay.com"
	}

	return &IyzicoProvider{
		APIKey:    apiKey,
		SecretKey: secretKey,
		BaseURL:   baseURL,
		IsLive:    isLive,
	}
}

// GetProviderName returns the provider name
func (p *IyzicoProvider) GetProviderName() string {
	return "iyzico"
}

// CreatePayment creates a payment session with Iyzico
func (p *IyzicoProvider) CreatePayment(ctx context.Context, request common.PaymentRequest) (*common.PaymentResponse, error) {
	// Iyzico checkout form request
	checkoutRequest := map[string]interface{}{
		"locale":              "tr",
		"conversationId":      request.TransactionID,
		"price":               fmt.Sprintf("%.2f", request.Amount),
		"paidPrice":           fmt.Sprintf("%.2f", request.Amount),
		"currency":            request.Currency,
		"basketId":            request.TransactionID,
		"paymentGroup":        "PRODUCT",
		"callbackUrl":         request.CallbackURL,
		"enabledInstallments": []int{1},
		"buyer": map[string]string{
			"id":                  request.UserID,
			"name":                request.UserName,
			"surname":             "User",
			"gsmNumber":           "+905350000000",
			"email":               request.UserEmail,
			"identityNumber":      "74300864791",
			"lastLoginDate":       time.Now().Format("2006-01-02 15:04:05"),
			"registrationDate":    time.Now().Format("2006-01-02 15:04:05"),
			"registrationAddress": "Nidakule Göztepe, Merdivenköy Mah. Bora Sok. No:1",
			"ip":                  "85.34.78.112",
			"city":                "Istanbul",
			"country":             "Turkey",
			"zipCode":             "34732",
		},
		"shippingAddress": map[string]string{
			"contactName": request.UserName,
			"city":        "Istanbul",
			"country":     "Turkey",
			"address":     "Nidakule Göztepe, Merdivenköy Mah. Bora Sok. No:1",
			"zipCode":     "34732",
		},
		"billingAddress": map[string]string{
			"contactName": request.UserName,
			"city":        "Istanbul",
			"country":     "Turkey",
			"address":     "Nidakule Göztepe, Merdivenköy Mah. Bora Sok. No:1",
			"zipCode":     "34732",
		},
		"basketItems": []map[string]string{
			{
				"id":        "TOKEN001",
				"name":      request.Description,
				"category1": "Token",
				"itemType":  "VIRTUAL",
				"price":     fmt.Sprintf("%.2f", request.Amount),
			},
		},
	}

	response, err := p.makeRequest(ctx, "/payment/iyzipos/checkoutform/initialize/auth/ecom", checkoutRequest)
	if err != nil {
		return nil, err
	}

	// Parse Iyzico response
	var iyzicoResp map[string]interface{}
	if err := json.Unmarshal(response, &iyzicoResp); err != nil {
		return nil, fmt.Errorf("failed to parse Iyzico response: %w", err)
	}

	if status, ok := iyzicoResp["status"].(string); !ok || status != "success" {
		return nil, fmt.Errorf("iyzico payment creation failed: %v", iyzicoResp["errorMessage"])
	}

	paymentPageURL, ok := iyzicoResp["paymentPageUrl"].(string)
	if !ok {
		return nil, fmt.Errorf("missing payment page URL in Iyzico response")
	}

	checkoutFormToken, ok := iyzicoResp["checkoutFormToken"].(string)
	if !ok {
		return nil, fmt.Errorf("missing checkout form token in Iyzico response")
	}

	return &common.PaymentResponse{
		PaymentID:  checkoutFormToken,
		PaymentURL: paymentPageURL,
		ProviderID: checkoutFormToken,
		Status:     common.StatusPending,
	}, nil
}

// VerifyPayment verifies a payment callback from Iyzico
func (p *IyzicoProvider) VerifyPayment(ctx context.Context, data map[string]interface{}) (*common.PaymentVerification, error) {
	token, ok := data["token"].(string)
	if !ok {
		return nil, fmt.Errorf("missing token in callback data")
	}

	// Retrieve checkout form result
	retrieveRequest := map[string]interface{}{
		"locale":            "tr",
		"conversationId":    data["conversationId"],
		"checkoutFormToken": token,
	}

	response, err := p.makeRequest(ctx, "/payment/iyzipos/checkoutform/auth/ecom/detail", retrieveRequest)
	if err != nil {
		return nil, err
	}

	var iyzicoResp map[string]interface{}
	if err := json.Unmarshal(response, &iyzicoResp); err != nil {
		return nil, fmt.Errorf("failed to parse Iyzico response: %w", err)
	}

	status, ok := iyzicoResp["status"].(string)
	if !ok {
		return nil, fmt.Errorf("missing status in Iyzico response")
	}

	isValid := status == "success"
	paymentStatus := common.StatusFailed
	if isValid {
		paymentStatus = common.StatusCompleted
	}

	conversationId, _ := iyzicoResp["conversationId"].(string)
	paidPrice, _ := iyzicoResp["paidPrice"].(float64)

	return &common.PaymentVerification{
		IsValid:       isValid,
		PaymentID:     token,
		TransactionID: conversationId,
		Amount:        paidPrice,
		Currency:      "TRY",
		Status:        paymentStatus,
		ProviderData:  iyzicoResp,
	}, nil
}

// GetPaymentStatus checks the status of a payment
func (p *IyzicoProvider) GetPaymentStatus(ctx context.Context, paymentID string) (*common.PaymentStatus, error) {
	// This would typically query Iyzico's API for payment status
	// For now, we'll return a basic implementation
	return &common.PaymentStatus{
		PaymentID: paymentID,
		Status:    common.StatusPending,
	}, nil
}

// makeRequest makes an HTTP request to Iyzico API
func (p *IyzicoProvider) makeRequest(ctx context.Context, endpoint string, data map[string]interface{}) ([]byte, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request data: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", p.BaseURL+endpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Generate authorization header
	randomKey := fmt.Sprintf("%d", time.Now().UnixNano())
	authString := p.generateAuthString(string(jsonData), randomKey)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", authString)
	req.Header.Set("x-iyzi-rnd", randomKey)

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
		return nil, fmt.Errorf("iyzico API error: %s", string(body))
	}

	return body, nil
}

// generateAuthString generates the authorization string for Iyzico API
func (p *IyzicoProvider) generateAuthString(requestBody, randomKey string) string {
	hashData := p.APIKey + randomKey + p.SecretKey + requestBody
	hash := sha1.Sum([]byte(hashData))
	hashString := base64.StdEncoding.EncodeToString(hash[:])
	return fmt.Sprintf("IYZWS %s:%s", p.APIKey, hashString)
}
