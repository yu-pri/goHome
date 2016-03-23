package main


import (
        "log"
        //"os"
        "time"
        //  "sync"
        "github.com/hybridgroup/gobot"
        "github.com/hybridgroup/gobot/platforms/gpio"
        "github.com/hybridgroup/gobot/platforms/raspi"
)

func main() {
  //state := 0
  gbot := gobot.NewGobot()

  r := raspi.NewRaspiAdaptor("raspi")
  /*
  pin17 := gpio.NewDirectPinDriver(r, "pin", "11")
  pin18 := gpio.NewDirectPinDriver(r, "pin", "12")

  pin21 := gpio.NewDirectPinDriver(r, "pin", "13")
  pin22 := gpio.NewDirectPinDriver(r, "pin", "15")

  //button := gpio.NewButtonDriver(r, "button", "15")
  */
  led := gpio.NewLedDriver(r, "led", "13")

  work := func() {
    gobot.Every(500*time.Millisecond, func() {
      log.Println("toggle");
      led.Toggle()
      })
    }



robot := gobot.NewRobot("blinkBot",
[]gobot.Connection{r},
[]gobot.Device{led},
work,
)

gbot.AddRobot(robot)

gbot.Start()

}
