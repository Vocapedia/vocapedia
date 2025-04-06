package entities

type User struct {
	Base
	Name      string    `json:"name"`
	Vocatoken string    `json:"vocatoken" gorm:"not null;unique"`
	Biography string    `json:"biography"`
	Username  string    `json:"username" gorm:"type:citext;not null;unique"`
	Email     string    `json:"-" gorm:"not null;unique"`
	Approved  bool      `json:"approved" gorm:""`
	Chapter   []Chapter `json:"chapters,omitempty" gorm:"foreignkey:creator_id"`
}
