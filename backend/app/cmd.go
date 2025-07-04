package app

import (
	"fmt"
	"net/http"

	"github.com/akifkadioglu/vocapedia/pkg/cache"
	"github.com/akifkadioglu/vocapedia/pkg/config"
	"github.com/akifkadioglu/vocapedia/pkg/database"
	"github.com/akifkadioglu/vocapedia/pkg/i18n"
	"github.com/akifkadioglu/vocapedia/pkg/mail"
	"github.com/akifkadioglu/vocapedia/pkg/payment"
	"github.com/akifkadioglu/vocapedia/pkg/search"
	"github.com/akifkadioglu/vocapedia/pkg/server"
	"github.com/akifkadioglu/vocapedia/pkg/snowflake"
	"github.com/akifkadioglu/vocapedia/pkg/token"
)

func Execute() {
	token.InitTokenAuth(
		config.ReadValue().JwtSecret,
	)
	i18n.InitI18n()
	mail.InitMail(
		config.ReadValue().SMTP.Host,
		config.ReadValue().SMTP.From,
		config.ReadValue().SMTP.Password,
		config.ReadValue().SMTP.Port,
	)
	snowflake.InitSnowflake()

	search.InitMeili(
		config.ReadValue().Meili.Host,
		config.ReadValue().Meili.APIKey,
		config.ReadValue().Meili.Index,
	)

	database.InitDB(
		config.ReadValue().Database.Host, config.ReadValue().Database.Port,
		config.ReadValue().Database.User, config.ReadValue().Database.Password,
		config.ReadValue().Database.Name, config.ReadValue().Database.SSLMode,
		config.ReadValue().AdminUsername,
		config.ReadValue().AdminEmail,
		config.ReadValue().AdminName,
		config.ReadValue().AdminBiography,
	)

	cache.InitRedis(
		config.ReadValue().Redis.Host,
		config.ReadValue().Redis.Port,
		config.ReadValue().Redis.Password,
		config.ReadValue().Redis.DB,
	)

	payment.InitializeProviders()

	s := server.CreateNewServer()

	s.MountHandlers(
		config.ReadValue().Host,
		config.ReadValue().Port,
		config.ReadValue().AllowMethods,
		config.ReadValue().AllowOrigins,
		config.ReadValue().AllowHeaders,
	)

	http.ListenAndServe(fmt.Sprintf(":%v", config.ReadValue().Port), s.Router)

}
