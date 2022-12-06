package sparkpost

import (
	"log"
	"os"

	sp "github.com/SparkPost/gosparkpost"
)

var SPClient *sp.Client

func CreateSparkPostClient() {
	apiKey := os.Getenv("SPARKPOST_API_KEY")
	cfg := &sp.Config{
		BaseUrl:    "https://api.sparkpost.com",
		ApiKey:     apiKey,
		ApiVersion: 1,
	}
	var client sp.Client
	err := client.Init(cfg)
	if err != nil {
		log.Fatalf("SparkPost client init failed: %s\n", err)
	}
	SPClient = &client
}
