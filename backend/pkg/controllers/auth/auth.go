package auth

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/render"

	"github.com/akifkadioglu/vocapedia/pkg/database"
	"github.com/akifkadioglu/vocapedia/pkg/entities"
	"github.com/akifkadioglu/vocapedia/pkg/mail"
	"github.com/akifkadioglu/vocapedia/pkg/token"
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
	params := loginBody{}
	err := render.DecodeJSON(r.Body, &params)
	if err != nil {
		render.JSON(w, r, map[string]string{
			"error": "something went wrong",
		})
		return
	}

	if tx := db.First(&user, "email = ?", params.Email); tx.Error != nil {
		render.JSON(w, r, map[string]string{
			"error": "user does not exist",
		})
		return
	}
	tokenClaim.UserID = user.ID
	tokenString, err := token.GenerateToken(tokenClaim)
	if err != nil {
		render.JSON(w, r, map[string]string{
			"error": "something went wrong",
		})
		return
	}
	err = mail.Mail(params.Email, "login", tokenString)

	if err != nil {
		render.JSON(w, r, map[string]string{
			"error": "something went wrong",
		})
		return
	}

}
