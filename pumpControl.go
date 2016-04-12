package main

import (
	"encoding/json"
	"goHome/home"
	"io"
	"log"
	"net/http"
)

func reportPump(state string) error {
	r := stringReport{"pumpStateChanged", "state", state}

	for _, ws := range rconns.ws {
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

func pump(w http.ResponseWriter, r *http.Request) {

	state := r.FormValue("state")
	log.Println(state)

	if !SENSORS {
		io.WriteString(w, "Auto")
		reportPump("Auto")
		return
	}

	if len(state) == 0 {
		log.Println("state requested:")
		io.WriteString(w, home.GetPumpState())
	}

	if state == "Auto" {
		err = home.OnHeatMotor1()
		if err != nil {
			log.Println(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = home.OffHeatMotor2()
		if err != nil {
			log.Println(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return
	}

	if state == "On" {
		err = home.OffHeatMotor1()
		if err != nil {
			log.Println(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = home.OnHeatMotor2()
		if err != nil {
			log.Println(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return
	}

	if state == "Off" {
		err = home.OffHeatMotor1()
		if err != nil {
			log.Println(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = home.OffHeatMotor2()
		if err != nil {
			log.Println(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		return
	}

	reportPump(state)
	return
}
