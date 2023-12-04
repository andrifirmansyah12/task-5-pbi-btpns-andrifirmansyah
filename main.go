package main

import (
	"log"

	"github.com/andrifirmansyah12/projectGo/database"
	"github.com/andrifirmansyah12/projectGo/models"
	"github.com/andrifirmansyah12/projectGo/router"
)

func main() {
	err := database.InitDatabase()
	if err != nil {
		log.Fatalln("could not create database", err)
	}
	database.GlobalDB.AutoMigrate(&models.User{}, &models.Photo{})
	r := router.SetupRouter()
	r.Run(":8080")
}
