package sparkpost

import (
	"api/config"
	"log"

	sp "github.com/SparkPost/gosparkpost"
)

var SPClient *sp.Client

func CreateSparkPostClient() {
	apiKey := config.Get("SPARKPOST_API_KEY")
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
