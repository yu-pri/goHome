package home

import (
	"log"

	"github.com/traetox/goDS18B20"
)

var (
//Sensors goDS18B20.ProbeGroup
)

/*
Sensors Reference to sensor
*/
type Sensors struct {
	Temp goDS18B20.ProbeGroup
}

func init() {
	log.Println("init sensors")

}

/*
NewSensors Creates reference to sensor
*/
func NewSensors() (*Sensors, error) {
	s, err := goDS18B20.New()
	if err != nil {
		return nil, err
	}

	err = s.AssignAlias("int", internalSensorID)
	err = s.AssignAlias("entry", entryRoomSensorID)
	err = s.AssignAlias("rev", reverseSensorID)

	//TODO: More sensors should be declared here

	return &Sensors{
		Temp: *s,
	}, err
}

/*
InternalSensor returns temperature from internal sensor
*/
func (s Sensors) InternalSensor() (float32, error) {
	err := s.Temp.Update()
	if err != nil {
		return -100, err
	}

	temp, err := s.Temp.ReadSingleAlias("int")

	return temp.Celsius(), err
}

/*
ReverseSensor returns temperature from internal sensor
*/
func (s Sensors) ReverseSensor() (float32, error) {
	err := s.Temp.Update()
	if err != nil {
		return -100, err
	}

	temp, err := s.Temp.ReadSingleAlias("rev")

	return temp.Celsius(), err
}

/*
EntryRoomSensor returns temperature from internal sensor
*/
func (s Sensors) EntryRoomSensor() (float32, error) {
	err := s.Temp.Update()
	if err != nil {
		return -100, err
	}

	temp, err := s.Temp.ReadSingleAlias("entry")

	return temp.Celsius(), err
}
