package home_test

import (
	"goHome/home"
	"testing"
)

func TestReportTempInternal(t *testing.T) {
	got := home.IOTReportTempInternal(22.5)
	//t.Log(got)
	if got != nil {
		t.Error(got)
	}
}
