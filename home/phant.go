package home

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

/*
ReportInternalTemp Reports data to spark
*/
func ReportInternalTemp(v float32) error {
	url := cloudURL + phantPublicTemp + "?internaltemp=" + fmt.Sprintf("%f", v)
	fmt.Println("URL:>", url)
	var jsonStr = []byte(``)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return err
	}

	req.Header.Set("Phant-Private-Key", phantPrivateTemp)
	//req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return err
	}

	defer resp.Body.Close()
	status := resp.StatusCode
	log.Println("response Status:", status)
	log.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	log.Println("response Body:", string(body))
	if status != 200 {
		return errors.New("Non OK response from sparkfun: " + string(body))
	}

	return nil
}
