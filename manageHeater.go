package main

import (
	"goHome/home"
	"log"
)

func managePump(dat *home.HData) {
	if dat.TempHeater >= home.HeaterPumpRunThreshold {
		if !home.GetPump() {
			log.Println("\t\tHeatPump On")
			err := home.OnHeatPump()
			if err != nil {
				home.ReportAlert(err.Error(), "ALARM: cannot turn on Heater Pump")
				log.Println(err.Error())
			}
		}
	} else {
		if home.GetPump() {
			log.Println("\t\tHeatPump Off")
			err := home.OffHeatPump()
			if err != nil {
				home.ReportAlert(err.Error(), "ALARM: cannot turn off Heater Pump")
				log.Println(err.Error())
			}
		}
	}
	dat.PumpState = home.GetPump()
}

func manageHeater(dat *home.HData) {
	if dat.TempEntryRoom < float32(home.HeaterOnThreshold) {
		if !home.GetHeat() {
			log.Println("\t\tHeat On")
			err := home.OnHeat()
			if err != nil {
				home.ReportAlert(err.Error(), "ALARM: cannot turn on Heater")
				log.Println(err.Error())
			}
		}
	} else {
		if home.GetHeat() {
			log.Println("\t\tHeat Off")
			err := home.OffHeat()
			if err != nil {
				home.ReportAlert(err.Error(), "ALARM: cannot turn off Heater")
				log.Println(err.Error())
			}
		}
	}
	dat.HeaterState = home.GetHeat()
}
