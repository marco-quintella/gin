package main

import (
	"gin/config"
	"gin/models"
	"gin/routes"
)

var (
	db = config.ConnectDB()
)

func main() {
	err := db.AutoMigrate(&models.Todo{})
	if err != nil {
		return
	}

	defer config.DisconnectDB(db)

	routes.Routes()
}
