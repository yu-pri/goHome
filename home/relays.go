package home

import (
	"log"
	"os"
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

	/*TESTS emanbles test mode*/
	//TESTS = true
)

var r = raspi.NewRaspiAdaptor("raspi")

//var relHeatPump = gpio.NewLedDriver(r, "led", "11") //17
var relHeatPump = gpio.NewLedDriver(r, "led", "16") //23
//var relHeatPump = gpio.NewLedDriver(r, "led", "33") //just test

//physical port 12 (BCM 18) seems broken, do not use it.
//var relHeat = gpio.NewLedDriver(r, "led", "13") //27
var relHeat = gpio.NewLedDriver(r, "led", "18") //24

var heatMode = AUTO
var heatPumpMode = AUTO

/*
SENSORS is sensors available? export SENSORS=0 run app in test mode without sensors
*/
var SENSORS = os.Getenv("SENSORS")

/*
Stop - Set relays to default position
*/
func Stop() {
	log.Println("Set relays to default state")
	err := OffHeat()
	if err != nil {
		log.Println(err.Error())
	}

	err = OnHeatPump()
	if err != nil {
		log.Println(err.Error())
	}
}

func init() {
	log.Println("Set init state for relays")
	if SENSORS != "0" {
		err := OffHeat()
		if err != nil {
			log.Fatal(err.Error())
		}

		err = OnHeatPump()
		if err != nil {
			log.Fatal(err.Error())
		}
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
OnHeatPump switches state for heatpump
*/
func OnHeatPump() error {
	log.Println("on heat motor r1")
	if GetPump() {
		log.Println("Heat Pump already on")
		return nil
	}
	return relHeatPump.On()
}

/*
OnHeat switches state for heat,  reversed
*/
func OnHeat() error {
	log.Println("on heat r1")
	if GetHeat() {
		log.Println("Heat already on")
		return nil
	}
	return relHeat.Off()
}

/*
OffHeatPump switches state for heatmotor, please note it's reversed
*/
func OffHeatPump() error {
	log.Println("off heat motor r1")
	if !GetPump() {
		log.Println("Heat Pump already off")
		return nil
	}

	return relHeatPump.Off()
}

/*
OffHeat switches state for heat, note that it's reversed
*/
func OffHeat() error {
	log.Println("off heat r1")

	log.Println("on heat motor r1")
	if !GetHeat() {
		log.Println("Heat already off")
		return nil
	}

	return relHeat.On()
}

/*
SetHeatMode switches state for heatmotor
*/
func SetHeatMode(state string) error {
	heatMode = state
	var err error
	switch state {
	case AUTO:
		err = OffHeat()
		break

	case ON:
		err = OnHeat()
		break

	case OFF:
		err = OffHeat()
		break

	}

	return err
}

/*
SetHeatPumpMode switches state for heatmotor
*/
func SetHeatPumpMode(state string) error {
	heatPumpMode = state
	var err error
	switch state {
	case AUTO:
		err = OnHeatPump()
		break

	case ON:
		err = OnHeatPump()
		break

	case OFF:
		err = OffHeatPump()
		break

	}

	return err
}

/*
GetHeat returns rely status, reversing
*/
func GetHeat() bool {
	return !relHeat.State()
}

/*
GetPump returns rely status, DO NOT reverse it
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
