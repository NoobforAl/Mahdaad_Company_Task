package main

import (
	"fmt"
	"log"

	"github.com/avast/retry-go"
)

type EventService struct {
	MailService *MailService
}

func NewEventService(mailService *MailService) *EventService {
	return &EventService{
		MailService: mailService,
	}
}

func (e *EventService) NotifyUser(email, subject, message string) error {
	err := retry.Do(
		func() error {
			err := e.MailService.SendEmail(email, subject, message)
			if err != nil {
				log.Printf("Failed to send email: %v", err)
			}
			return err
		},

		// TODO: Add Configuration for Retry
		retry.Attempts(3),
		retry.DelayType(retry.FixedDelay),
		retry.Delay(1000),
	)

	if err != nil {
		return fmt.Errorf("failed to notify user after 3 attempts: %s", err)
	}

	return nil
}
