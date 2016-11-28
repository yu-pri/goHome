package main

import (
	//	"goHome/home"
	"goHome/home"
	"log"
)

func managePump(dat *home.HData) error {
	var err error
	if home.GetHeatPumpMode() != home.AUTO {
		dat.HeaterState = home.GetHeat()
		dat.PumpState = home.GetPump()
		return nil
	}

	if dat.TempHeater >= home.HeaterPumpRunThreshold &&
		(dat.TempHeater-dat.TempInside) > 5 {
		err = home.OnHeatPump()
		if err != nil {
			home.ReportAlert(err.Error(), "ALARM: cannot turn on Heater Pump")
			log.Println(err.Error())
		}

	} else {
		err = home.OffHeatPump()
		if err != nil {
			home.ReportAlert(err.Error(), "ALARM: cannot turn off Heater Pump")
			log.Println(err.Error())
		}
	}

	dat.HeaterState = home.GetHeat()
	dat.PumpState = home.GetPump()
	return err
}

func manageHeater(dat *home.HData) error {
	var err error
	if home.GetHeatMode() != home.AUTO {
		dat.HeaterState = home.GetHeat()
		dat.PumpState = home.GetPump()

		return err
	}

	//hour := time.Now().Hour() + 1
	if dat.TempEntryRoom < float32(home.HeaterOnThreshold) &&
		home.IsChipTimeZone() {
		err = home.OnHeat()
		if err != nil {
			home.ReportAlert(err.Error(), "ALARM: cannot turn on Heater")
			log.Println(err.Error())
		}
	} else {
		err = home.OffHeat()
		if err != nil {
			home.ReportAlert(err.Error(), "ALARM: cannot turn off Heater")
			log.Println(err.Error())
		}
	}

	dat.HeaterState = home.GetHeat()
	dat.PumpState = home.GetPump()

	return err
}
