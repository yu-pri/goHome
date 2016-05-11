package main

import (
	"encoding/json"
	"goHome/home"
	"io"
	"log"
	"net/http"
)

func reportHeat() error {
	st := home.GetHeatMode()

	r := stringReport{"heatStateChanged", "state", st}

	b, err01 := json.Marshal(r)
	if err01 != nil {
		return err01
	}

	for _, ws := range rconns.ws {
		m, err02 := ws.Write(b)
		if err02 != nil {
			return err02
		}
		log.Println(m)
	}

	err := reportCurrentState(&currentState)
	if err != nil {
		home.ReportAlert(err.Error(), "Cannot report Temp to socket")
	}

	return nil
}

func heat(w http.ResponseWriter, r *http.Request) {

	defer reportHeat()

	state := r.FormValue("state")
	log.Println(state)

	if !SENSORS {
		io.WriteString(w, home.AUTO)
		reportHeat()
		return
	}

	if len(state) == 0 {
		log.Println("state requested:")
		io.WriteString(w, home.GetHeatMode())
		return
	}

	errr := home.SetHeatMode(state)
	if errr != nil {
		http.Error(w, errr.Error(), http.StatusInternalServerError)
	}
}
