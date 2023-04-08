package hardware_api

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

	currentState.HeaterState = home.GetHeat()
	err := reportCurrentState(&currentState)
	if err != nil {
		home.ReportAlert(err.Error(), "Cannot report Temp to socket")
	}

	return nil
}

func heat(w http.ResponseWriter, r *http.Request) {

	defer reportHeat()
	w.Header().Set("Access-Control-Allow-Origin", "*")

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

	conf.HeaterState = home.GetHeatMode()
	err := conf.saveConfig()
	if err != nil {
		log.Println("Cannot save config: ", err.Error())
	}

	err = manageHeater(&currentState)
	if err != nil {
		http.Error(w, errr.Error(), http.StatusInternalServerError)
		return
	}
}
