package user

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"strings"
	"time"

	"github.com/akifkadioglu/vocapedia/pkg/cache"
	"github.com/akifkadioglu/vocapedia/pkg/database"
	"github.com/akifkadioglu/vocapedia/pkg/entities"
	"github.com/akifkadioglu/vocapedia/pkg/i18n"
	"github.com/akifkadioglu/vocapedia/pkg/mail"
	"github.com/akifkadioglu/vocapedia/pkg/token"
	"github.com/akifkadioglu/vocapedia/pkg/utils"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func GetByUsername(w http.ResponseWriter, r *http.Request) {
	db := database.Manager()

	username := r.URL.Query().Get("username")

	if username == "" {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{
			"error": "Username query parameter is required",
		})
		return
	}

	var user entities.User
	err := db.Where("LOWER(username) = ?", strings.ToLower(username)).First(&user).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			render.Status(r, http.StatusNotFound)
			render.JSON(w, r, map[string]string{
				"error": "User not found",
			})
		} else {
			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, map[string]string{
				"error": "Database error",
			})
		}
		return
	}

	render.JSON(w, r, user)
}

func EditUser(w http.ResponseWriter, r *http.Request) {
	db := database.Manager()

	userID := token.User(r).UserID
	if userID == "" {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{
			"error": i18n.Localizer(r, "error.user_id_required"),
		})
		return
	}
	var user entities.User
	err := db.Where("id = ?", userID).First(&user).Error
	if err != nil {
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, map[string]string{
			"error": i18n.Localizer(r, "error.user_not_found"),
		})
		return
	}
	var updatedUser _updateUser
	if err := render.Decode(r, &updatedUser); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{
			"error": i18n.Localizer(r, "error.invalid_request_body"),
		})
		return
	}
	device, err := utils.StructToMap(updatedUser.Device)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{
			"error": i18n.Localizer(r, "error.something_went_wrong"),
		})
		return
	}

	deviceString := ""
	for _, k := range device {
		deviceString = deviceString + " | " + fmt.Sprintf("%v", k)
	}
	if err := db.Model(&user).Updates(map[string]any{
		"name":      updatedUser.Name,
		"username":  updatedUser.Username,
		"biography": updatedUser.Biography,
	}).Where("id = ?", userID).Error; err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{
			"error": i18n.Localizer(r, "error.failed_to_update_user"),
		})
		return
	}

	tmpl, err := template.ParseFiles("pkg/mail/templates/edit.user.html")
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{
			"error": i18n.Localizer(r, "error.something_went_wrong"),
		})
		return
	}
	var body bytes.Buffer

	err = tmpl.Execute(&body, _emailData{
		Header:      i18n.Localizer(r, "mail.edit.user.header"),
		Description: i18n.Localizer(r, "mail.edit.user.description"),
		Device:      deviceString,
	})
	if err != nil {
		render.Status(r, http.StatusInternalServerError)

		render.JSON(w, r, map[string]string{
			"error": i18n.Localizer(r, "error.something_went_wrong"),
		})
		return
	}
	subject := i18n.Localizer(r, "mail.edit.user.header")

	isSent, err := mail.Send(r, user.Email, subject, body.String())
	if err != nil && isSent {
		render.Status(r, http.StatusBadRequest)

		render.JSON(w, r, map[string]string{
			"error": i18n.Localizer(r, "error.something_went_wrong"),
		})
		return
	}
	jwtModel := entities.JwtModel{
		UserID:   userID,
		Username: user.Username,
		Device:   device,
	}

	tokenString, err := token.GenerateToken(jwtModel)
	if err != nil {
		render.Status(r, http.StatusBadRequest)

		render.JSON(w, r, map[string]string{
			"error": i18n.Localizer(r, "error.something_went_wrong"),
		})
		return
	}
	render.JSON(w, r, map[string]string{
		"token": tokenString,
	})
}

func Tokens(w http.ResponseWriter, r *http.Request) {
	db := database.Manager()
	var tokenEntity []entities.Token
	tokenStr := strings.Split(r.Header.Get("Authorization"), " ")[1]
	err := db.Where("token = ?", tokenStr).First(&tokenEntity).Error
	if err != nil {
		render.JSON(w, r, map[string][]entities.Token{
			"tokens": {},
		})
		return
	}
	userID := token.User(r).UserID

	err = db.Where("user_id = ?", userID).Find(&tokenEntity).Error
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{
			"error": i18n.Localizer(r, "error.tokens_fetch_failed"),
		})
		return
	}

	render.JSON(w, r, map[string][]entities.Token{
		"tokens": tokenEntity,
	})

}

func DeleteToken(w http.ResponseWriter, r *http.Request) {
	db := database.Manager()
	tokenIDStr := chi.URLParam(r, "id")

	userID := token.User(r).UserID
	var tokenEntity entities.Token
	if err := db.Where("id = ? AND user_id = ?", tokenIDStr, userID).First(&tokenEntity).Error; err != nil {
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, map[string]string{
			"error": i18n.Localizer(r, "error.token_not_found"),
		})
		return
	}

	if err := db.Delete(&tokenEntity).Error; err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{
			"error": i18n.Localizer(r, "error.something_went_wrong"),
		})
		return
	}

	render.JSON(w, r, map[string]string{
		"message": i18n.Localizer(r, "success.token_deleted"),
	})
}

func Check(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, map[string]string{})
}
func UpdateVocaToken(w http.ResponseWriter, r *http.Request) {
	db := database.Manager()
	userID := token.User(r).UserID
	vocatoken, err := utils.GenerateVocaToken(10)
	if err != nil {
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, map[string]string{
			"error": i18n.Localizer(r, "error.something_went_wrong"),
		})
		return
	}
	err = db.Model(&entities.User{}).
		Where("id = ?", userID).
		Update("vocatoken", vocatoken).Error
	if err != nil {
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, map[string]string{
			"error": i18n.Localizer(r, "error.something_went_wrong"),
		})
		return
	}
	render.JSON(w, r, map[string]string{
		"vocatoken": vocatoken,
	})
}

func GetVocaToken(w http.ResponseWriter, r *http.Request) {
	db := database.Manager()
	userID := token.User(r).UserID
	var user entities.User
	err := db.Where("id = ?", userID).First(&user).Error
	if err != nil {
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, map[string]string{
			"error": i18n.Localizer(r, "error.something_went_wrong"),
		})
		return
	}
	render.JSON(w, r, map[string]any{
		"vocatoken": user.Vocatoken,
	})
}

func DailyStreak(w http.ResponseWriter, r *http.Request) {
	rdb := cache.Redis()
	userID := token.User(r).UserID
	db := database.Manager()

	data, err := rdb.Get(r.Context(), fmt.Sprintf("streak:%s", userID)).Result()
	var streak _streak
	if err == redis.Nil {
		streak = _streak{Count: 0, LastDate: "", Rewarded: false}
	} else if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{
			"error": i18n.Localizer(r, "error.something_went_wrong"),
		})
		return
	} else {
		json.Unmarshal([]byte(data), &streak)
	}

	today := time.Now().Format("2006-01-02")
	yesterday := time.Now().AddDate(0, 0, -1).Format("2006-01-02")

	if streak.LastDate == today {
	} else if streak.LastDate == yesterday {
		streak.Count++
		if streak.Count > 7 {
			streak.Count = 1
			streak.Rewarded = false
		}

		streak.LastDate = today
	} else {
		streak.Count = 1
		streak.LastDate = today
		streak.Rewarded = false
	}

	if streak.Count == 7 && !streak.Rewarded {
		streak.Rewarded = true
		if err := db.Model(&entities.User{}).Where("id = ?", userID).Update("vocatoken_val", gorm.Expr("vocatoken_val + ?", 20)).Error; err != nil {
			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, map[string]string{
				"error": i18n.Localizer(r, "error.something_went_wrong"),
			})
			return
		}
	}

	streakJSON, _ := json.Marshal(streak)
	rdb.Set(r.Context(), fmt.Sprintf("streak:%s", userID), streakJSON, 0)

	render.JSON(w, r, map[string]any{
		"streak": streak,
	})
}
