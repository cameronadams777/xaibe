package database

import (
	"api/config"
	"api/models"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	host := config.Get("DB_HOST")
	user := config.Get("DB_USER")
	password := config.Get("DB_PASSWORD")
	dbname := config.Get("DB_NAME")
	port := config.Get("DB_PORT")

	dsn := "host=" + host + " user=" + user + " password=" + password + " dbname=" + dbname + " port=" + port
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect to database")
	}

	DB = database

	fmt.Println("Connection Opened to Database")

	DB.AutoMigrate(&models.User{}, &models.Team{}, &models.Application{}, &models.AlertSchema{}, &models.ServiceToken{}, &models.TeamInvite{})

	fmt.Println("Database Migrated")
}
