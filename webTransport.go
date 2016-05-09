package main

import (
	"encoding/json"
	"goHome/home"
	"log"
)

type floatReport struct {
	Key  string
	Name string
	Val  float32
}

type stringReport struct {
	Type  string
	Key   string
	Value string
}

var prev *home.HData

/*
"temperature" works with sensors data

temp1
temp2
temp3
temp4

*/

func reportFloat(id string, t float32) error {
	r := floatReport{"temperature", id, t}

	for _, ws := range conns.ws {
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

func reportCurrentState(dat *home.HData) error {

	d, errs := dat.ToJSON()
	if errs != nil {
		log.Println(errs.Error())
		return errs
	}
	for _, ws := range conns.ws {
		m, err02 := ws.Write(d)
		if err02 != nil {
			return err02
		}
		log.Println(m)
	}

	return nil
}

/*
func reportCurrentState(dat *home.HData) error {
	var err error
	fmt.Println(dat)

	if prev == nil {
		prev = &home.HData{}
	}

	fmt.Println(prev)

	if prev.TempInside != dat.TempInside {
		err = reportFloat("temp1", dat.TempInside)
		if err != nil {
			return err
		}
	}

	if prev.TempOutside != dat.TempOutside {
		err = reportFloat("temp2", dat.TempOutside)
		if err != nil {
			return err
		}
	}

	if prev.TempHeater != dat.TempHeater {
		err = reportFloat("temp3", dat.TempHeater)
		if err != nil {
			return err
		}
	}

	if prev.TempReverse != dat.TempReverse {
		err = reportFloat("temp4", dat.TempReverse)
		if err != nil {
			return err
		}
	}

	*prev = *dat
	return nil
}
*/
