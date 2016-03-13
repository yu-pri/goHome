package home

import (
	"gopkg.in/gomail.v2"
)

/*
ReportAlert sends message to the system owner "emailTo"
*/
func ReportAlert(b string, s string) error {

	m := gomail.NewMessage()
	m.SetHeader("From", emailFrom)
	m.SetHeader("To", emailTo)

	m.SetHeader("Subject", s)
	m.SetBody("text/html", b)
	//m.Attach("/home/Alex/lolcat.jpg")

	d := gomail.NewPlainDialer("smtp.gmail.com", 587, googleAccountName, googleAPIKey)

	// Send the email
	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}