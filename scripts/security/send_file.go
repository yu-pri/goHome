package main

import (
        "log"
				//"fmt"
        "flag"
        "gopkg.in/gomail.v2"
				//"scripts/amx/lib"
				//"strconv"
				"os"
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

  flag.Parse();

  var fname string;
  if (len(flag.Args()) > 0 ) {
    fname = flag.Args()[0];
  }

  log.Println(fname)

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

  m.SetHeader("Subject", "Motion file: " + fname)
  m.SetBody("text/html", "Hello <b>Alexander</b> </br> Motion detected! </br> \n motion file attached")

  if (len(fname) > 0)  {
    if _, err := os.Stat(fname); err == nil {
      m.Attach(fname)
    }
  }

  d := gomail.NewPlainDialer("smtp.gmail.com", 587, "alPrihodko@gmail.com", "ahvpiuembqkvszpi")

  // Send the email
  if err := d.DialAndSend(m); err != nil {
      panic(err)
  }

}
