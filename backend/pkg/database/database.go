package database

import (
	"fmt"
	"log"

	"gorm.io/gorm"

	"github.com/akifkadioglu/vocapedia/pkg/entities"
	"github.com/brianvoe/gofakeit"
)

var (
	db  *gorm.DB
	err error
)

func InitDB(host string, port int, user, password, dbname string) {
	db, err = pg(host, port, user, password, dbname)
	if err != nil {
		panic(err)
	}
	if err := db.Exec(`CREATE EXTENSION IF NOT EXISTS pg_trgm`).Error; err != nil {
		panic(err)
	}

	db.AutoMigrate(
		&entities.User{},
		&entities.Chapter{},
		&entities.WordBase{},
		&entities.Word{},
	)
	//seed()
	db.Exec(`CREATE INDEX IF NOT EXISTS idx_trgm_title ON chapters USING gin (title gin_trgm_ops)`)
	db.Exec(`CREATE INDEX IF NOT EXISTS idx_trgm_description ON chapters USING gin (description gin_trgm_ops)`)
	db.Exec(`CREATE EXTENSION IF NOT EXISTS fuzzystrmatch`)

	log.Println("Database is ready")
}
func Manager() *gorm.DB {
	return db
}
func seed() {
	var user entities.User
	gofakeit.Struct(&user)
	db.Create(&user)
	for range 60 {
		chapter := entities.Chapter{
			UserID: user.ID,
			Lang:   "en",
			TargetLang: "tr",
			Title:       gofakeit.Sentence(3),
			Description: gofakeit.Paragraph(1, 2, 3, " "),
		}
		fmt.Println(chapter)
		if err := db.Create(&chapter).Error; err != nil {
			fmt.Println("veri eklerken hata:", err)
		}
		for range 15 {
			wb := entities.WordBase{
				ChapterID: chapter.ID,
			}
			if err := db.Create(&wb).Error; err != nil {
				fmt.Println("veri eklerken hata:", err)
			}
			for i := range 2 {
				langTerm := "en"
				if i == 1 {
					langTerm = "tr"
				}
				word := entities.Word{
					Lang:        langTerm,
					Word:        gofakeit.Word(),
					Type:        gofakeit.Word(),
					Description: gofakeit.Paragraph(1, 2, 3, " "),
					Examples:    []string{gofakeit.Sentence(3)},
					WordBaseID:  wb.ID,
				}
				if err := db.Create(&word).Error; err != nil {
					fmt.Println("veri eklerken hata:", err)
				}
			}
		}
	}
}
