package home

import (
	"log"

	"github.com/hybridgroup/gobot/platforms/gpio"
)

/*
SmokeAlarmKitchen smoke/gas sensor in the kitchen
*/
var SmokeAlarmKitchen = gpio.NewButtonDriver(r, "alm1", "16") //23

/*
SmokeAlarmSauna smoke/gas sensor in the sauna
*/
var SmokeAlarmSauna = gpio.NewButtonDriver(r, "alm2", "18") //24

func init() {
	log.Println("Initialise smoke alarms")
}
