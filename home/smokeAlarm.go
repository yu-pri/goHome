package home

import (
	"log"

	"github.com/hybridgroup/gobot/platforms/gpio"
)

var smokeAlarmKitchen = gpio.NewButtonDriver(r, "alm1", "18") //24
var smokeAlarmSauna = gpio.NewButtonDriver(r, "alm2", "16")   //23

func init() {
	log.Println("Initialise smoke alarms")
}
