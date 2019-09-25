package api

import (
	"alif/quota/api/config"
	"alif/quota/api/server"
	"alif/quota/app/models"
)

func Run() {
	// Please, read ATTENTION mark in db connection.go file
	// There is an explanation why ConnectDB is in models package
	models.ConnectDB()
	defer models.DisconnectDB()

	server.Start(config.Peek().Service)
}
