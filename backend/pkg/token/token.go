package token

import (
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/jwtauth/v5"

	"github.com/akifkadioglu/vocapedia/pkg/config"
	"github.com/akifkadioglu/vocapedia/pkg/database"
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
	db := database.Manager()
	var tokenEntity entities.Token
	mapData, err := utils.StructToMap(model)
	if err != nil {
		return "", err
	}
	_, token, err := tokenAuth.Encode(mapData)
	if err != nil {
		return "", err
	}
	tokenEntity.Token = token
	tokenEntity.UserID, _ = strconv.ParseInt(model.UserID, 10, 64)
	tokenEntity.Device = model.Device

	err = db.Create(&tokenEntity).Error
	if err != nil {
		return "", err
	}

	return token, err
}

func User(r *http.Request) entities.JwtModel {
	_, claims, _ := jwtauth.FromContext(r.Context())
	var user entities.JwtModel
	utils.MapToStruct(claims, &user)
	if user.UserID == "" {
		user.UserID = "0"
	}
	return user
}
