package main

import "goHome/home"

func managePump(dat *home.HData) {
	if dat.TempHeater >= home.HeaterPumpRunThreshold {
		home.OnHeatMotor1()
		home.OnHeatMotor2()
		home.OnHeat1()
		home.OffHeat1()
	} else {
		home.OffHeat1()
		home.OffHeat2()
		home.OffHeatMotor1()
		home.OffHeatMotor2()
	}
}
