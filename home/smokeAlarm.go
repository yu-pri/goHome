package home

import (
	"log"

	"github.com/hybridgroup/gobot/platforms/gpio"
)

/*
SmokeAlarmKitchen smoke/gas sensor in the kitchen
*/
var SmokeAlarmKitchen = gpio.NewButtonDriver(r, "alm1", "18") //24

/*
SmokeAlarmSauna smoke/gas sensor in the sauna
*/
var SmokeAlarmSauna = gpio.NewButtonDriver(r, "alm2", "16") //23

func init() {
	log.Println("Initialise smoke alarms")
}
