package main

import (
	"flag"
	"goHome/home"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/hybridgroup/gobot"

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
var rconns socketConns

var currentState home.HData
var historyData home.HistoryData

func main() {

	var stop chan bool
	gbot := gobot.NewGobot()

	conns = socketConns{make(map[int32]*websocket.Conn), &sync.Mutex{}}
	rconns = socketConns{make(map[int32]*websocket.Conn), &sync.Mutex{}}
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
	http.Handle("/relays", websocket.Handler(relHandler))

	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/control/pump", pump)
	http.HandleFunc("/control/heat", heat)

	http.HandleFunc("/control/hdata", hdata)

	go func() {
		processInput(&currentState)
	}()

	work := func() {
		gobot.Every(100000*time.Millisecond, func() {
			log.Println("gobot heartbeat")
			//      led.Toggle()
		})
	}

	robot := gobot.NewRobot("blinkBot",
		[]gobot.Connection{home.GetRelayAdaptor()},
		[]gobot.Device{home.GetHeat1(), home.GetHeat2(), home.GetRelHeatMotor1(),
			home.GetRelHeatMotor2()},
		work,
	)

	gbot.AddRobot(robot)

	go gbot.Start()

	err = http.ListenAndServe(":1234", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}

	if stop != nil {
		stop <- true
	}
}
