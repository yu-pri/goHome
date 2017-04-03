package home

import (
	"log"

	"github.com/hybridgroup/gobot/platforms/gpio"
	"github.com/hybridgroup/gobot/platforms/raspi"
)

//"github.com/hybridgroup/gobot"

type irrigationRelay struct {
	RelayMode string `json:"RelayMode, string"`
	Relay     *raspi.RaspiAdaptor
}

/* Ir - irrigation relay type */
type Ir irrigationRelay

/*
Stop - Set relays to default position
*/
func (r *Ir) Stop() {
	log.Println("Set relay to default state")

}

/*
GetRelHeat returns led driver reference to RelayMotor1
*/
func (r *Ir) GetIr() *gpio.LedDriver {
	return r.Relay
}
