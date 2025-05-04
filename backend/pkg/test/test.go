package test

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/akifkadioglu/vocapedia/pkg/cache"
	"github.com/akifkadioglu/vocapedia/pkg/config"
	"github.com/akifkadioglu/vocapedia/pkg/database"
	"github.com/akifkadioglu/vocapedia/pkg/i18n"
	"github.com/akifkadioglu/vocapedia/pkg/mail"
	"github.com/akifkadioglu/vocapedia/pkg/search"
	"github.com/akifkadioglu/vocapedia/pkg/server"
	"github.com/akifkadioglu/vocapedia/pkg/snowflake"
	"github.com/akifkadioglu/vocapedia/pkg/token"
)

func TestInit() *server.Server {
	token.InitTokenAuth(
		config.ReadTestValue().JwtSecret,
	)
	i18n.InitI18n()
	mail.InitMail(
		config.ReadTestValue().SMTP.Host,
		config.ReadTestValue().SMTP.From,
		config.ReadTestValue().SMTP.Password,
		config.ReadTestValue().SMTP.Port,
	)
	snowflake.InitSnowflake()
	database.InitDB(
		config.ReadTestValue().Database.Host, config.ReadTestValue().Database.Port,
		config.ReadTestValue().Database.User, config.ReadTestValue().Database.Password,
		config.ReadTestValue().Database.Name,
		config.ReadTestValue().AdminUsername,
		config.ReadTestValue().AdminEmail,
		config.ReadTestValue().AdminName,
		config.ReadTestValue().AdminBiography,
	)
	search.InitMeili(
		config.ReadTestValue().Meili.Host,
		config.ReadTestValue().Meili.APIKey,
		config.ReadTestValue().Meili.Index,
	)
	cache.InitRedis(
		config.ReadTestValue().Redis.Host,
		config.ReadTestValue().Redis.Port,
		config.ReadTestValue().Redis.Password,
		config.ReadTestValue().Redis.DB,
	)
	s := server.CreateNewServer()
	s.MountHandlers(
		config.ReadTestValue().Host,
		config.ReadTestValue().Port,
		config.ReadTestValue().AllowMethods,
		config.ReadTestValue().AllowOrigins,
		config.ReadTestValue().AllowHeaders,
	)
	return s
}

func ExecuteRequest(req *http.Request, s *server.Server) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	s.Router.ServeHTTP(rr, req)

	return rr
}

func CheckResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func ConvertBody(obj any) (*bytes.Reader, error) {
	jsonData, err := json.Marshal(obj)
	if err != nil {
		log.Println("Error marshalling JSON:", err)
		return nil, err
	}
	body := bytes.NewReader(jsonData)
	return body, nil
}

func ConvertBodyToStruct(body *bytes.Buffer, obj any) error {
	err := json.NewDecoder(body).Decode(obj)
	if err != nil {
		log.Println("Error decoding JSON:", err)
	}
	return err
}
