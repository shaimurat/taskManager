package main

import (
	"log"
	"taskManager/config"
	"taskManager/internal/repository/postgres/models"
	"taskManager/utils"
)

func main() {
	utils.LoadEnv()
	db := config.ConnectPostgres()
	err := db.AutoMigrate(&models.UserDoc{}, &models.TaskDoc{})
	if err != nil {
		log.Fatalf("Error when auto migrating DB:%v", err)
	}
}
