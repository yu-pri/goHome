package home

import (
	"log"

	"github.com/traetox/goDS18B20"
)

var (
//Sensors goDS18B20.ProbeGroup
)

/*
HeaterOnThreshold - getting cold and need heater
*/
var HeaterOnThreshold = 22.0

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
	err = s.AssignAlias("heater", heaterSensorID)
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

	return temp.Celsius(), err
}
*/

/*
OutsideSensor returns temperature from the sensor
*/
/*
func (s Sensors) OutsideSensor() (float32, error) {
	//err := s.Temp.Update()

	temp, err := s.Temp.ReadSingleAlias("outside")

	return temp.Celsius(), err
}
*/

/*
WaterBoilerSensor returns temperature from the sensor
*/
func (s Sensors) WaterBoilerSensor() (float32, error) {
	//err := s.Temp.Update()

	temp, err := s.Temp.ReadSingleAlias("waterboiler")

	return temp.Celsius(), err
}

/*
ReverseSensor returns temperature from internal sensor
*/
func (s Sensors) ReverseSensor() (float32, error) {
	//err := s.Temp.Update()

	temp, err := s.Temp.ReadSingleAlias("rev")

	return temp.Celsius(), err
}

/*
EntryRoomSensor returns temperature from internal sensor
*/
func (s Sensors) EntryRoomSensor() (float32, error) {
	//err := s.Temp.Update()

	temp, err := s.Temp.ReadSingleAlias("entry")

	return temp.Celsius(), err
}

/*
HeaterSensor returns temperature from internal sensor
*/
func (s Sensors) HeaterSensor() (float32, error) {
	//err := s.Temp.Update()

	temp, err := s.Temp.ReadSingleAlias("heater")

	return temp.Celsius(), err
}
