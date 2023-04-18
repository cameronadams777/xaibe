package sparkpost_service

import (
	"api/helpers"
	"api/initializers/sparkpost"
	"log"

	sp "github.com/SparkPost/gosparkpost"
)

func get_subject_by_template(template_key string) string {
	switch template_key {
	case "reset_password":
		return "Reset Password: Xaibe"
	case "verify_email":
		return "Verify Email: Xaibe"
	case "invite_new_user":
		return "You're invited!"
	default:
		return ""
	}
}

func SendEmail(to string, template_key string, template_vars interface{}) error {
	subject := get_subject_by_template(template_key)
	html_template, temp_err := helpers.ReadEmailTemplate(template_key)

	if temp_err != nil {
		return temp_err
	}

	// Create a Transmission using an inline Recipient List
	// and inline email Content.
	tx := &sp.Transmission{
		Recipients: []string{to},
		Content: sp.Content{
			HTML:    *html_template,
			From:    "noreply@galata.app",
			Subject: subject,
		},
		Metadata: template_vars,
	}

	id, _, err := sparkpost.SPClient.Send(tx)
	if err != nil {
		return err
	}

	log.Printf("Transmission sent with id [%s]\n", id)

	return nil
}
