package entities

import (
	"time"
)

type Stream struct {
	Base
	RoomID      string `json:"room_id" gorm:"uniqueIndex"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Password    string `json:"password"`
	CreatedBy   int64  `json:"created_by"`
	IsActive    bool   `json:"is_active" gorm:"default:true"`
	// Scheduled lesson time
	ScheduledAt time.Time `json:"scheduled_at"`
	// Source and target languages
	Lang       string `json:"lang"`
	TargetLang string `json:"target_lang"`
	// Duration in minutes
	Duration int `json:"duration"`
	// Maximum number of participants
	MaxParticipants int `json:"max_participants"`
	// Actual start and end times
	StartTime *time.Time `json:"start_time,omitempty"`
	EndTime   *time.Time `json:"end_time,omitempty"`

	// Relationships
	Creator User `json:"creator,omitempty" gorm:"foreignKey:CreatedBy"`

	// Virtual fields for API response
	Participants int `json:"participants" gorm:"-"`
}
