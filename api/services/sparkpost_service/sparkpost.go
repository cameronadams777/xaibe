package sparkpost_service

import (
  "go/format"
	"api/helpers"
	"api/initializers/sparkpost"
	"bytes"
	"html/template"
	"log"

	sp "github.com/SparkPost/gosparkpost"
)

func get_subject_by_template(template_key string) string {
  switch template_key {
  case "reset_password":
    return "Reset Password: Galata.app"
  case "verify_email":
    return "Verify Email: Galata.app"
  default:
    return ""
  }
}

func format_html_template(html_template string, template_vars interface{}) (*string, error) {
  var processed bytes.Buffer
  template_compiler, _ := template.New("email_temp").Parse(html_template)
  
  compl_err := template_compiler.Execute(&processed, template_vars)

  if compl_err != nil {
    return nil, compl_err
  }

  formatted, format_err := format.Source(processed.Bytes())

  if format_err != nil {
    return nil, format_err
  }

  fmt_as_string := string(formatted)

  return &fmt_as_string, nil
}

func SendEmail(to string, template_key string, template_vars interface{}) error {
  subject := get_subject_by_template(template_key)
	html_template, temp_err := helpers.ReadEmailTemplate(template_key)

	if temp_err != nil {
		return temp_err
	}

  formatted_html_template, format_err := format_html_template(*html_template, template_vars)

  if format_err != nil {
    return format_err
  }

	// Create a Transmission using an inline Recipient List
	// and inline email Content.
	tx := &sp.Transmission{
		Recipients: []string{to},
		Content: sp.Content{
			HTML:    *formatted_html_template,
			From:    "noreply@galata.app",
			Subject: subject,
		},
	}

	id, _, err := sparkpost.SPClient.Send(tx)
	if err != nil {
		return err
	}

	log.Printf("Transmission sent with id [%s]\n", id)

	return nil
}
