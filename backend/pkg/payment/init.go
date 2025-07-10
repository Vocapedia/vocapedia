package payment

import (
	"github.com/akifkadioglu/vocapedia/pkg/database"
	"github.com/akifkadioglu/vocapedia/pkg/entities"
)

var GlobalPaymentService *PaymentService

// InitializeProviders initializes the payment service
func InitializeProviders() {
	db := database.Manager()
	repo := NewGormPaymentRepository(db)
	GlobalPaymentService = NewPaymentService(repo)

	// Auto-migrate payment tables
	db.AutoMigrate(&entities.PaymentRecord{})
}
