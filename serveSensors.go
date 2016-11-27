package main

import (
	"goHome/home"
	"log"
	"time"
)

/*
ReportInternalSensor reports Sensors to web app and cloud
*/
func reportSensors(s *home.Sensors) {
	err := s.Temp.Update()

	v, err := s.InternalSensor()
	if err != nil {
		home.ReportAlert(err.Error(), "Cannot get internal Temp")
	}
	//log.Println("\t", v)
	/*report to Phant*/
	/*
		err = home.ReportInternalTemp(v)
		if err != nil {
			home.ReportAlert(err.Error(), "Cannot report internal Temp to Phant")
		}
	*/

	if err != nil {
		log.Println(err.Error())
	}

	/*report to Cloud*/
	/*
		err = home.IOTReportTempInternal(v)
		if err != nil {
			home.ReportAlert(err.Error(), "Cannot report internal Temp to IOT Cloud")
		}
	*/

	reverse, err := s.ReverseSensor()
	if err != nil {
		log.Println(err.Error())
	}
	//log.Println("reverse \t", reverse)

	entry, err := s.EntryRoomSensor()
	if err != nil {
		log.Println(err.Error())
	}
	//log.Println("entry \t", entry)

	heater, err := s.HeaterSensor()
	if err != nil {
		log.Println(err.Error())
	}
	//log.Println("heater \t", entry)

	//waterBoiler, err := s.WaterBoilerSensor()
	//if err != nil {
	//	log.Println(err.Error())
	//}

	recuperator, err := s.RecuperatorSensor()
	if err != nil {
		log.Println(err.Error())
	}

	outside, err := s.OutsideSensor()
	if err != nil {
		log.Println(err.Error())
	}

	//x := &home.HData{}
	currentState.TempInside = v
	currentState.TempReverse = reverse
	currentState.TempEntryRoom = entry
	currentState.TempHeater = heater
	currentState.TempWaterBoiler = 0
	currentState.TempRecuperator = recuperator
	currentState.TempOutside = outside

	currentState.Timestamp = int(time.Now().Unix())

	managePump(&currentState)
	manageHeater(&currentState)

	/*update UI*/
	err = reportCurrentState(&currentState)
	if err != nil {
		home.ReportAlert(err.Error(), "Cannot report Temp to socket")
	}

	x := currentState
	//x.TempInside = v
	//x.TempOutside = reverse
	//currentState.TempEntryRoom = entry
	//x.Timestamp = int(time.Now().Unix())

	historyData.Push(&x)
	/*
		err = home.IOTReportDev(x)
		if err != nil {
			home.ReportAlert(err.Error(), "Cannot report values to dev server")
		}
	*/
}

func schedule(what func(*home.Sensors), delay time.Duration, s *home.Sensors) chan bool {
	stop := make(chan bool)

	go func() {
		for {
			what(s)
			select {
			case <-time.After(delay):
			case <-stop:
				return
			}
		}
	}()

	return stop
}

func scheduleT(what func(id string, t float32) error, delay time.Duration, id string, t float32) chan bool {
	stop := make(chan bool)

	go func() {
		for {
			what(id, t)
			select {
			case <-time.After(delay):
			case <-stop:
				return
			}
		}
	}()

	return stop
}
