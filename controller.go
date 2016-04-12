package main

import (
	"flag"
	"goHome/home"
	"log"
	"net/http"
	"sync"
	"time"

	"golang.org/x/net/websocket"
)

/*
SENSORS  Sensors exists
*/
var SENSORS bool

var err error

type socketConns struct {
	ws   map[int32]*websocket.Conn
	lock *sync.Mutex
}

var conns socketConns
var currentState home.HData

func main() {

	var stop chan bool

	conns = socketConns{make(map[int32]*websocket.Conn), &sync.Mutex{}}
	currentState = home.HData{}
	currentState.Index = 2
	log.Println(currentState.Index)
	flag.BoolVar(&SENSORS, "sensors", true, "Sensors are there?")
	flag.Parse()

	if SENSORS {
		sensors, errs := home.NewSensors()
		if errs != nil {
			panic("Sensors: " + err.Error())
		}

		stop = schedule(reportInternalSensor, 60*time.Second, sensors)
	}

	//stop = scheduleT(reportFloat, 10*time.Second, "temp1", 10)

	http.Handle("/echo", websocket.Handler(echoHandler))
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/control/pump", pump)
	http.HandleFunc("/control/heat", heat)

	go func() {
		processInput(&currentState)
	}()

	err = http.ListenAndServe(":1234", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}

	if stop != nil {
		stop <- true
	}
}
