package main

import (
	"fmt"
	"log"
)

func main() {
	mailServer := NewMailService(
		"smtp.example.com",
		443,
		"username",
		"password",
	)

	eventService := NewEventService(mailServer)

	var maxUser = 20

	for i := range maxUser {
		email := fmt.Sprintf("test%d@example.com", i)
		subject := fmt.Sprintf("Test Subject %d", i)
		message := fmt.Sprintf("Test Message %d", i)

		err := eventService.NotifyUser(email, subject, message)
		if err != nil {
			log.Printf("Error notifying user: %v", err)
		}
	}
}
