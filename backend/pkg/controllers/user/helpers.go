package user

// Calculate token price based on quantity and currency
func calculateTokenPrice(tokens int, currency string) float64 {
	basePrice := float64(tokens)

	// Apply discount based on quantity
	if tokens >= 100 {
		basePrice = basePrice * 0.8 // 20% discount
	} else if tokens >= 50 {
		basePrice = basePrice * 0.9 // 10% discount
	}

	// Convert to appropriate currency
	switch currency {
	case "USD":
		return basePrice * 0.03 // $0.03 per token
	case "TRY":
		return basePrice * 1.00 // 1 TRY per token
	default:
		return basePrice * 0.03 // Default to USD
	}
}

// Get discount percentage based on token count
func getDiscountPercentage(tokens int) int {
	if tokens < 50 {
		return 0
	} else if tokens < 100 {
		return 10
	} else {
		return 20
	}
}
