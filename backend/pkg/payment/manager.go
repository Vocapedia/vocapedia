package payment

import (
	"context"
	"fmt"
	"sync"

	"github.com/akifkadioglu/vocapedia/pkg/payment/common"
)

// Manager manages different payment providers
type Manager struct {
	providers map[string]PaymentProvider
	mu        sync.RWMutex
}

// NewManager creates a new payment manager
func NewManager() *Manager {
	return &Manager{
		providers: make(map[string]PaymentProvider),
	}
}

// RegisterProvider registers a payment provider
func (m *Manager) RegisterProvider(name string, provider PaymentProvider) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.providers[name] = provider
}

// GetProvider returns a payment provider by name
func (m *Manager) GetProvider(name string) (PaymentProvider, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	provider, exists := m.providers[name]
	if !exists {
		return nil, fmt.Errorf("%w: %s", common.ErrProviderNotFound, name)
	}

	return provider, nil
}

// GetAvailableProviders returns list of available provider names
func (m *Manager) GetAvailableProviders() []string {
	m.mu.RLock()
	defer m.mu.RUnlock()

	var providers []string
	for name := range m.providers {
		providers = append(providers, name)
	}

	return providers
}

// CreatePayment creates a payment using the specified provider
func (m *Manager) CreatePayment(ctx context.Context, providerName string, request common.PaymentRequest) (*common.PaymentResponse, error) {
	provider, err := m.GetProvider(providerName)
	if err != nil {
		return nil, err
	}

	return provider.CreatePayment(ctx, request)
}

// VerifyPayment verifies a payment using the specified provider
func (m *Manager) VerifyPayment(ctx context.Context, providerName string, data map[string]interface{}) (*common.PaymentVerification, error) {
	provider, err := m.GetProvider(providerName)
	if err != nil {
		return nil, err
	}

	return provider.VerifyPayment(ctx, data)
}

// GetPaymentStatus gets payment status using the specified provider
func (m *Manager) GetPaymentStatus(ctx context.Context, providerName string, paymentID string) (*common.PaymentStatus, error) {
	provider, err := m.GetProvider(providerName)
	if err != nil {
		return nil, err
	}

	return provider.GetPaymentStatus(ctx, paymentID)
}

// Global manager instance
var globalManager *Manager
var once sync.Once

// GetManager returns the global payment manager instance
func GetManager() *Manager {
	once.Do(func() {
		globalManager = NewManager()
	})
	return globalManager
}
