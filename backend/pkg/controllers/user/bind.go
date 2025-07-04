package user

import "github.com/akifkadioglu/vocapedia/pkg/entities"

type _emailData struct {
	Header      string
	Description string
	Warning     string
	Device      string
}

type _updateUser struct {
	Name        string          `json:"name"`
	Username    string          `json:"username"`
	Description string          `json:"description"`
	Biography   string          `json:"biography"`
	Device      entities.Device `json:"device"`
}

type _streak struct {
	Count    int    `json:"count"`
	LastDate string `json:"lastDate"`
	Rewarded bool   `json:"rewarded"`
}

// Token purchase structures
type _tokenPurchaseRequest struct {
	Tokens int `json:"tokens" binding:"required,min=1"`
}

type _tokenPurchaseResponse struct {
	Success       bool    `json:"success"`
	Tokens        int     `json:"tokens"`
	Price         float64 `json:"price"`
	Discount      int     `json:"discount"`
	PaymentURL    string  `json:"payment_url"`
	TransactionID string  `json:"transaction_id"`
}

type _tokenPurchaseConfirmation struct {
	TransactionID string `json:"transaction_id"`
	PaymentStatus string `json:"payment_status"`
}
