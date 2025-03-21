package database

import (
	"log"

	"gorm.io/gorm"

	"github.com/akifkadioglu/vocapedia/pkg/entities"
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

	db.AutoMigrate(&entities.User{}, &entities.Chapter{})
	db.Exec(`CREATE INDEX IF NOT EXISTS idx_trgm_title ON chapters USING gin (title gin_trgm_ops)`)
	db.Exec(`CREATE INDEX IF NOT EXISTS idx_trgm_description ON chapters USING gin (description gin_trgm_ops)`)
	db.Exec(`CREATE EXTENSION IF NOT EXISTS fuzzystrmatch`)

	/* for range 60 {
		chapter := entities.Chapter{
			Title:       gofakeit.Sentence(3),
			Description: gofakeit.Paragraph(1, 2, 3, " "),
		}
		fmt.Println(chapter)
		if err := db.Create(&chapter).Error; err != nil {
			fmt.Println("veri eklerken hata: %v", err)
		}
	} */

	log.Println("Database is ready")
}
func Manager() *gorm.DB {
	return db
}
