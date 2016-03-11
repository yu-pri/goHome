package home

import (
        "gopkg.in/gomail.v2"
)


func reportAlert (b string, s string) error {
    //var config = home.LoadConfiguration()
    m := gomail.NewMessage()
    m.SetHeader("From", SOME)
    //m.SetHeader("From", EMAIL_FROM)
    //m.SetHeader("To", EMAIL_TO)
    m.SetHeader("Subject", s)
    m.SetBody("text/html", b)
    //m.Attach("/home/Alex/lolcat.jpg")

    d := gomail.NewPlainDialer("smtp.gmail.com", 587, "alPrihodko@gmail.com", "ahvpiuembqkvszpi")

    // Send the email
    if err := d.DialAndSend(m); err != nil {
        return err
    }

    return nil
}
