package chapters

import (
	"net/http"
	"strconv"
	"time"

	"github.com/akifkadioglu/vocapedia/pkg/database"
	"github.com/akifkadioglu/vocapedia/pkg/entities"
	"github.com/akifkadioglu/vocapedia/pkg/i18n"
	"github.com/akifkadioglu/vocapedia/pkg/token"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	"gorm.io/gorm/clause"
)

func GetByID(w http.ResponseWriter, r *http.Request) {
	db := database.Manager()
	var chapters entities.Chapter
	id := chi.URLParam(r, "id")

	db.Where("id = ?", id).Preload("Creator").Preload("WordBase.Word").First(&chapters)
	render.JSON(w, r, map[string]any{
		"chapter": chapters,
	})
}

func DeleteFavorite(w http.ResponseWriter, r *http.Request) {
	db := database.Manager()
	tx := db.Unscoped().Delete(&entities.UserFavorite{}, "user_id = ? AND chapter_id = ?", token.User(r).UserID, uuid.MustParse(r.URL.Query().Get("chapter_id")))
	if tx.Error != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{
			"error": i18n.Localizer(r, "error.something_went_wrong"),
		})
		return
	}

	render.JSON(w, r, i18n.Localizer(r, "ok.success"))
}

func Favorite(w http.ResponseWriter, r *http.Request) {
	db := database.Manager()

	tx := db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "user_id"}, {Name: "chapter_id"}},
		DoNothing: true,
	}).Create(&entities.UserFavorite{
		UserID:    token.User(r).UserID,
		ChapterID: uuid.MustParse(r.URL.Query().Get("chapter_id")),
	})

	if tx.Error != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{
			"error": i18n.Localizer(r, "error.something_went_wrong"),
		})
		return
	}
	render.JSON(w, r, i18n.Localizer(r, "ok.success"))
}

func GetTrendingChapters(w http.ResponseWriter, r *http.Request) {
	var chapters []entities.Chapter
	db := database.Manager()
	sevenDaysAgo := time.Now().AddDate(0, 0, -7)

	err := db.
		Joins("JOIN user_favorites ON user_favorites.chapter_id = chapters.id::text").
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
	var wordBases []entities.WordBase
	var words []entities.Word
	var format _gameFormat
	db := database.Manager()
	id := chi.URLParam(r, "id")

	limitStr := r.URL.Query().Get("limit")

	limit := 2

	if l, err := strconv.Atoi(limitStr); err == nil && l > 0 {
		limit = l
	}

	if err := db.Where("chapter_id = ?", id).Preload("Word").Order("RANDOM()").Limit(limit).Find(&wordBases).Error; err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{
			"error": i18n.Localizer(r, "error.something_went_wrong"),
		})
		return
	}
	answers := []string{}
	for _, v := range wordBases {
		answers = append(answers, v.Word[1].Word)
		format.Matchs = append(format.Matchs, _match{
			Word:   v.Word[0].Word,
			Answer: v.Word[1].Word,
		})
	}

	if err := db.Where("chapter_id = ?", id).
		Not("word IN ?", answers).
		Order("RANDOM()").
		Limit(50).
		Find(&words).Error; err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{
			"error": i18n.Localizer(r, "error.something_went_wrong"),
		})
		return
	}
	for _, v := range words {
		format.Answers = append(format.Answers, v.Word)
	}

	render.JSON(w, r, format)
}
