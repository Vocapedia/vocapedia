package payment

import (
	"context"
	"database/sql"
	"time"

	"github.com/akifkadioglu/vocapedia/pkg/entities"
	"gorm.io/gorm"
)

// GormPaymentRepository implements PaymentRepository using GORM
type GormPaymentRepository struct {
	db *gorm.DB
}

// NewGormPaymentRepository creates a new GORM payment repository
func NewGormPaymentRepository(db *gorm.DB) *GormPaymentRepository {
	return &GormPaymentRepository{db: db}
}

// Create implements PaymentRepository.Create
func (r *GormPaymentRepository) Create(ctx context.Context, payment *entities.PaymentRecord) error {
	return r.db.WithContext(ctx).Create(payment).Error
}

// GetByID implements PaymentRepository.GetByID
func (r *GormPaymentRepository) GetByID(ctx context.Context, id string) (*entities.PaymentRecord, error) {
	var payment entities.PaymentRecord
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&payment).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrPaymentNotFound
		}
		return nil, err
	}
	return &payment, nil
}

// GetByProviderID implements PaymentRepository.GetByProviderID
func (r *GormPaymentRepository) GetByProviderID(ctx context.Context, providerID string) (*entities.PaymentRecord, error) {
	var payment entities.PaymentRecord
	err := r.db.WithContext(ctx).Where("provider_id = ?", providerID).First(&payment).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrPaymentNotFound
		}
		return nil, err
	}
	return &payment, nil
}

// UpdateStatus implements PaymentRepository.UpdateStatus
func (r *GormPaymentRepository) UpdateStatus(ctx context.Context, id string, status string) error {
	updates := map[string]interface{}{
		"status":     status,
		"updated_at": time.Now(),
	}

	// Set completed_at if status is completed
	if status == PaymentStatusCompleted {
		now := time.Now()
		updates["completed_at"] = &now
	}

	result := r.db.WithContext(ctx).Model(&entities.PaymentRecord{}).Where("id = ?", id).Updates(updates)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return ErrPaymentNotFound
	}
	return nil
}

// GetUserPayments implements PaymentRepository.GetUserPayments
func (r *GormPaymentRepository) GetUserPayments(ctx context.Context, userID string) ([]*entities.PaymentRecord, error) {
	var payments []*entities.PaymentRecord
	err := r.db.WithContext(ctx).
		Where("user_id = ?", userID).
		Order("created_at DESC").
		Find(&payments).Error
	if err != nil {
		return nil, err
	}
	return payments, nil
}

// GetPaymentsByStatus returns payments by status with optional pagination
func (r *GormPaymentRepository) GetPaymentsByStatus(ctx context.Context, status string, limit, offset int) ([]*entities.PaymentRecord, error) {
	var payments []*entities.PaymentRecord
	query := r.db.WithContext(ctx).Where("status = ?", status)

	if limit > 0 {
		query = query.Limit(limit)
	}
	if offset > 0 {
		query = query.Offset(offset)
	}

	err := query.Order("created_at DESC").Find(&payments).Error
	if err != nil {
		return nil, err
	}
	return payments, nil
}

// GetPaymentStats returns payment statistics
func (r *GormPaymentRepository) GetPaymentStats(ctx context.Context, userID string) (*PaymentStats, error) {
	var stats PaymentStats

	// Get total payments count
	err := r.db.WithContext(ctx).Model(&entities.PaymentRecord{}).
		Where("user_id = ? AND status = ?", userID, PaymentStatusCompleted).
		Count(&stats.TotalPayments).Error
	if err != nil {
		return nil, err
	}

	// Get total amount spent
	var totalCents sql.NullInt64
	err = r.db.WithContext(ctx).Model(&entities.PaymentRecord{}).
		Where("user_id = ? AND status = ?", userID, PaymentStatusCompleted).
		Select("SUM(price_cents)").Scan(&totalCents).Error
	if err != nil {
		return nil, err
	}
	stats.TotalAmountCents = int(totalCents.Int64)

	// Get total tokens purchased
	var totalTokens sql.NullInt64
	err = r.db.WithContext(ctx).Model(&entities.PaymentRecord{}).
		Where("user_id = ? AND status = ?", userID, PaymentStatusCompleted).
		Select("SUM(token_count)").Scan(&totalTokens).Error
	if err != nil {
		return nil, err
	}
	stats.TotalTokens = int(totalTokens.Int64)

	return &stats, nil
}

// PaymentStats represents payment statistics
type PaymentStats struct {
	TotalPayments    int64 `json:"total_payments"`
	TotalAmountCents int   `json:"total_amount_cents"`
	TotalTokens      int   `json:"total_tokens"`
}
