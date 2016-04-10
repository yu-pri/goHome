package home

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

/*
IOTReportValue Reports data to spark
*/
func IOTReportValue(dev string, name string, v float32) error {
	url := iotURL + "/device/" + dev + "/var/" + name + "?apikey=" + iotCloudAPIKey
	fmt.Println("URL:>", url)

	/*
		var jsonStr = []byte(`{"value":` + fmt.Sprintf("%f", v) + "}")

		req, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonStr))
	*/
	//var jsonStr = []byte()
	//fmt.Sprintf("%f", v)
	jsonStr := "{\"value\":" + fmt.Sprintf("%f", v) + "}"
	req, err := http.NewRequest("PUT", url, strings.NewReader(jsonStr))
	req.Header.Add("Content-Type", "application/json")

	if err != nil {
		return err
	}

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
		return errors.New("Non OK response from IOT: " + string(body))
	}
	return nil
}

/*
IOTReportTempInternal Reports data to spark
*/
func IOTReportTempInternal(v float32) error {
	err := IOTReportValue(iotCloudHomeDevice, iotCloudInternalTempVar, v)
	return err
}
