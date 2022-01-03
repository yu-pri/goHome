package main

// import (
// 	"goHome/home"

// 	"github.com/stacktic/dropbox"
// )

// /*
// DB - dropbox database
// */
// var DB *dropbox.Dropbox

// func init() {

// 	// 1. Create a new dropbox object.
// 	DB = dropbox.NewDropbox()

// 	// 2. Provide your clientid and clientsecret (see prerequisite).
// 	DB.SetAppInfo(home.DropboxClientid, home.DropboxClientsecret)

// 	// 3. Provide the user token.
// 	DB.SetAccessToken(home.DropboxToken)

// }

// /*
// TestUploadDropbox - test case for file upload
// */
// /*
// func TestUploadDropbox(t *testing.T) {
// 	tt := &home.HData{}
// 	tt.TempOutside = 10
// 	//var h home.HistoryData
// 	h := home.HistoryData{}
// 	h.Push(tt)

// 	tt = &home.HData{}
// 	tt.TempOutside = 12
// 	h.Push(tt)

// 	fname := "/tmp/hdata.b64"
// 	err := h.SerializeToFile(fname)
// 	if err != nil {
// 		t.Error(err)
// 	}

// 	backupHistoryData(&h, fname)
// }
// */
