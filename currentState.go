package main

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
	io.WriteString(w, string(d))
}