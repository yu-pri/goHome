package main

import (
	"goHome/home"
	"log"
)


type stringReport struct {
	Type  string
	Key   string
	Value string
}


/*
"temperature" works with sensors data

temp1
temp2
temp3
temp4

*/
func reportCurrentState(dat *home.HData) error {

	d, errs := dat.ToJSON()
	if errs != nil {
		log.Println(errs.Error())
		return errs
	}
	for _, ws := range conns.ws {
		_, err02 := ws.Write(d)
		if err02 != nil {
			return err02
		}
	}

	return nil
}
