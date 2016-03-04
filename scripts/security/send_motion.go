package main

import (
        //"log"
				//"fmt"
        "gopkg.in/gomail.v2"
				//"scripts/amx/lib"
				//"strconv"
				//"os/exec"
				//"bytes"
)

var (
	from string
	to string
)

func init() {
	from = "alPrihodko@gmail.com"
	to = "alPrihodko@gmail.com" //report errors to this email
}

/*
func reportErr(point string, err error) {
		er := tpcsmtp.Send(from, to, point, err.Error())
    if err != nil {
       log.Fatal(er)
    }
		log.Fatal(err)
}
*/


func main() {

	body := "event registered!\n"

	body = body + "=====================================\r\n"

  /*
	err := tpcsmtp.Send(from, "alPrihodko@gmail.com", "[MOTION DETECTED]", body)
	if err != nil {
       	reportErr("sending to owner", err)
  }
  */

  m := gomail.NewMessage()
  m.SetHeader("From", "alPrihodko@gmail.com")
  m.SetHeader("To", "alPrihodko@gmail.com")
  //m.SetAddressHeader("Cc", "dan@example.com", "Dan")
  m.SetHeader("Subject", "MOTION DETECTED")
  m.SetBody("text/html", "Hello <b>Alexander</b> </br> Motion detected!")
  //m.Attach("/home/Alex/lolcat.jpg")

  d := gomail.NewPlainDialer("smtp.gmail.com", 587, "alPrihodko@gmail.com", "ahvpiuembqkvszpi")

  // Send the email
  if err := d.DialAndSend(m); err != nil {
      panic(err)
  }

}
