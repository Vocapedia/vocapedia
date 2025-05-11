package token

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"hash"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/go-chi/jwtauth/v5"

	"github.com/akifkadioglu/vocapedia/pkg/database"
	"github.com/akifkadioglu/vocapedia/pkg/entities"
	"github.com/akifkadioglu/vocapedia/pkg/utils"
)

var tokenAuth *jwtauth.JWTAuth
var hmacProvider hash.Hash

func InitTokenAuth(secret string) {
	tokenAuth = jwtauth.New("HS256", []byte(secret), nil)
	if tokenAuth == nil {
		log.Println("Failed to initialize TokenAuth")
	} else {
		log.Println("Token auth is ready")
	}
	hmacProvider = hmac.New(sha256.New, []byte(secret))
}
func TokenAuth() *jwtauth.JWTAuth {
	return tokenAuth
}

func GenerateDeterministicToken(input string, length int) string {
	hmacProvider.Write([]byte(input))
	hash := hmacProvider.Sum(nil)
	encoded := base64.URLEncoding.EncodeToString(hash)
	encoded = strings.ToLower(strings.TrimRight(encoded, "="))
	if len(encoded) > length {
		return encoded[:length]
	}
	return encoded
}

func GenerateJitsiToken(user entities.JwtModel, room, domain, appID string) (string, error) {
	claims := map[string]any{
		"aud":  "jitsi",
		"iss":  appID,
		"sub":  domain,
		"room": room,
		"exp":  time.Now().Add(1 * time.Hour).Unix(),
		"context": map[string]interface{}{
			"user": map[string]string{
				"name":    user.Username,
				"user_id": user.UserID,
			},
		},
	}

	_, token, err := tokenAuth.Encode(claims)
	if err != nil {
		return "", err
	}

	return token, nil
}

func GenerateToken(model entities.JwtModel) (string, error) {
	db := database.Manager()
	var tokenEntity entities.Token

	userID, err := strconv.ParseInt(model.UserID, 10, 64)
	if err != nil {
		return "", err
	}

	mapData, err := utils.StructToMap(model)
	if err != nil {
		return "", err
	}
	_, token, err := tokenAuth.Encode(mapData)
	if err != nil {
		return "", err
	}
	if model.UserID == "0" {
		tokenEntity.Token = token
		tokenEntity.UserID, _ = strconv.ParseInt(model.UserID, 10, 64)
		tokenEntity.Device = model.Device
		err = db.Create(&tokenEntity).Error
	} else {

		err = db.Where(entities.Token{UserID: userID, Device: model.Device}).
			Assign(entities.Token{Token: token}).
			FirstOrCreate(&tokenEntity).Error

	}
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
