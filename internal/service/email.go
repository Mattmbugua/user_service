package service

import (
	"log"
	"time"
)

func sendWelcomeEmailAsync(name, email string) {
	go func() {
		log.Printf("Sending welcome email to %s <%s>\n", name, email)

		time.Sleep(1 * time.Second)
		log.Printf("Welcome email sent to %s <%s>\n", name, email)
	}()
}
