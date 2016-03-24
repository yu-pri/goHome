package main

import (
	"encoding/json"
	"goHome/home"
	"log"
	"time"
)

type termoReport struct {
	Key  string
	Name string
	Val  float32
}

func reportInternalSensor(s *home.Sensors) {
	v, err := s.InternalSensor()
	if err != nil {
		home.ReportAlert(err.Error(), "Cannot get internal Temp to Phant")
	}
	log.Println("\t", v)
	err = home.ReportInternalTemp(v)
	if err != nil {
		home.ReportAlert(err.Error(), "Cannot report internal Temp to Phant")
	}
	err = report(v)
	if err != nil {
		home.ReportAlert(err.Error(), "Cannot report internal Temp to socket")
	}
}

func report(t float32) error {
	r := termoReport{"temperature", "temp1", t}

	for _, ws := range conns.ws {
		b, err01 := json.Marshal(r)
		if err01 != nil {
			return err01
		}

		m, err02 := ws.Write(b)
		if err02 != nil {
			return err02
		}
		log.Println(m)
	}
	return nil
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

func scheduleT(what func(t float32) error, delay time.Duration, t float32) chan bool {
	stop := make(chan bool)

	go func() {
		for {
			what(t)
			select {
			case <-time.After(delay):
			case <-stop:
				return
			}
		}
	}()

	return stop
}
