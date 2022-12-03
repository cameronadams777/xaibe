package sparkpost_service

import (
	"api/initializers/sparkpost"
	"log"

	sp "github.com/SparkPost/gosparkpost"
)

func get_template_by_key(key string) (string, string) {
	return "", ""
}

func SendEmail(to string, template_key string) {
	template, subject := get_template_by_key(template_key)

	// Create a Transmission using an inline Recipient List
	// and inline email Content.
	tx := &sp.Transmission{
		Recipients: []string{to},
		Content: sp.Content{
			HTML:    template,
			From:    "admin@galata.app",
			Subject: subject,
		},
	}

	id, _, err := sparkpost.SPClient.Send(tx)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Transmission sent with id [%s]\n", id)
}
