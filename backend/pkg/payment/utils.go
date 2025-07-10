package payment

import (
	"fmt"
)

// CalculateTokenPrice calculates the price for tokens based on quantity tiers
func CalculateTokenPrice(tokenCount int) int {
	if tokenCount < 1 {
		return 0
	}

	if tokenCount < 50 {
		return tokenCount * TokenPriceTier1
	} else if tokenCount < 100 {
		return tokenCount * TokenPriceTier2
	} else {
		return tokenCount * TokenPriceTier3
	}
}

// ValidateTokenCount validates if the token count is within acceptable limits
func ValidateTokenCount(tokenCount int) error {
	if tokenCount < 1 {
		return fmt.Errorf("token count must be at least 1")
	}
	if tokenCount > 1000 {
		return fmt.Errorf("token count cannot exceed 1000")
	}
	return nil
}

// FormatPrice formats price in cents to a human-readable string
func FormatPrice(cents int) string {
	dollars := float64(cents) / 100
	return fmt.Sprintf("$%.2f", dollars)
}
