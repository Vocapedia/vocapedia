package auth

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"text/template"
	"time"

	"github.com/go-chi/render"
	"gorm.io/gorm"

	"github.com/akifkadioglu/vocapedia/pkg/cache"
	"github.com/akifkadioglu/vocapedia/pkg/database"
	"github.com/akifkadioglu/vocapedia/pkg/entities"
	"github.com/akifkadioglu/vocapedia/pkg/i18n"
	"github.com/akifkadioglu/vocapedia/pkg/mail"
	"github.com/akifkadioglu/vocapedia/pkg/token"
	"github.com/akifkadioglu/vocapedia/utils"
)

func Token(w http.ResponseWriter, r *http.Request) {
	tmpl, err := os.ReadFile("pkg/template/auth.token.html")
	if err != nil {
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")

	fmt.Fprint(w, string(tmpl))
}

func SendOTP(w http.ResponseWriter, r *http.Request) {
	db := database.Manager()
	var user entities.User
	params := _login{}
	err := render.DecodeJSON(r.Body, &params)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{
			"error": i18n.Localizer(r, "error.something_went_wrong"),
		})
		return
	}

	if tx := db.First(&user, "email = ?", params.Email); tx.Error != nil && tx.Error != gorm.ErrRecordNotFound {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{
			"error": tx.Error.Error(),
		})
		return
	} else {
		user.Email = params.Email
		user.Username = strings.Split(params.Email, "@")[0] + utils.RandomString(5)
		db.Create(&user)
	}
	otp := generateOTP()

	err = cache.Redis().Set(r.Context(), user.Email, otp, 5*time.Minute).Err()
	if err != nil {
		panic(err)
	}

	tmpl, err := template.ParseFiles("pkg/mail/templates/auth.login.html")
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{
			"error": i18n.Localizer(r, "error.something_went_wrong"),
		})
		return
	}
	var body bytes.Buffer

	err = tmpl.Execute(&body, _emailData{
		Code:        otp,
		Header:      i18n.Localizer(r, "mail.auth.header"),
		Description: i18n.Localizer(r, "mail.auth.description"),
		Action:      i18n.Localizer(r, "mail.auth.action"),
		Warning:     i18n.Localizer(r, "mail.auth.warning"),
	})
	if err != nil {
		render.Status(r, http.StatusInternalServerError)

		render.JSON(w, r, map[string]string{
			"error": i18n.Localizer(r, "error.something_went_wrong"),
		})
		return
	}
	subject := i18n.Localizer(r, "mail.auth.login.subject")

	isSent, err := mail.Send(r, params.Email, subject, body.String())
	if err != nil {
		render.Status(r, http.StatusBadRequest)

		render.JSON(w, r, map[string]string{
			"error": i18n.Localizer(r, "error.something_went_wrong"),
		})
		return
	}

	render.JSON(w, r, map[string]any{
		"is_mail_sent": isSent,
		"message":      i18n.Localizer(r, "auth.otp.success"),
	})

}

func VerifyOTP(w http.ResponseWriter, r *http.Request) {
	params := _otp{}
	var tokenClaim entities.JwtModel
	var user entities.User
	db := database.Manager()
	err := render.DecodeJSON(r.Body, &params)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{
			"error": i18n.Localizer(r, "error.something_went_wrong"),
		})
		return
	}
	storedOTP, err := cache.Redis().Get(r.Context(), params.Email).Result()
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{
			"error": i18n.Localizer(r, "error.something_went_wrong"),
		})
		return
	}
	if storedOTP != params.OTP {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{
			"error": i18n.Localizer(r, "error.something_went_wrong"),
		})
		return
	}
	if err := db.Where("email = ?", params.Email).First(&user).Error; err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{
			"error": i18n.Localizer(r, "error.something_went_wrong"),
		})
		return
	}
	tokenClaim.UserID = user.ID
	tokenClaim.Username = user.Username
	tokenString, err := token.GenerateToken(tokenClaim)
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

func GetToken(w http.ResponseWriter, r *http.Request) {
	tokenClaim := entities.JwtModel{}
	tokenClaim.UserID, _ = strconv.ParseInt(r.URL.Query().Get("email"), 10, 64)
	tokenString, err := token.GenerateToken(tokenClaim)
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
