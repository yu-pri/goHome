package home

import (
  "log"
  //"os"
  //"time"
  //  "sync"
  "github.com/traetox/goDS18B20"
  )

var (
  Sensors goDS18B20.ProbeGroup
)

func init () {
  log.Println("init sensors")
  Sensors, err := goDS18B20.New()
  if (err != nil) {
    log.Fatal(err)
  }

  err = Sensors.AssignAlias("int", "28-0315a14596ff");
  if (err != nil) {
    log.Fatal(err)
  }

}

func ReportInternalTemp () {
  err := Sensors.Update()
  if (err != nil) {
    log.Fatal(err)
  }

  temp, err := Sensors.ReadSingleAlias("int");
  if (err != nil) {
    log.Fatal(err)
  }

  log.Println(temp.Celsius())

}
