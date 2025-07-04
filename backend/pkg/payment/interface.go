package payment

import (
	"context"

	"github.com/akifkadioglu/vocapedia/pkg/payment/common"
)

// PaymentProvider defines the interface for payment providers
type PaymentProvider interface {
	// CreatePayment creates a payment session and returns payment URL
	CreatePayment(ctx context.Context, request common.PaymentRequest) (*common.PaymentResponse, error)

	// VerifyPayment verifies a payment using webhook data or callback
	VerifyPayment(ctx context.Context, data map[string]interface{}) (*common.PaymentVerification, error)

	// GetPaymentStatus checks the status of a payment
	GetPaymentStatus(ctx context.Context, paymentID string) (*common.PaymentStatus, error)

	// GetProviderName returns the name of the provider
	GetProviderName() string
}
