package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func Get(key string) string {
	// load .env file
  mode := os.Getenv("GIN_MODE")
  if mode != "release" {
    err := godotenv.Load(".env")
    if err != nil {
      fmt.Print("Error loading .env file")
    }
  }
	return os.Getenv(key)
}
