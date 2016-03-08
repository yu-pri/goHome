package home

import (
  "log"
  //"errors"
  //"os"
  //"time"
  //  "sync"
  "github.com/traetox/goDS18B20"
  )

var (
  //Sensors goDS18B20.ProbeGroup
)

type Sensors struct {
  Temp goDS18B20.ProbeGroup
}

func init () {
  log.Println("init sensors")

}

func New() (*Sensors, error) {
  s, err := goDS18B20.New()
  if (err != nil) {
    return nil, err
  }

  err = s.AssignAlias("int", "28-0315a14596ff");

  return &Sensors{
    Temp: *s,
  }, err
}

func (s Sensors)ReportInternalTemp () {
  err := s.Temp.Update()
  if (err != nil) {
    log.Fatal(err)
  }

  temp, err := s.Temp.ReadSingleAlias("int");
  if (err != nil) {
    log.Fatal(err)
  }

  log.Println(temp.Celsius())
}
