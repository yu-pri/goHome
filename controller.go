package main

import (
	"flag"
	"goHome/home"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	//_ "github.com/icattlecoder/godaemon"

	"github.com/hybridgroup/gobot"

	"golang.org/x/net/websocket"
)

const configFileName = "/etc/goHome.conf"

/*
HISTORYDATASERIAL file which contains history data for my home
*/
const HISTORYDATASERIAL = "/home/pi/goHomeHistoryData.b64"

/*
SENSORS  Sensors exists
*/
var SENSORS bool

/*
INTERVAL  Check sensors status with interval
*/
var INTERVAL int

//var err error

var conf Config

type socketConns struct {
	ws   map[int32]*websocket.Conn
	lock *sync.Mutex
}

var conns socketConns
var rconns socketConns

var currentState home.HData
var historyData home.HistoryData
var sensors *home.Sensors

//var stop chan bool

func main() {

	err := conf.loadConfig()
	if err != nil {
		log.Println("Likely use default configuration")
	}

	gbot := gobot.NewGobot()

	conns = socketConns{make(map[int32]*websocket.Conn), &sync.Mutex{}}
	rconns = socketConns{make(map[int32]*websocket.Conn), &sync.Mutex{}}
	currentState = home.HData{}
	//currentState.Index = 2
	log.Println(currentState.Index)
	flag.BoolVar(&SENSORS, "sensors", true, "Sensors are there?")
	flag.IntVar(&INTERVAL, "timeout", 60, "Timeout?")
	flag.Parse()

	if SENSORS {
		var errs error
		sensors, errs = home.NewSensors()
		if errs != nil {
			panic("Sensors: " + errs.Error())
		}
	}

	log.Println("Timeout interval to track sensors: ", INTERVAL)
	//stop = scheduleT(reportFloat, 10*time.Second, "temp1", 10)
	historyData.RestoreFromFile(HISTORYDATASERIAL)

	http.Handle("/echo", websocket.Handler(echoHandler))
	http.Handle("/relays", websocket.Handler(relHandler))

	http.Handle("/", http.FileServer(http.Dir("/home/pi/w/go/src/goHome")))
	http.HandleFunc("/control/pump", pump)
	http.HandleFunc("/control/heat", heat)

	http.HandleFunc("/control/hdata", hdata)
	http.HandleFunc("/control/currentState", cState)

	if !SENSORS {
		go func() {
			processInput(&currentState)
		}()
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	signal.Notify(c, os.Kill)
	signal.Notify(c, syscall.SIGABRT)

	go func() {
		<-c
		log.Println("Save history data...")
		historyData.SerializeToFile(HISTORYDATASERIAL)
		//home.Stop()
		//TODO: return relay to initial state
		os.Exit(1)
	}()

	//stop = schedule(reportSensors, time.Duration(INTERVAL)*time.Second, sensors)
	//      led.Toggle()

	work := func() {
		//defer home.Stop()
		gobot.Every(time.Duration(INTERVAL)*time.Second, func() {
			log.Println("gobot heartbeat")

			if SENSORS {
				reportSensors(sensors)
			}

		})

		gobot.On(kitchenSmokeAlarm.Event("push"), func(data interface{}) {
			log.Println("Smoke alarm in kitchen: On")
		})

		gobot.On(kitchenSmokeAlarm.Event("release"), func(data interface{}) {
			log.Println("Smoke alarm in kitchen: Off")
		})

		gobot.On(saunaSmokeAlarm.Event("push"), func(data interface{}) {
			log.Println("Smoke alarm in kitchen: On")
		})

		gobot.On(saunaSmokeAlarm.Event("release"), func(data interface{}) {
			log.Println("Smoke alarm in kitchen: Off")
		})

	}

	//TODO: more relays here
	//home.GetHeat2(),
	//home.GetRelHeatMotor2()
	robot := gobot.NewRobot("blinkBot",
		[]gobot.Connection{home.GetRelayAdaptor()},
		[]gobot.Device{home.GetRelHeat(), home.GetRelHeatPump()},
		work,
	)

	gbot.AddRobot(robot)

	go gbot.Start()

	/*
		TODO: Set initial state when app launched
	*/
	errr := home.SetHeatMode(conf.HeaterState)
	if errr != nil {
		log.Println("Cannot set Heat mode to: " + conf.HeaterState)
	}

	errr = home.SetHeatPumpMode(conf.PumpState)
	if errr != nil {
		log.Println("Cannot set HeatPump mode to: " + conf.PumpState)
	}

	reportSensors(sensors)

	err = manageHeater(&currentState)
	if err != nil {
		log.Println(err.Error())
	}

	err = managePump(&currentState)
	if err != nil {
		log.Println(err.Error())
	}
	//done TODO

	stop := scheduleBackup(backupHistoryData, time.Duration(INTERVAL*60)*time.Second, &historyData, HISTORYDATASERIAL)

	err = http.ListenAndServe(":1234", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}

	if stop != nil {
		stop <- true
	}
}

func scheduleBackup(what func(*home.HistoryData, string), delay time.Duration,
	q *home.HistoryData, l string) chan bool {
	stop := make(chan bool)

	go func() {
		for {
			what(q, l)
			select {
			case <-time.After(delay):
			case <-stop:
				return
			}
		}
	}()

	return stop
}

func backupHistoryData(q *home.HistoryData, local string) {
	historyData.SerializeToFile(local)

	if _, err := DB.UploadFile(local, "/backup/goHome.b64", true, ""); err != nil {
		log.Printf("Error uploading %s: %s\n", local, err)
	} else {
		log.Printf("File %s successfully uploaded\n", local)
	}
}
