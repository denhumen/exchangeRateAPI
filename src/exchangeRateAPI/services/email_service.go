package services

import (
	"exchangeRateAPI/src/exchangeRateAPI/db"
	"fmt"
	"log"
	"os"
	"strconv"

	"gopkg.in/gomail.v2"
)

// SendCurrentRateToSubscribers gets subscribers from database and sends them current rate
func SendCurrentRateToSubscribers(rate float64) {
	var subscribers []db.Subscriber
	db.DB.Find(&subscribers)

	mailer := gomail.NewMessage()
	mailer.SetHeader("From", os.Getenv("SMTP_USER"))
	mailer.SetHeader("Subject", "Daily USD to UAH Exchange Rate")

	smtpPort, err := strconv.Atoi(os.Getenv("SMTP_PORT"))
	if err != nil {
		log.Fatalf("Invalid SMTP_PORT: %s", err)
	}

	for _, subscriber := range subscribers {
		mailer.SetHeader("To", subscriber.Email)
		mailer.SetBody("text/plain", "Current USD to UAH exchange rate: "+fmt.Sprintf("%.2f", rate))

		dialer := gomail.NewDialer(
			os.Getenv("SMTP_HOST"),
			smtpPort,
			os.Getenv("SMTP_USER"),
			os.Getenv("SMTP_PASSWORD"),
		)

		if err := dialer.DialAndSend(mailer); err != nil {
			log.Printf("Could not send email to %s: %v", subscriber.Email, err)
		}

		log.Printf("Message sent to: %s", subscriber.Email)
	}
}
