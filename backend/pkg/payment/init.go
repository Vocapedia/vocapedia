package payment

import (
	"log"

	"github.com/akifkadioglu/vocapedia/pkg/config"
	"github.com/akifkadioglu/vocapedia/pkg/payment/providers/iyzico"
	"github.com/akifkadioglu/vocapedia/pkg/payment/providers/paypal"
)

// InitializeProviders initializes all payment providers
func InitializeProviders() {
	manager := GetManager()

	// Initialize Iyzico provider
	iyzicoProvider := iyzico.NewIyzicoProvider(
		config.ReadValue().Payments.Providers.Iyzico.APIKey,     // API Key from config
		config.ReadValue().Payments.Providers.Iyzico.SecretKey,  // Secret Key from config
		config.ReadValue().Payments.Providers.Iyzico.Production, // Set to true for production
	)
	manager.RegisterProvider("iyzico", iyzicoProvider)
	log.Println("✅ Iyzico payment provider registered")

	// Initialize PayPal provider
	paypalProvider := paypal.NewPayPalProvider(
		config.ReadValue().Payments.Providers.Paypal.ClientID,   // Client ID from config
		config.ReadValue().Payments.Providers.Paypal.SecretKey,  // Client Secret from config
		config.ReadValue().Payments.Providers.Paypal.Production, // Set to true for production
	)
	manager.RegisterProvider("paypal", paypalProvider)
	log.Println("✅ PayPal payment provider registered")

	log.Printf("✅ Payment system initialized with %d providers", len(manager.GetAvailableProviders()))
}
