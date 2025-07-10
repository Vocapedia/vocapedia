package payment

import (
	"context"
	"errors"
	"fmt"

	"github.com/akifkadioglu/vocapedia/pkg/entities"
)

var (
	ErrInvalidTokenCount = errors.New("invalid token count")
	ErrInvalidProvider   = errors.New("invalid payment provider")
	ErrPaymentFailed     = errors.New("payment processing failed")
	ErrPaymentNotFound   = errors.New("payment record not found")
	ErrWebhookInvalid    = errors.New("invalid webhook payload")
)

// PaymentProvider defines the interface for payment providers
type PaymentProvider interface {
	CreateCheckout(ctx context.Context, req *CheckoutRequest) (*CheckoutResult, error)
	ProcessWebhook(ctx context.Context, payload []byte, signature string) (*WebhookResult, error)
	GetPaymentStatus(ctx context.Context, paymentID string) (string, error)
}

// CheckoutRequest represents a checkout creation request
type CheckoutRequest struct {
	UserID     string
	Email      string
	TokenCount int
	PriceCents int
	ProductID  string
	VariantID  string
	CustomData map[string]any
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
	CustomData map[string]any
}

// PaymentService handles payment operations
type PaymentService struct {
	providers map[string]PaymentProvider
	repo      PaymentRepository
}

// PaymentRepository defines database operations for payments
type PaymentRepository interface {
	Create(ctx context.Context, payment *entities.PaymentRecord) error
	GetByID(ctx context.Context, id string) (*entities.PaymentRecord, error)
	GetByProviderID(ctx context.Context, providerID string) (*entities.PaymentRecord, error)
	UpdateStatus(ctx context.Context, id string, status string) error
	GetUserPayments(ctx context.Context, userID string) ([]*entities.PaymentRecord, error)
}

// NewPaymentService creates a new payment service
func NewPaymentService(repo PaymentRepository) *PaymentService {
	return &PaymentService{
		providers: make(map[string]PaymentProvider),
		repo:      repo,
	}
}

// RegisterProvider registers a payment provider
func (s *PaymentService) RegisterProvider(name string, provider PaymentProvider) {
	s.providers[name] = provider
}

// GetProvider returns a payment provider by name
func (s *PaymentService) GetProvider(name string) (PaymentProvider, error) {
	provider, exists := s.providers[name]
	if !exists {
		return nil, ErrInvalidProvider
	}
	return provider, nil
}

// CreateCheckout creates a checkout session
func (s *PaymentService) CreateCheckout(ctx context.Context, providerName string, req *CheckoutRequest) (*CheckoutResult, error) {
	provider, err := s.GetProvider(providerName)
	if err != nil {
		return nil, err
	}

	result, err := provider.CreateCheckout(ctx, req)
	if err != nil {
		return nil, err
	}

	// Save payment record to database
	payment := &entities.PaymentRecord{
		UserID:      req.UserID,
		Provider:    providerName,
		ProviderID:  result.ProviderID,
		TokenCount:  req.TokenCount,
		PriceCents:  req.PriceCents,
		Status:      PaymentStatusPending,
		CheckoutURL: result.CheckoutURL,
	}

	if err := s.repo.Create(ctx, payment); err != nil {
		return nil, err
	}

	result.PaymentID = fmt.Sprintf("%v", payment.ID)
	return result, nil
}

// ProcessWebhook processes a webhook from a payment provider
func (s *PaymentService) ProcessWebhook(ctx context.Context, providerName string, payload []byte, signature string) error {
	provider, err := s.GetProvider(providerName)
	if err != nil {
		return err
	}

	result, err := provider.ProcessWebhook(ctx, payload, signature)
	if err != nil {
		return err
	}

	// Find payment by provider ID
	payment, err := s.repo.GetByProviderID(ctx, result.ProviderID)
	if err != nil {
		return err
	}

	// Update payment status
	if result.Status != payment.Status {
		if err := s.repo.UpdateStatus(ctx, fmt.Sprintf("%v", payment.ID), result.Status); err != nil {
			return err
		}
	}

	return nil
}
