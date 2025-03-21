package token

import (
	"github.com/go-chi/jwtauth/v5"

	"github.com/akifkadioglu/vocapedia/pkg/config"
	"github.com/akifkadioglu/vocapedia/pkg/entities"
	"github.com/akifkadioglu/vocapedia/utils"
)

var _tokenAuth *jwtauth.JWTAuth

func InitTokenAuth() {

	_tokenAuth = jwtauth.New("HS256", []byte(config.ReadValue().JwtSecret), nil)
}
func TokenAuth() *jwtauth.JWTAuth {
	return _tokenAuth
}

func GenerateToken(model entities.JwtModel) (string, error) {
	mapData, err := utils.StructToMap(model)
	
	if err != nil {
		return "", err
	}

	_, token, err := _tokenAuth.Encode(mapData)

	if err != nil {
		return "", err
	}

	return token, err
}
