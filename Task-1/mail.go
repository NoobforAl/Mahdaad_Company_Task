package main

import (
	"fmt"
	"log"
	"math/rand/v2"
)

type MailService struct {
	SMTPServer string
	Port       int
	Username   string
	Password   string
}

func NewMailService(smtpServer string, port int, username, password string) *MailService {
	return &MailService{
		SMTPServer: smtpServer,
		Port:       port,
		Username:   username,
		Password:   password,
	}
}

// Example send email to user by 80% chance
func (m *MailService) SendEmail(to, subject, body string) error {
	if rand.Float64() < 0.8 {
		log.Printf("Email sent to %s: %s - %s\n", to, subject, body)
		return nil
	}

	return fmt.Errorf("failed to send email to %s", to)
}
