package entities

import "github.com/google/uuid"

type JwtModel struct {
	UserID uuid.UUID `json:"user_id" example:"f0a0a1e9-69bd-4bef-b8c6-4e8c0d3a1212"`
}
