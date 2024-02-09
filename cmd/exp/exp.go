package main

import (
	"fmt"

	"github.com/xiaoshanjiang/lenslocked/models"
)

const (
	host     = "sandbox.smtp.mailtrap.io"
	port     = 587
	username = "f54af7c050e7b2"
	password = "09f8a9202ff4af"
)

func main() {
	email := models.Email{
		// Comment this out to test the default sender logic we added.
		From:    "test@lenslocked.com",
		To:      "jon@calhoun.io",
		Subject: "This is a test email",
		// Try sending emails with only one of these two fields set.
		Plaintext: "This is the body of the email",
		HTML:      `<h1>Hello there buddy!</h1><p>This is the email</p><p>Hope you enjoy it</p>`,
	}
	es := models.NewEmailService(models.SMTPConfig{
		Host:     host,
		Port:     port,
		Username: username,
		Password: password,
	})
	// Uncomment this to test the default sender logic we added.
	// es.DefaultSender = "bob@calhoun.io"
	err := es.Send(email)
	if err != nil {
		panic(err)
	}
	fmt.Println("Email sent")
}
