package auth

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"strings"
	"text/template"

	"github.com/go-chi/render"
	"github.com/google/uuid"
	"gorm.io/gorm"

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

func Login(w http.ResponseWriter, r *http.Request) {
	db := database.Manager()
	var user entities.User
	var tokenClaim entities.JwtModel
	params := _login{}
	err := render.DecodeJSON(r.Body, &params)
	if err != nil {
		render.JSON(w, r, map[string]string{
			"error": i18n.Localizer(r, "error.something_went_wrong"),
		})
		return
	}

	if tx := db.First(&user, "email = ?", params.Email); tx.Error != gorm.ErrRecordNotFound {
		render.JSON(w, r, map[string]string{
			"error": "user does not exist",
		})
		return
	} else {
		user.Email = params.Email
		user.Username = strings.Split(params.Email, "@")[0] + utils.RandomString(5)
		db.Create(&user)
	}
	tokenClaim.UserID = user.ID
	tokenString, err := token.GenerateToken(tokenClaim)
	if err != nil {
		render.JSON(w, r, map[string]string{
			"error": i18n.Localizer(r, "error.something_went_wrong"),
		})
		return
	}
	tmpl, err := template.ParseFiles("pkg/mail/template/auth.login.html")
	if err != nil {
		render.JSON(w, r, map[string]string{
			"error": i18n.Localizer(r, "error.something_went_wrong"),
		})
		return
	}
	var body bytes.Buffer

	err = tmpl.Execute(&body, _emailData{
		Code:        tokenString,
		Header:      i18n.Localizer(r, "mail.auth.header"),
		Description: i18n.Localizer(r, "mail.auth.description"),
		Action:      i18n.Localizer(r, "mail.auth.action"),
		Warning:     i18n.Localizer(r, "mail.auth.warning"),
	})
	if err != nil {
		render.JSON(w, r, map[string]string{
			"error": i18n.Localizer(r, "error.something_went_wrong"),
		})
		return
	}
	msg := []byte("Subject: Giriş Kodunuz\nMIME-Version: 1.0\nContent-Type: text/html; charset=UTF-8\n\n" + body.String())

	err = mail.Send(params.Email, msg, body.String())
	if err != nil {
		render.JSON(w, r, map[string]string{
			"error": i18n.Localizer(r, "error.something_went_wrong"),
		})
		return
	}

}

func GetToken(w http.ResponseWriter, r *http.Request) {
	tokenClaim := entities.JwtModel{}
	tokenClaim.UserID = uuid.New()
	tokenString, err := token.GenerateToken(tokenClaim)
	if err != nil {
		render.JSON(w, r, map[string]string{
			"error": i18n.Localizer(r, "error.something_went_wrong"),
		})
		return
	}
	render.JSON(w, r, map[string]string{
		"token": tokenString,
	})
}
