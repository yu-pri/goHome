package main

import (
	"goHome/home"
	"log"
	"testing"
	"time"

	"github.com/stacktic/dropbox"
)

/*
DB - dropbox database
*/
var DB *dropbox.Dropbox

func init() {

	// 1. Create a new dropbox object.
	DB = dropbox.NewDropbox()

	// 2. Provide your clientid and clientsecret (see prerequisite).
	DB.SetAppInfo(home.DropboxClientid, home.DropboxClientsecret)

	// 3. Provide the user token.
	DB.SetAccessToken(home.DropboxToken)
	scheduleBackup(backupHistoryData, time.Duration(INTERVAL*1)*time.Second, &historyData, HISTORYDATASERIAL)

}

func scheduleBackup(what func(*home.HistoryData, string), delay time.Duration,
	q *home.HistoryData, l string) chan bool {
	stop := make(chan bool)

	go func() {
		for {
			what(q, l)
			select {
			case <-time.After(delay):
			case <-stop:
				return
			}
		}
	}()

	return stop
}

func backupHistoryData(q *home.HistoryData, local string) {
	historyData.SerializeToFile(local)

	if _, err := DB.UploadFile(local, "/backup", true, "1"); err != nil {
		log.Printf("Error uploading %s: %s\n", local, err)
	} else {
		log.Printf("File %s successfully uploaded\n", local)
	}
}

/*
TestUploadDropbox - test case for file upload
*/
func TestUploadDropbox(t *testing.T) {
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

	backupHistoryData(&h, fname)
}
