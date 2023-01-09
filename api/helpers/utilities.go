package helpers

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
)

func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}

func ReadEmailTemplate(template_key string) (*string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return nil, errors.New("ReadEmailTemplate: Error getting current working directory")
	}
	path := filepath.Join(cwd, "email_templates", template_key + ".html")
	dat, read_err := os.ReadFile(path)
	if read_err != nil {
		return nil, errors.New("ReadEmailTemplate: Error reading in " + template_key)
	}
	contents_as_string := string(dat)
	return &contents_as_string, nil
}
