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
	err = reportFloat("temp1", v)
	if err != nil {
		home.ReportAlert(err.Error(), "Cannot report internal Temp to socket")
	}

	reverse, err := s.ReverseSensor()
	log.Println("reverse \t", v)

	/*update UI*/
	err = reportFloat("temp2", v)
	if err != nil {
		home.ReportAlert(err.Error(), "Cannot report Temp to socket")
	}

	entry, err := s.EntryRoomSensor()
	log.Println("entry \t", v)

	/*update UI*/
	err = reportFloat("temp3", v)
	if err != nil {
		home.ReportAlert(err.Error(), "Cannot report Temp to socket")
	}

	x := &home.HData{}
	currentState.TempInside = v
	currentState.TempOutside = reverse
	currentState.TempEntryRoom = entry

	currentState.Timestamp = int(time.Now().Unix())
	x.TempInside = v
	x.Timestamp = int(time.Now().Unix())
	historyData.Push(x)
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
