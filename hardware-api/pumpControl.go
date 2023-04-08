package hardware_api

import (
	"encoding/json"
	"goHome/home"
	"io"
	"log"
	"net/http"
)

func reportPump() error {
	st := home.GetHeatPumpMode()

	r := stringReport{"pumpStateChanged", "state", st}

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

	currentState.PumpState = home.GetPump()
	err := reportCurrentState(&currentState)
	if err != nil {
		home.ReportAlert(err.Error(), "Cannot report Temp to socket")
	}

	return nil
}

func pump(w http.ResponseWriter, r *http.Request) {
	defer reportPump()

	w.Header().Set("Access-Control-Allow-Origin", "*")
	state := r.FormValue("state")
	log.Println(state)

	if !SENSORS {
		io.WriteString(w, home.AUTO)
		reportPump()
		return
	}

	if len(state) == 0 {
		log.Println("state requested:")
		io.WriteString(w, home.GetHeatPumpMode())
		return
	}

	errr := home.SetHeatPumpMode(state)
	if errr != nil {
		http.Error(w, errr.Error(), http.StatusInternalServerError)
		return
	}

	conf.PumpState = home.GetHeatPumpMode()
	err := conf.saveConfig()
	if err != nil {
		log.Println("Cannot save config: ", err.Error())
	}

	err = managePump(&currentState)
	if err != nil {
		http.Error(w, errr.Error(), http.StatusInternalServerError)
		return
	}
}
