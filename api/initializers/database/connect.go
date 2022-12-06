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
  host := config.Get("REDIS_HOST")
  user := config.Get("REDIS_USER")
  password := config.Get("REDIS_PASSWORD")
  dbname := config.Get("REDIS_DB")
  port := config.Get("REDIS_PORT")

	dsn := "host=" + host + " user=" + user + " password=" + password + " dbname=" + dbname + " port=" + port
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect to database")
	}

	DB = database

	fmt.Println("Connection Opened to Database")

	DB.AutoMigrate(&models.User{}, &models.Team{}, &models.Application{}, &models.AlertSchema{}, &models.ServiceToken{})

	fmt.Println("Database Migrated")
}
