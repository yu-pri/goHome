package main

import (
	"encoding/json"
	"goHome/home"
	"io/ioutil"
	"log"
	"os"
)

type config struct {
	PumpState   string `json:"PumpState, string"`
	HeaterState string `json:"HeaterState, string"`
	DesiredTemp string `json:"DesiredTemp, number"`
}

/*Config contains user preferences*/
type Config config

/*
ToJSON returns serialized date
*/
func (q *Config) ToJSON() (d []byte, err error) {
	//now := int(time.Now().Unix())

	b, err := json.Marshal(q)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func (q *Config) saveConfig() error {
	b, err := json.Marshal(q)
	if err != nil {
		log.Println(err)
		return err
	}

	err = ioutil.WriteFile(configFileName, b, 0644)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (q *Config) setDefault() {
	q.HeaterState = home.AUTO
	q.PumpState = home.AUTO
}

func (q *Config) loadConfig() error {

	configFile, err := os.Open(configFileName)
	if err != nil {
		log.Println("Cannot open config: ", err.Error())
		q.setDefault()
		return err
	}

	jsonParser := json.NewDecoder(configFile)
	if err = jsonParser.Decode(q); err != nil {
		log.Println("Error parsing config: ", err.Error())
		q.setDefault()
		return err
	}

	return nil
}
