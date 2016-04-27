package main

import (
	"io"
	"log"
	"net/http"
)

func hdata(w http.ResponseWriter, r *http.Request) {
	d, errs := historyData.ToJSON()
	if errs != nil {
		log.Println(errs.Error())
		http.Error(w, errs.Error(), http.StatusInternalServerError)
		return
	}
	io.WriteString(w, string(d))
}
