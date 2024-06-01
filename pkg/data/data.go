package data

import (
	"bytes"
	"html/template"
)

type Data struct {
	BillsDue           int
	Date               string
	Selected           bool
	ConfirmationNumber int
	Message            string
}

func NewData() *Data {
	return &Data{}
}

func (d *Data) GenerateEmail() (subject, content string, err error) {
	tmpl, err := template.New("email").Parse(emailTemplate)
	if err != nil {
		return "", "", err
	}

	// Execute the template with the provided data
	var emailBody bytes.Buffer
	err = tmpl.Execute(&emailBody, *d)
	if err != nil {
		return "", "", err
	}

	if !d.Selected {
		subject = "NOT SELECTED for testing today"
	} else {
		subject = "YES SELECTED for testing today"
	}

	content = emailBody.String()

	return subject, content, nil
}
