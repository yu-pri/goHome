package hardware_api

import (
	"io"
	"log"
	"net/http"
)

func cState(w http.ResponseWriter, r *http.Request) {

	d, errs := currentState.ToJSON()
	if errs != nil {
		log.Println(errs.Error())
		http.Error(w, errs.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	io.WriteString(w, string(d))
}
