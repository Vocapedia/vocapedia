package app

import (
	"github.com/akifkadioglu/vocapedia/pkg/config"
	"github.com/akifkadioglu/vocapedia/pkg/database"
	"github.com/akifkadioglu/vocapedia/pkg/server"
	"github.com/akifkadioglu/vocapedia/pkg/token"
)

func Execute() {
	token.InitTokenAuth()

	database.InitDB(
		config.ReadValue().Database.Host, config.ReadValue().Database.Port,
		config.ReadValue().Database.User, config.ReadValue().Database.Password,
		config.ReadValue().Database.Name,
	)

	server.HttpServer(
		config.ReadValue().Host,
		config.ReadValue().Port,
		config.ReadValue().AllowMethods,
		config.ReadValue().AllowOrigins,
		config.ReadValue().AllowHeaders,
	)

}
