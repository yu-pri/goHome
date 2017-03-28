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

	s.AssignAlias("int", internalSensorID)
	s.AssignAlias("entry", entryRoomSensorID)
	s.AssignAlias("rev", reverseSensorID)
	s.AssignAlias("heater", heaterSensorID)
	//err = s.AssignAlias("waterBoiler", waterBoilerSensorID)
	//err = s.AssignAlias("recuperator", recuperatorSensorID)
	//err = s.AssignAlias("outside", outsideSensorID)

	//TODO: More sensors should be declared here

	s.Update()

	return &Sensors{
		Temp: *s,
	}, err
}

/*
InternalSensor returns temperature from internal sensor
*/
func (s Sensors) InternalSensor() (float32, error) {
	//err := s.Temp.Update()

	temp, err := s.Temp.ReadSingleAlias("int")

	return temp.Celsius(), err
}

/*
RecuperatorSensor returns temperature from the sensor
*/
/*
func (s Sensors) RecuperatorSensor() (float32, error) {
	//err := s.Temp.Update()

	temp, err := s.Temp.ReadSingleAlias("recuperator")
	//c := float64(temp.Celsius())
	return Round32(temp.Celsius(), 0.5, 1), err
}
*/

/*
OutsideSensor returns temperature from the sensor
*/
/*
func (s Sensors) OutsideSensor() (float32, error) {
	//err := s.Temp.Update()

	temp, err := s.Temp.ReadSingleAlias("outside")

	return Round32(temp.Celsius(), 0.5, 1), err
}
*/

/*
WaterBoilerSensor returns temperature from the sensor
*/
func (s Sensors) WaterBoilerSensor() (float32, error) {
	//err := s.Temp.Update()

	temp, err := s.Temp.ReadSingleAlias("waterboiler")

	return Round32(temp.Celsius(), 0.5, 0), err
}

/*
ReverseSensor returns temperature from internal sensor
*/
func (s Sensors) ReverseSensor() (float32, error) {
	//err := s.Temp.Update()

	temp, err := s.Temp.ReadSingleAlias("rev")

	return Round32(temp.Celsius(), 0.5, 0), err
}

/*
EntryRoomSensor returns temperature from internal sensor
*/
func (s Sensors) EntryRoomSensor() (float32, error) {
	//err := s.Temp.Update()

	temp, err := s.Temp.ReadSingleAlias("entry")

	return Round32(temp.Celsius(), 0.5, 0), err
}

/*
HeaterSensor returns temperature from internal sensor
*/
func (s Sensors) HeaterSensor() (float32, error) {
	//err := s.Temp.Update()

	temp, err := s.Temp.ReadSingleAlias("heater")

	return Round32(temp.Celsius(), 0.5, 0), err
}
