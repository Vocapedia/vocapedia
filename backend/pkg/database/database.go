package database

import (
	"fmt"
	"log"

	"gorm.io/gorm"

	"github.com/akifkadioglu/vocapedia/pkg/entities"
	"github.com/akifkadioglu/vocapedia/pkg/utils"
	"github.com/brianvoe/gofakeit"
)

var (
	db  *gorm.DB
	err error
)

func InitDB(host string, port int, user, password, sslMode, dbname, adminusername, adminemail, adminname, adminbiography string) {
	db, err = pg(host, port, user, password, dbname, sslMode)
	if err != nil {
		panic(err)
	}
	if err := db.Exec(`CREATE EXTENSION IF NOT EXISTS pg_trgm`).Error; err != nil {
		panic(err)
	}
	if err := db.Exec(`CREATE EXTENSION IF NOT EXISTS citext`).Error; err != nil {
		panic(err)
	}
	if err := db.Exec(`CREATE EXTENSION IF NOT EXISTS pgroonga`).Error; err != nil {
		panic(err)
	}

	db.AutoMigrate(
		&entities.User{},
		&entities.Chapter{},
		&entities.UserFavorite{},
		&entities.WordBase{},
		&entities.Word{},
		&entities.Token{},
	)

	db.Exec(`CREATE INDEX IF NOT EXISTS idx_trgm_title ON chapters USING gin (title gin_trgm_ops)`)
	db.Exec(`CREATE INDEX IF NOT EXISTS idx_trgm_description ON chapters USING gin (description gin_trgm_ops)`)
	db.Exec(`CREATE EXTENSION IF NOT EXISTS fuzzystrmatch`)
	db.Exec("CREATE UNIQUE INDEX IF NOT EXISTS unique_user_chapter ON user_favorites (user_id, chapter_id);")
	db.Exec(`CREATE INDEX IF NOT EXISTS idx_pgroonga_title ON chapters USING pgroonga (title)`)
	db.Exec(`CREATE INDEX IF NOT EXISTS idx_pgroonga_description ON chapters USING pgroonga (description)`)

	//seed()

	createAdmin(adminusername, adminemail, adminname, adminbiography)

	log.Println("Database is ready with PGroonga")
}
func Manager() *gorm.DB {
	return db
}

func createAdmin(username, email, name, biography string) {
	vocatoken, err := utils.GenerateVocaToken(10)
	if err != nil {
		fmt.Println("veri eklerken hata:", err)
	}
	user := entities.User{
		Vocatoken: vocatoken,
		Username:  username,
		Email:     email,
		Name:      name,
		Biography: biography,
		Approved:  true,
	}
	if err := db.Save(&user).Error; err != nil {
		fmt.Println("veri eklerken hata:", err)
	}
}
func seed(username, email, name string) {
	user := entities.User{
		Username: username,
		Email:    email,
		Name:     name,
	}

	db.FirstOrCreate(&user, entities.User{Username: user.Username})
	for range 50 {
		chapter := entities.Chapter{
			CreatorID:   user.ID,
			Lang:        "en",
			TargetLang:  "tr",
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
				Type:      "noun",
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
					ChapterID:   chapter.ID,
					Lang:        langTerm,
					Word:        gofakeit.Word(),
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
