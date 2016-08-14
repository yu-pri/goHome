package home

import (
	"log"

	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/platforms/gpio"
)

var firstTimeKitchen = true
var firstTimeSauna = true

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
	gobot.On(SmokeAlarmKitchen.Event("release"), func(data interface{}) {
		log.Println("Smoke alarm in the kitchen: On")
		/*
			err = home.Sms("Test", "Smoke/Gas detected in the Kitchen", home.Recipients)
			if err != nil {
				log.Println(err.Error())
			}
			home.ReportAlert("Something is really not OK", "Smoke/Gas detected in the Kitchen")
		*/
	})

	gobot.On(SmokeAlarmKitchen.Event("push"), func(data interface{}) {
		if firstTimeKitchen {
			firstTimeKitchen = false
			return
		}
		log.Println("Smoke alarm in the kitchen: Off")
		/*
			err = home.Sms("Test", "Smoke/Gas in the Kitchen - all good", home.Recipients)
			if err != nil {
				log.Println(err.Error())
			}
			home.ReportAlert("Now it's better", "Smoke/Gas in the Kitchen - all good")
		*/
	})

	gobot.On(SmokeAlarmSauna.Event("release"), func(data interface{}) {

		log.Println("Smoke alarm in the Sauna: On")
		/*
			err = home.Sms("Test", "Smoke/Gas detected in the Sauna", home.Recipients)
			if err != nil {
				log.Println(err.Error())
				}

			home.ReportAlert("Something is really not OK", "Smoke/Gas detected in the Sauna")
		*/
	})

	gobot.On(SmokeAlarmSauna.Event("push"), func(data interface{}) {
		if firstTimeSauna {
			firstTimeSauna = false
			return
		}

		log.Println("Smoke alarm in the Sauna: Off")
		/*
			err = home.Sms("Test", "Smoke/Gas in Sauna - all good", home.Recipients)
			if err != nil {
				log.Println(err.Error())
			}
			home.ReportAlert("Now it's better", "Smoke/Gas in Sauna - all good")
		*/
	})
}
