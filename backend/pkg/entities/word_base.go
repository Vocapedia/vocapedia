package entities

type WordBase struct {
	Base
	ChapterID int64    `json:"chapter_id" gorm:"not null"`
	Chapter   *Chapter `json:"chapter,omitempty" gorm:"foreignKey:ChapterID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Type      string   `json:"type" gorm:"not null"`
	Word      []Word   `json:"words,omitempty"`
}
