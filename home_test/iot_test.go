package home_test

import (
	"goHome/home"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReportTempInternal(t *testing.T) {
	got := home.IOTReportTempInternal(22.5)
	//t.Log(got)
	if got != nil {
		t.Error(got)
	}
}

func TestHistoryToJSON(t *testing.T) {
	tt := &home.HData{}
	tt.TempOutside = 10
	//var h home.HistoryData
	h := home.HistoryData{}
	h.Push(tt)

	tt = &home.HData{}
	tt.TempOutside = 12
	h.Push(tt)

	b, err := h.ToJSON(0)
	if err != nil {
		t.Error(err)
	}
	log.Println(string(b))
	//t.Log(got)
	//if got != nil {
	//	t.Error(got)
	//}
}

func TestHistoryDataEncode(t *testing.T) {
	tt := &home.HData{}
	tt.TempOutside = 10
	//var h home.HistoryData
	h := home.HistoryData{}
	h.Push(tt)

	tt = &home.HData{}
	tt.TempOutside = 12
	h.Push(tt)

	fname := "/tmp/hdata.b64"
	err := h.SerializeToFile(fname)
	if err != nil {
		t.Error(err)
	}
	tt = &home.HData{}
	tt.TempOutside = 14
	h.Push(tt)

	h.RestoreFromFile(fname)
	/*
		if err != nil {
			t.Error(err)
		}
	*/
	assert.Equal(t, 2, h.Len())

}
