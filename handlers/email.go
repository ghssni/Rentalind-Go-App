package handlers

import (
	"gopkg.in/gomail.v2"
)

func SendMail(email, subject, content string) {
	m := gomail.NewMessage()
	m.SetHeader("From", "opoaeseh@example.com") 
	m.SetHeader("To", email)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", content)

	d := gomail.NewDialer(
		"your_smtp_server.com", 
		587,
		"your_email@example.com", 
		"your_email_password",
	)

	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}

func SendSuccessCreateRent(email string) {
	SendMail(
		email,
		"Rental Successful",
		"Your rental has been successfully created. Thank you for using our service!",
	)
}