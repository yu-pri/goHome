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
	v, err := s.InternalSensor()
	if err != nil {
		home.ReportAlert(err.Error(), "Cannot get internal Temp")
	}
	log.Println("\t", v)

	/*report to Phant*/
	/*
		err = home.ReportInternalTemp(v)
		if err != nil {
			home.ReportAlert(err.Error(), "Cannot report internal Temp to Phant")
		}
	*/

	/*report to Cloud*/
	err = home.IOTReportTempInternal(v)
	if err != nil {
		home.ReportAlert(err.Error(), "Cannot report internal Temp to IOT Cloud")
	}

	/*update UI*/
	err = reportFloat("internal", v)
	if err != nil {
		home.ReportAlert(err.Error(), "Cannot report internal Temp to socket")
	}

	reverse, err := s.ReverseSensor()
	if err != nil {
		log.Println(err.Error())
	}
	log.Println("reverse \t", reverse)

	/*update UI*/
	err = reportFloat("temp2", reverse)
	if err != nil {
		home.ReportAlert(err.Error(), "Cannot report Temp to socket")
	}

	entry, err := s.EntryRoomSensor()
	if err != nil {
		log.Println(err.Error())
	}
	log.Println("entry \t", entry)

	/*update UI*/
	err = reportFloat("temp3", entry)
	if err != nil {
		home.ReportAlert(err.Error(), "Cannot report Temp to socket")
	}

	//x := &home.HData{}
	currentState.TempInside = v
	currentState.TempOutside = reverse
	currentState.TempEntryRoom = entry
	currentState.Timestamp = int(time.Now().Unix())
	x := currentState
	//x.TempInside = v
	//x.TempOutside = reverse
	//currentState.TempEntryRoom = entry
	//x.Timestamp = int(time.Now().Unix())

	historyData.Push(&x)
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
