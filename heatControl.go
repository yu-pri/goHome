package main

import (
	"encoding/json"
	"goHome/home"
	"io"
	"log"
	"net/http"
)

func reportHeat() error {
	st := home.GetHeaterState()

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
		io.WriteString(w, home.GetHeaterState())
		return
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
}
