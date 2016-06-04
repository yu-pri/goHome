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
IOTReportValue Reports data to cloud
*/
func IOTReportValue(server string, key string, dev string, name string, v float32) error {
	url := server + "/device/" + dev + "/var/" + name + "?apikey=" + key
	log.Println("URL:>", url)

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
IOTReportBoolValue Reports data to cloud
*/
func IOTReportBoolValue(server string, key string, dev string, name string, v bool) error {
	url := server + "/device/" + dev + "/var/" + name + "?apikey=" + key
	log.Println("URL:>", url)

	jsonStr := "{\"value\":" + fmt.Sprintf("%t", v) + "}"
	log.Println(jsonStr)

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
	err := IOTReportValue(iotURL, iotCloudAPIKey, iotCloudHomeDevice, iotCloudInternalTempVar, v)
	return err
}

/*
IOTReportDev Reports data to spark
*/
func IOTReportDev(h HData) error {
	err := IOTReportValue(iotDevURL, iotDevCloudAPIKey, iotDevCloudHomeDevice,
		iotCloudEntry, h.TempEntryRoom)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	err = IOTReportValue(iotDevURL, iotDevCloudAPIKey, iotDevCloudHomeDevice,
		iotCloudKitchen, h.TempInside)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	err = IOTReportValue(iotDevURL, iotDevCloudAPIKey, iotDevCloudHomeDevice,
		iotCloudHout, h.TempHeater)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	err = IOTReportValue(iotDevURL, iotDevCloudAPIKey, iotDevCloudHomeDevice,
		iotCloudHin, h.TempReverse)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	err = IOTReportBoolValue(iotDevURL, iotDevCloudAPIKey, iotDevCloudHomeDevice,
		iotCloudHeater, h.HeaterState)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	err = IOTReportBoolValue(iotDevURL, iotDevCloudAPIKey, iotDevCloudHomeDevice,
		iotCloudPump, h.PumpState)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return err
}
