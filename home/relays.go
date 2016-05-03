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
var relHeatMotor1 = gpio.NewLedDriver(r, "led", "11")
var relHeatMotor2 = gpio.NewLedDriver(r, "led", "12")

var relHeat1 = gpio.NewLedDriver(r, "led", "13")
var relHeat2 = gpio.NewLedDriver(r, "led", "15")

/*
GetRelayAdaptor returns raspi adaptor for robot
*/
func GetRelayAdaptor() *raspi.RaspiAdaptor {
	return r
}

/*
GetRelHeatMotor1 returns led driver reference to RelayMotor1
*/
func GetRelHeatMotor1() *gpio.LedDriver {
	return relHeatMotor1
}

/*
GetRelHeatMotor2 returns led driver reference to RelayMotor1
*/
func GetRelHeatMotor2() *gpio.LedDriver {
	return relHeatMotor2
}

/*
GetHeat1 returns led driver reference to RelayMotor1
*/
func GetHeat1() *gpio.LedDriver {
	return relHeat1
}

/*
GetHeat2 returns led driver reference to RelayMotor1
*/
func GetHeat2() *gpio.LedDriver {
	return relHeat2
}

func init() {
	log.Println("init relays")
}

/*
ToggleHeatMotor1 switches state for heatmotor
*/
func ToggleHeatMotor1() error {
	log.Println("toggle heat motor r1")
	return relHeatMotor1.Toggle()
}

/*
ToggleHeatMotor2 switches state for heatmotor
*/
func ToggleHeatMotor2() error {
	log.Println("toggle heat motor r2")
	return relHeatMotor2.Toggle()
}

/*
ToggleHeat1 switches state for heatmotor
*/
func ToggleHeat1() error {
	log.Println("toggle heat r1")
	return relHeat1.Toggle()
}

/*
ToggleHeat2 switches state for heatmotor
*/
func ToggleHeat2() error {
	log.Println("toggle heat r2")
	return relHeat2.Toggle()
}

/*
OnHeatMotor1 switches state for heatmotor
*/
func OnHeatMotor1() error {
	log.Println("toggle heat motor r1")
	return relHeatMotor1.On()
}

/*
OnHeatMotor2 switches state for heatmotor
*/
func OnHeatMotor2() error {
	log.Println("toggle heat motor r2")
	return relHeatMotor2.On()
}

/*
OnHeat1 switches state for heatmotor
*/
func OnHeat1() error {
	log.Println("toggle heat r1")
	return relHeat1.On()
}

/*
OnHeat2 switches state for heatmotor
*/
func OnHeat2() error {
	log.Println("toggle heat r2")
	return relHeat2.On()
}

/*
OffHeatMotor1 switches state for heatmotor
*/
func OffHeatMotor1() error {
	log.Println("toggle heat motor r1")
	return relHeatMotor1.Off()
}

/*
OffHeatMotor2 switches state for heatmotor
*/
func OffHeatMotor2() error {
	log.Println("toggle heat motor r2")
	return relHeatMotor2.Off()
}

/*
OffHeat1 switches state for heatmotor
*/
func OffHeat1() error {
	log.Println("toggle heat r1")
	return relHeat1.Off()
}

/*
OffHeat2 switches state for heatmotor
*/
func OffHeat2() error {
	log.Println("toggle heat r2")
	return relHeat2.Off()
}

/*
SetHeat switches state for heatmotor
*/
func SetHeat(state string) error {
	if state == AUTO {
		err := OnHeat1()
		if err != nil {
			log.Println(err.Error())
			return err
		}

		err = OffHeat2()
		if err != nil {
			log.Println(err.Error())
			return err
		}
		return nil
	}

	if state == ON {
		err := OffHeat1()
		if err != nil {
			log.Println(err.Error())
			return err
		}

		err = OnHeat2()
		if err != nil {
			log.Println(err.Error())
			return err
		}
		return nil
	}

	if state == OFF {
		err := OffHeat1()
		if err != nil {
			log.Println(err.Error())
			return err
		}

		err = OffHeat2()
		if err != nil {
			log.Println(err.Error())
			return err
		}
		return nil
	}
	return nil
}

/*
SetBoiler switches state for heatmotor
*/
func SetBoiler(state string) error {
	if state == AUTO {
		err := OnHeatMotor1()
		if err != nil {
			log.Println(err.Error())
			return err
		}

		err = OffHeatMotor2()
		if err != nil {
			log.Println(err.Error())
			return err
		}
		return nil
	}

	if state == "On" {
		err := OffHeatMotor1()
		if err != nil {
			log.Println(err.Error())
			return err
		}

		err = OnHeatMotor2()
		if err != nil {
			log.Println(err.Error())
			return err
		}
		return nil
	}

	if state == "Off" {
		err := OffHeatMotor1()
		if err != nil {
			log.Println(err.Error())
			return err
		}

		err = OffHeatMotor2()
		if err != nil {
			log.Println(err.Error())
			return err
		}
		return nil
	}
	return nil
}

/*
GetRelHeat1 returns rely status
*/
func GetRelHeat1() bool {
	return relHeat1.State()
}

/*
GetRelHeat2 returns rely status
*/
func GetRelHeat2() bool {
	return relHeat2.State()
}

/*
GetRelMotor1 returns rely status
*/
func GetRelMotor1() bool {
	return relHeatMotor1.State()
}

/*
GetRelMotor2 returns rely status
*/
func GetRelMotor2() bool {
	return relHeatMotor2.State()
}

/*
GetHeaterState rteturns heater state
*/
func GetHeaterState() string {
	if !GetRelHeat1() && !GetRelHeat2() {
		return OFF
	}

	if GetRelHeat2() {
		return ON
	}

	return AUTO
}

/*
GetPumpState rteturns heater state
*/
func GetPumpState() string {
	if !GetRelMotor1() && !GetRelMotor2() {
		return OFF
	}

	if GetRelMotor2() {
		return ON
	}

	return AUTO
}
