package main

import (
	"os"

	"github.com/go-mail/mail/v2"
)

const (
	host     = "sandbox.smtp.mailtrap.io"
	port     = 587
	username = "f54af7c050e7b2"
	password = "09f8a9202ff4af"
)

func main() {
	from := "test@lenslocked.com"
	to := "jon@calhoun.io"
	subject := "This is a test email"
	plaintext := "This is the body of the email"
	html := `<h1>Hello there buddy!</h1><p>This is the email</p><p>Hope you enjoy it</p>`

	msg := mail.NewMessage()
	msg.SetHeader("To", to)
	msg.SetHeader("From", from)
	msg.SetHeader("Subject", subject)
	msg.SetBody("text/plain", plaintext)
	msg.AddAlternative("text/html", html)
	msg.WriteTo(os.Stdout)

	dialer := mail.NewDialer(host, port, username, password)
	err := dialer.DialAndSend(msg)
	if err != nil {
		panic(err)
	}
}
