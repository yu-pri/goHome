package home

import (
	"log"
	//"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/platforms/gpio"
	"github.com/hybridgroup/gobot/platforms/raspi"
)

const (
	/*ON Constant */
	ON = "On"
	/*OFF constant*/
	OFF = "Off"
	/*AUTO constant*/
	AUTO = "Auto"
)

var r = raspi.NewRaspiAdaptor("raspi")
var relHeatPump = gpio.NewLedDriver(r, "led", "11")
var relHeat = gpio.NewLedDriver(r, "led", "12")

var heatMode = AUTO
var heatPumpMode = AUTO

func init() {
	log.Println("Set init state for relays")
	err := OffHeat()
	if err != nil {
		log.Fatal(err.Error())
	}

	err = OnHeatPump()
	if err != nil {
		log.Fatal(err.Error())
	}

}

//TODO: more relays can be served here :-)
//var relHeatMotor2 = gpio.NewLedDriver(r, "led", "13")
//var relHeat2 = gpio.NewLedDriver(r, "led", "15")

/*
GetRelayAdaptor returns raspi adaptor for robot
*/
func GetRelayAdaptor() *raspi.RaspiAdaptor {
	return r
}

/*
GetRelHeatPump returns led driver reference to RelayMotor1
*/
func GetRelHeatPump() *gpio.LedDriver {
	return relHeatPump
}

/*
GetRelHeat returns led driver reference to RelayMotor1
*/
func GetRelHeat() *gpio.LedDriver {
	return relHeat
}

/*
ToggleHeatPump returns led driver reference to RelayMotor1
*/
func ToggleHeatPump() error {
	log.Println("toggle heat motor r1")
	return relHeatPump.Toggle()
}

/*
ToggleHeat switches state for heatmotor
*/
func ToggleHeat() error {
	log.Println("toggle heat r1")
	return relHeat.Toggle()
}

/*
OnHeatPump switches state for heatmotor
*/
func OnHeatPump() error {
	log.Println("on heat motor r1")
	return relHeatPump.On()
}

/*
OnHeat switches state for heatmotor
*/
func OnHeat() error {
	log.Println("on heat r1")
	return relHeat.On()
}

/*
OffHeatPump switches state for heatmotor
*/
func OffHeatPump() error {
	log.Println("off heat motor r1")
	return relHeatPump.Off()
}

/*
OffHeat switches state for heatmotor
*/
func OffHeat() error {
	log.Println("off heat r1")
	return relHeat.Off()
}

/*
SetHeatMode switches state for heatmotor
*/
func SetHeatMode(state string) error {
	heatMode = state
	return nil
}

/*
SetHeatPumpMode switches state for heatmotor
*/
func SetHeatPumpMode(state string) error {
	heatPumpMode = state
	return nil
}

/*
GetHeat returns rely status
*/
func GetHeat() bool {
	return relHeat.State()
}

/*
GetPump returns rely status
*/
func GetPump() bool {
	return relHeatPump.State()
}

/*
GetHeatMode rteturns heater state
*/
func GetHeatMode() string {
	return heatMode
}

/*
GetHeatPumpMode - returns heat pump mode
*/
func GetHeatPumpMode() string {
	return heatPumpMode
}
