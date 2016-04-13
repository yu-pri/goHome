package home_test

import (
	"goHome/home"
	"log"
	"testing"
)

func TestReportTempInternal(t *testing.T) {
	got := home.IOTReportTempInternal(22.5)
	//t.Log(got)
	if got != nil {
		t.Error(got)
	}
}

func TestHistoryToJSON(t *testing.T) {
	var tt = &home.HData{}
	tt.TempOutside = 10
	//var h home.HistoryData
	h := home.HistoryData{}
	h.Push(tt)
	b, err := h.ToJSON()
	if err != nil {
		t.Error(err)
	}
	log.Println(string(b))
	//t.Log(got)
	//if got != nil {
	//	t.Error(got)
	//}
}
