package main

import (
	"github.com/akifkadioglu/vocapedia/pkg/config"
	"github.com/akifkadioglu/vocapedia/pkg/server"
)

func main() {
	//app.Execute()

	server.HttpServer(
		config.ReadValue().Host,
		config.ReadValue().Port,
		config.ReadValue().AllowMethods,
		config.ReadValue().AllowOrigins,
		config.ReadValue().AllowHeaders,
	)
}
