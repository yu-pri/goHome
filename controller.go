package main

import (
	"fmt"
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
const SENSORS = false

type socketConns struct {
	ws   map[int32]*websocket.Conn
	lock *sync.Mutex
}

var conns socketConns
var currentState home.HData

func echoHandler(ws *websocket.Conn) {
	var id = int32(time.Now().Unix())
	conns.lock.Lock()
	conns.ws[id] = ws
	conns.lock.Unlock()
	defer func() {
		conns.lock.Lock()
		delete(conns.ws, id)
		conns.lock.Unlock()
	}()
	msg := make([]byte, 512)
	for {
		n, err := ws.Read(msg)
		if err != nil {
			log.Println(err)
			return
		}

		fmt.Printf("Receive: %s\n", msg[:n])
		/*
			home.ToggleHeatMotor1()
			m, err := ws.Write(msg[:n])
			if err != nil {
				log.Println(err)
				return
			}
			fmt.Printf("Send: %s\n", msg[:m])
		*/
	}
}

func main() {
	var stop chan bool
	var err error
	conns = socketConns{make(map[int32]*websocket.Conn), &sync.Mutex{}}
	currentState = home.HData{}
	currentState.Index = 2
	log.Println(currentState.Index)

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
