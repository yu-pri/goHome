package main

import (
	"goHome/home"
	"log"
)

func managePump(dat *home.HData) {
	if home.GetHeatPumpMode() != home.AUTO {
		return
	}

	if dat.TempHeater >= home.HeaterPumpRunThreshold &&
		(dat.TempHeater-dat.TempInside) > 5 {
		err := home.OnHeatPump()
		if err != nil {
			home.ReportAlert(err.Error(), "ALARM: cannot turn on Heater Pump")
			log.Println(err.Error())
		}

	} else {
		err := home.OffHeatPump()
		if err != nil {
			home.ReportAlert(err.Error(), "ALARM: cannot turn off Heater Pump")
			log.Println(err.Error())
		}
	}

	dat.PumpState = home.GetPump()
}

func manageHeater(dat *home.HData) {
	if home.GetHeatMode() == home.AUTO {
		return
	}

	if dat.TempEntryRoom < float32(home.HeaterOnThreshold) {
		err := home.OnHeat()
		if err != nil {
			home.ReportAlert(err.Error(), "ALARM: cannot turn on Heater")
			log.Println(err.Error())
		}
	} else {
		err := home.OffHeat()
		if err != nil {
			home.ReportAlert(err.Error(), "ALARM: cannot turn off Heater")
			log.Println(err.Error())
		}
	}
	dat.HeaterState = home.GetHeat()
}
