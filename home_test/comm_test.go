package home_test

import (
	"goHome/home"
	"log"
	"testing"
)

func TestSmsFail(t *testing.T) {
	recipients := []string{"380939760324"}
	log.Println("Start")
	err := home.Sms("TT", "hey here", recipients)
	if err != nil {
		t.Error(err)
	}
}

func TestSmsOk(t *testing.T) {
	recipients := []string{"380939760324"}
	t.Log("Stat successful test")
	err := home.Sms("TTest", "hey here", recipients)
	if err != nil {
		t.Error(err)
	}
}
