package main

import (
	"encoding/json"
	"goHome/home"
	"io"
	"log"
	"net/http"
)

func reportHeat(state string) error {
	r := stringReport{"heatStateChanged", "state", state}

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

func heat(w http.ResponseWriter, r *http.Request) {
	state := r.FormValue("state")
	log.Println(state)

	if !SENSORS {
		io.WriteString(w, "Auto")
		reportHeat("Auto")
		return
	}

	if len(state) == 0 {
		log.Println("state requested:")
		io.WriteString(w, home.GetHeaterState())
	}

	if state == "Auto" {
		err = home.OnHeat1()
		if err != nil {
			log.Println(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = home.OffHeat2()
		if err != nil {
			log.Println(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return
	}

	if state == "On" {
		err = home.OffHeat1()
		if err != nil {
			log.Println(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = home.OnHeat2()
		if err != nil {
			log.Println(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return
	}

	if state == "Off" {
		err = home.OffHeat1()
		if err != nil {
			log.Println(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = home.OffHeat2()
		if err != nil {
			log.Println(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return
	}

	reportHeat(state)

}
