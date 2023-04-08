package hardware_api

import (
	"encoding/json"
	"goHome/home"
	"io"
	"log"
	"net/http"
	"strconv"
)

func reportConfig() error {
	r := stringReport{"desiredTempChanged", "state", conf.DesiredTemp}

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
	currentState.HeaterState = home.GetHeat()
	err := reportCurrentState(&currentState)
	if err != nil {
		home.ReportAlert(err.Error(), "Cannot report Temp to socket")
	}

	return nil
}

func configHandler(w http.ResponseWriter, r *http.Request) {
	defer reportConfig()
	d, errs := conf.ToJSON()
	if errs != nil {
		log.Println(errs.Error())
		http.Error(w, errs.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	//io.WriteString(w, string(d))

	state := r.FormValue("state")
	if len(state) == 0 {
		log.Println("config requested:")
		io.WriteString(w, string(d))
		return
	}
	_, errs = strconv.Atoi(state)
	if errs != nil {
		log.Println(errs.Error())
		http.Error(w, errs.Error(), http.StatusInternalServerError)
		return
	}

	conf.DesiredTemp = state
	log.Println("desired temp", state)
	err := conf.saveConfig()
	if err != nil {
		log.Println("Cannot save config: ", err.Error())
	}

}
