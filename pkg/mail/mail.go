package mail

import (
	"fmt"
	"net/smtp"
)

const (
	smtpPort = 587
	smtpHost = "smtp.gmail.com"
)

type MailClient struct {
	email    string
	password string
}

func NewMailClient(email, password string) *MailClient {
	return &MailClient{
		email:    email,
		password: password,
	}
}

func (m *MailClient) SendMail(subject, content, to string) error {
	auth := smtp.PlainAuth("", m.email, m.password, smtpHost)
	addr := fmt.Sprintf("%s:%d", smtpHost, smtpPort)

	sub := fmt.Sprintf("Subject: %s\r\n", subject)
	mime := "MIME-version: 1.0;\r\nContent-Type: text/html; charset=\"UTF-8\";\r\n\r\n"
	msg := []byte(sub + mime + content)

	err := smtp.SendMail(addr, auth, m.email, []string{to}, msg)
	if err != nil {
		return err
	}

	return nil
}
