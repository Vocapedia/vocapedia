package user

// Calculate token price based on quantity (same logic as frontend)
func calculateTokenPrice(tokens int) float64 {
	if tokens < 50 {
		return float64(tokens) * 1.00
	} else if tokens < 100 {
		return float64(tokens) * 0.90
	} else {
		return float64(tokens) * 0.80
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
