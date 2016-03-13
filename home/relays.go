package home

import (
	"log"
	//"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/platforms/gpio"
	"github.com/hybridgroup/gobot/platforms/raspi"
)

var r = raspi.NewRaspiAdaptor("raspi")
var relHeatMotor1 = gpio.NewLedDriver(r, "led", "11")
var relHeatMotor2 = gpio.NewLedDriver(r, "led", "12")

var relHeat1 = gpio.NewLedDriver(r, "led", "13")
var relHeat2 = gpio.NewLedDriver(r, "led", "15")

/*
type Relays struct {

  RelHeatMotor2

  RelHeat1
  RelHeat2
}
*/

func init() {
	log.Println("init relays")
}

/*
ToggleHeatMotor1 switches state for heatmotor
*/
func ToggleHeatMotor1() {
	log.Println("toggle motor")
	relHeatMotor1.Toggle()
}
