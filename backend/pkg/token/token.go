package token

import (
	"log"
	"net/http"

	"github.com/go-chi/jwtauth/v5"

	"github.com/akifkadioglu/vocapedia/pkg/config"
	"github.com/akifkadioglu/vocapedia/pkg/entities"
	"github.com/akifkadioglu/vocapedia/utils"
)

var tokenAuth *jwtauth.JWTAuth

func InitTokenAuth() {
	tokenAuth = jwtauth.New("HS256", []byte(config.ReadValue().JwtSecret), nil)
	if tokenAuth == nil {
		log.Println("Failed to initialize TokenAuth")
	} else {
		log.Println("Token auth is ready")
	}

}
func TokenAuth() *jwtauth.JWTAuth {
	return tokenAuth
}

func GenerateToken(model entities.JwtModel) (string, error) {
	mapData, err := utils.StructToMap(model)
	if err != nil {
		return "", err
	}
	_, token, err := tokenAuth.Encode(mapData)
	if err != nil {
		return "", err
	}

	return token, err
}

func User(r *http.Request) entities.JwtModel {
	_, claims, _ := jwtauth.FromContext(r.Context())

	var user entities.JwtModel
	utils.MapToStruct(claims, &user)

	return user
}
