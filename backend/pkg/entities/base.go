package entities

import (
	"context"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	"github.com/akifkadioglu/vocapedia/pkg/snowflake"
	"gorm.io/gorm"
)

type Base struct {
	ID        int64          `gorm:"primary_key" json:"id"`
	CreatedAt time.Time      `json:"created_at" example:"2021-01-01T00:00:00Z"`
	UpdatedAt time.Time      `json:"updated_at" example:"2021-01-01T00:00:00Z"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index" swaggertype:"string" format:"date-time"`
}

type Attrs map[string]any

func (a Attrs) Value() (driver.Value, error) {
	return json.Marshal(a)
}

func (a *Attrs) Scan(value any) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &a)
}

func (g *Base) BeforeCreate(tx *gorm.DB) (err error) {
	g.ID = snowflake.GenerateID()
	return nil
}

func (g *Base) ToSearch(ctx context.Context, db *gorm.DB) (map[string]any, error) {
	indexDto := map[string]any{
		"id":         g.ID,
		"created_at": g.CreatedAt.Unix(),
		"updated_at": g.UpdatedAt.Unix(),
		"deleted_at": g.DeletedAt.Time.Unix(),
	}
	return indexDto, nil
}
