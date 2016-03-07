package home

import (
  "log"
  //"os"
  //"time"
  //  "sync"
  //"github.com/hybridgroup/gobot"
  "github.com/hybridgroup/gobot/platforms/gpio"
  "github.com/hybridgroup/gobot/platforms/raspi"
  )

var r = raspi.NewRaspiAdaptor("raspi")
var relHeatMotor1 = gpio.NewLedDriver(r, "led", "11")
 /*
type Relays struct {

  RelHeatMotor2

  RelHeat1
  RelHeat2
}
*/

func init () {
  log.Println("init relays")

  relHeatMotor1.Off()
  /*
  r.RelHeatMotor2 := gpio.NewLedDriver(r, "led", "12")
  r.RelHeat1 := gpio.NewLedDriver(r, "led", "13")
  r.RelHeat2 := gpio.NewLedDriver(r, "led", "15")
*/
}

func ToggleHeatMotor1 () {
  log.Println("toggle motor")
  relHeatMotor1.Toggle()
}
