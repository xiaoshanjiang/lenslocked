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
	es := models.NewEmailService(models.SMTPConfig{
		Host:     host,
		Port:     port,
		Username: username,
		Password: password,
	})
	err := es.ForgotPassword("jon@calhoun.io", "https://lenslocked.com/reset-pw?token=abc123")
	if err != nil {
		panic(err)
	}
	fmt.Println("Email sent")
}
