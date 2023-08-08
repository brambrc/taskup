package main

import (
	"log"
	"taskup/Database"
	"taskup/Model"
	"taskup/Router"

	"github.com/joho/godotenv"
)

func main() {
	loadEnv()
	loadDatabase()
	Router.ServeApps()
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func loadDatabase() {

	Database.Connect()
	Database.Database.AutoMigrate(&Model.User{})
	Database.Database.AutoMigrate(&Model.Project{})
	Database.Database.AutoMigrate(&Model.Task{})
}
