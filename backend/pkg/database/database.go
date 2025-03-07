package database

import (
	"fmt"

	"github.com/brianvoe/gofakeit/v6"
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
	db.Exec(`CREATE INDEX IF NOT EXISTS idx_trgm_name ON chapters USING gin (name gin_trgm_ops)`)
	db.Exec(`CREATE INDEX IF NOT EXISTS idx_trgm_description ON chapters USING gin (description gin_trgm_ops)`)

	for i := 0; i < 60; i++ {
		chapter := entities.Chapter{
			Title:       gofakeit.Sentence(3),             // 3 kelimelik rastgele başlık
			Description: gofakeit.Paragraph(1, 2, 3, " "), // 1-2 cümle açıklama
		}

		if err := db.Create(&chapter).Error; err != nil {
			fmt.Println("veri eklerken hata: %v", err)
		}
	}
	fmt.Println("Database is ready")
}
func Manager() *gorm.DB {
	return db
}
