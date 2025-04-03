package chapters

import (
	"fmt"
	"log"
	"math/rand/v2"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/akifkadioglu/vocapedia/pkg/database"
	"github.com/akifkadioglu/vocapedia/pkg/entities"
	"github.com/akifkadioglu/vocapedia/pkg/i18n"
	"github.com/akifkadioglu/vocapedia/pkg/token"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/xuri/excelize/v2"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func GetByID(w http.ResponseWriter, r *http.Request) {
	db := database.Manager()
	var chapters ChapterDTO
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return
	}
	userID := token.User(r).UserID
	tx := db.Model(&entities.Chapter{}).
		Select("chapters.*, COUNT(user_favorites.chapter_id) AS fav_count, (SELECT COUNT(*) FROM word_bases WHERE chapter_id = chapters.id) AS word_count, EXISTS(SELECT 1 FROM user_favorites WHERE user_favorites.chapter_id = chapters.id AND user_favorites.user_id = ? LIMIT 1) AS is_favorited", userID).
		Joins("LEFT JOIN user_favorites ON user_favorites.chapter_id = chapters.id").
		Where("chapters.id = ?", id).
		Group("chapters.id").
		Preload("Creator").
		Preload("WordBase.Word").
		First(&chapters)
	if tx.Error != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{
			"error": i18n.Localizer(r, "error.something_went_wrong"),
		})
		return
	}
	render.JSON(w, r, map[string]any{
		"chapter": chapters,
	})
}

func DeleteFavorite(w http.ResponseWriter, r *http.Request) {
	db := database.Manager()
	snowflakeID, err := strconv.ParseInt(r.URL.Query().Get("chapter_id"), 10, 64)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{
			"error": i18n.Localizer(r, "error.something_went_wrong"),
		})
		return
	}
	tx := db.Unscoped().Delete(&entities.UserFavorite{}, "user_id = ? AND chapter_id = ?", token.User(r).UserID, snowflakeID)
	if tx.Error != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{
			"error": i18n.Localizer(r, "error.something_went_wrong"),
		})
		return
	}
	render.JSON(w, r, map[string]string{
		"message": i18n.Localizer(r, "ok.success"),
	})
}

func Favorites(w http.ResponseWriter, r *http.Request) {
	db := database.Manager()
	userID := token.User(r).UserID
	var favorites []ChapterDTO
	tx := db.Table("user_favorites").
		Joins("JOIN chapters ON chapters.id = user_favorites.chapter_id").
		Joins("JOIN word_bases ON word_bases.chapter_id = chapters.id").
		Where("user_favorites.user_id = ?", userID).
		Unscoped().
		Select("chapters.*, user_favorites.chapter_id as fav_chapter_id, COUNT(DISTINCT user_favorites.chapter_id) as fav_count, COUNT(DISTINCT word_bases.id) as word_count, EXISTS(SELECT 1 FROM user_favorites WHERE user_favorites.chapter_id = chapters.id AND user_favorites.user_id = ? LIMIT 1) AS is_favorited", userID).
		Preload("Creator").
		Group("chapters.id, user_favorites.chapter_id").
		Order("id desc").
		Find(&favorites)

	if tx.Error != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	render.JSON(w, r, map[string]any{
		"list": favorites,
	})

}
func Favorite(w http.ResponseWriter, r *http.Request) {
	db := database.Manager()
	if token.User(r).UserID == "" {
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	snowflakeID, err := strconv.ParseInt(r.URL.Query().Get("chapter_id"), 10, 64)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{
			"error": i18n.Localizer(r, "error.id_convert"),
		})
		return
	}
	userID, _ := strconv.ParseInt(token.User(r).UserID, 10, 64)
	tx := db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "user_id"}, {Name: "chapter_id"}},
		DoNothing: true,
	}).Create(&entities.UserFavorite{
		UserID:    userID,
		ChapterID: snowflakeID,
	})

	if tx.Error != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{
			"error": i18n.Localizer(r, "error.something_went_wrong"),
		})
		return
	}

	render.JSON(w, r, map[string]string{
		"message": i18n.Localizer(r, "ok.success"),
	})
}

func GetTrendingChapters(w http.ResponseWriter, r *http.Request) {
	var chapters []entities.Chapter
	db := database.Manager()
	sevenDaysAgo := time.Now().AddDate(0, 0, -7)

	err := db.
		Joins("JOIN user_favorites ON user_favorites.chapter_id = chapters.id").
		Where("user_favorites.created_at >= ?", sevenDaysAgo).
		Group("chapters.id").
		Order("COUNT(user_favorites.chapter_id) DESC, SUM(EXTRACT(EPOCH FROM NOW() - user_favorites.created_at)) DESC").
		Limit(10).
		Find(&chapters).Error
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, err)
		return
	}

	render.JSON(w, r, map[string]any{
		"trends": chapters,
	})
}

func GameFormat(w http.ResponseWriter, r *http.Request) {
	db := database.Manager()

	id := chi.URLParam(r, "id")
	lang := r.URL.Query().Get("lang")
	targetLang := r.URL.Query().Get("target_lang")
	limitStr := r.URL.Query().Get("limit")

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 10 {
		limit = 10
	}

	var wordBases []entities.WordBase
	err = db.Preload("Word").
		Where("chapter_id = ?", id).
		Order("RANDOM()").
		Limit(limit).
		Find(&wordBases).Error

	if err != nil || len(wordBases) == 0 {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{
			"error": "No word bases found for chapter ID: " + id,
		})
		return
	}

	var targetLangWords []entities.Word
	err = db.Table("words").
		Where("lang = ?", targetLang).
		Where("chapter_id = ?", id).
		Order("RANDOM()").
		Limit(limit * 3).
		Find(&targetLangWords).Error

	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{
			"error": "Failed to fetch target language words",
		})
		return
	}

	response := map[string]any{
		"questions": []map[string]any{},
	}
	for _, wb := range wordBases {
		var correctWord string
		var correctDescription string
		var correctAnswer string
		for _, word := range wb.Word {
			if word.Lang == lang {
				correctWord = word.Word
				correctDescription = word.Description
			}
			if word.Lang == targetLang {
				correctAnswer = word.Word
			}

			if correctWord != "" && correctAnswer != "" {
				break
			}
		}
		if correctWord == "" || correctAnswer == "" {
			continue
		}

		var wrongWords []string
		for _, word := range targetLangWords {
			if word.Word != correctWord {
				wrongWords = append(wrongWords, word.Word)
			}
		}

		if len(wrongWords) < 4 {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, map[string]string{
				"error": "Not enough wrong answers to generate questions.",
			})
			return
		}

		var wrongChoices []string
		for _, word := range wrongWords {
			if word != correctAnswer {
				log.Println(wrongChoices, word)

				wrongChoices = append(wrongChoices, word)
			}
		}
		rand.Shuffle(len(wrongChoices), func(i, j int) { wrongChoices[i], wrongChoices[j] = wrongChoices[j], wrongChoices[i] })
		wrongChoices = wrongChoices[:3]

		choices := append(wrongChoices, correctAnswer)
		rand.Shuffle(len(choices), func(i, j int) { choices[i], choices[j] = choices[j], choices[i] })

		correctIndex := indexOf(choices, correctAnswer)

		question := map[string]any{
			"id":           wb.ID,
			"word":         correctWord,
			"description":  correctDescription,
			"choices":      choices,
			"correctIndex": correctIndex,
		}

		response["questions"] = append(response["questions"].([]map[string]any), question)
	}

	render.JSON(w, r, response)
}

func Search(w http.ResponseWriter, r *http.Request) {
	db := database.Manager()
	query := r.URL.Query().Get("q")
	pageStr := r.URL.Query().Get("page")
	limitStr := r.URL.Query().Get("limit")

	page := 1
	limit := 10

	if p, err := strconv.Atoi(pageStr); err == nil && p > 0 {
		page = p
	}
	if l, err := strconv.Atoi(limitStr); err == nil && l > 0 {
		limit = l
	}
	offset := (page - 1) * limit
	var results []ChapterDTO
	userID := token.User(r).UserID
	tx := db.Model(&entities.Chapter{}).Raw(`
WITH ranked_chapters AS (
        SELECT *,
            (0.7 * similarity(title, ?) + 0.3 * similarity(description, ?)) AS sim_score,
            LEAST(levenshtein(title, ?), levenshtein(description, ?)) AS lev_score
        FROM chapters
        WHERE title % ? OR description % ?
    )
    SELECT rc.*,
           (SELECT COUNT(*) FROM user_favorites WHERE chapter_id = rc.id) AS fav_count,
           (SELECT COUNT(*) FROM word_bases WHERE chapter_id = rc.id) AS word_count,
		   EXISTS(SELECT 1 FROM user_favorites WHERE user_favorites.chapter_id = rc.id AND user_favorites.user_id = ? LIMIT 1) AS is_favorited
    FROM ranked_chapters rc
    ORDER BY sim_score DESC, lev_score ASC
    LIMIT ? OFFSET ?
`, query, query, query, query, query, query, userID, limit, offset)
	err := tx.Preload("Creator").Order("id desc").Find(&results).Error

	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{
			"error": err.Error(),
		})
		return
	}
	render.JSON(w, r, map[string]any{
		"page":  page,
		"limit": limit,
		"list":  results,
	})
}

func ComposeByExcel(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "File upload failed", http.StatusBadRequest)
		return
	}
	defer file.Close()

	f, err := excelize.OpenReader(file)
	if err != nil {
		http.Error(w, "Failed to read Excel file", http.StatusInternalServerError)
		return
	}
	rows, err := f.Rows("Sheet1")
	if err != nil {
		http.Error(w, "Failed to read rows from sheet", http.StatusInternalServerError)
		return
	}
	log.Println(rows)

	w.Write([]byte("Chapters uploaded successfully"))

}
func Compose(w http.ResponseWriter, r *http.Request) {
	var params _compose
	var chapter entities.Chapter
	db := database.Manager()

	err := render.DecodeJSON(r.Body, &params)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{
			"error": i18n.Localizer(r, "error.something_went_wrong"),
		})
		return
	}
	userID, err := strconv.ParseInt(token.User(r).UserID, 10, 64)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{
			"error": i18n.Localizer(r, "error.something_went_wrong"),
		})
		return
	}
	err = db.Transaction(func(tx *gorm.DB) error {
		var err error
		chapter.CreatorID = userID
		chapter.Description = params.Description
		chapter.Lang = params.Lang
		chapter.TargetLang = params.TargetLang
		chapter.Title = params.Title
		chapter.Tutorial = params.Tutorial
		if err := tx.Create(&chapter).Error; err != nil {
			return err
		}

		for _, wb := range params.WordBase {
			if len(wb.Word) < 2 {
				continue
			}
			var termWB entities.WordBase
			termWB.ChapterID = chapter.ID
			termWB.Type = wb.Type
			if err := tx.Create(&termWB).Error; err != nil {
				return err
			}

			for _, word := range wb.Word {
				var termW entities.Word
				termW.ChapterID = chapter.ID
				termW.Description = word.Description
				if word.Examples != "" {
					termW.Examples = []string{word.Examples}
				}
				termW.Lang = word.Lang
				termW.Word = word.Word
				termW.WordBaseID = termWB.ID

				if err := tx.Create(&termW).Error; err != nil {
					return err
				}
			}
		}
		return err
	})
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{
			"error": i18n.Localizer(r, "error.something_went_wrong"),
		})
		return
	}
	render.JSON(w, r, map[string]string{
		"chapter_id": fmt.Sprintf("%v", chapter.ID),
	})

	log.Println(params)
}
func UserChapters(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	db := database.Manager()
	var chapters []ChapterDTO
	userID := token.User(r).UserID
	tx := db.Model(&entities.Chapter{}).
		Select("chapters.*,(SELECT COUNT(*) FROM user_favorites WHERE user_favorites.chapter_id = chapters.id) AS fav_count, (SELECT COUNT(*) FROM word_bases WHERE word_bases.chapter_id = chapters.id) AS word_count, creator.username, EXISTS(SELECT 1 FROM user_favorites WHERE user_favorites.chapter_id = chapters.id AND user_favorites.user_id = ? LIMIT 1) AS is_favorited", userID).
		Joins("LEFT JOIN users AS creator ON creator.id = chapters.creator_id").
		Where("LOWER(creator.username) = ?", strings.ToLower(username)).
		Group("chapters.id, creator.username").
		Preload("Creator").
		Order("id desc").
		Find(&chapters)

	if tx.Error != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]any{
			"error": i18n.Localizer(r, "error.something_went_wrong"),
			"list":  []ChapterDTO{},
		})
		return
	}
	render.JSON(w, r, map[string]any{
		"list": chapters,
	})
}
