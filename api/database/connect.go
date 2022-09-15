package database

import (
	"api/models"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := "host=localhost user=postgres password=5rPwU1=x9tUHyT=jiC+KL0ht dbname=geminiapp_db port=40300"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect to database")
	}

	DB = database

	fmt.Println("Connection Opened to Database")
	DB.AutoMigrate(&models.User{})
	fmt.Println("Database Migrated")
}
