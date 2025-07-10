package entities

import "time"

type PaymentRecord struct {
	Base
	UserID      string     `json:"user_id" gorm:"not null;index"`
	Provider    string     `json:"provider" gorm:"not null"`
	ProviderID  string     `json:"provider_id" gorm:"index"`
	TokenCount  int        `json:"token_count" gorm:"not null"`
	PriceCents  int        `json:"price_cents" gorm:"not null"`
	Status      string     `json:"status" gorm:"not null;default:'pending'"`
	CheckoutURL string     `json:"checkout_url"`
	CompletedAt *time.Time `json:"completed_at,omitempty"`
}
