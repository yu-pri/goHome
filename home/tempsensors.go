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

func NewSensors() (*Sensors, error) {
  s, err := goDS18B20.New()
  if (err != nil) {
    return nil, err
  }

  err = s.AssignAlias("int", "28-0315a14596ff");

  return &Sensors{
    Temp: *s,
  }, err
}

func (s Sensors)InternalSensor () (float32, error) {
  err := s.Temp.Update()
  if (err != nil) {
    return -100, err
  }

  temp, err := s.Temp.ReadSingleAlias("int");

  return temp.Celsius(), err
}
