package main

import (
  "fmt"
  "net/http"
  "bytes"
  "io/ioutil"
  //"strconv"
)

func main() {
    url := "http://data.sparkfun.com/input/rozL4pN5jMC466MKl8AE.json?internaltemp=" + fmt.Sprintf("%f", 23.5000)
    fmt.Println("URL:>", url)
    var jsonStr = []byte(``)

    req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
    req.Header.Set("Phant-Private-Key", "jkybBZw5j9F0xxmwWar1")
    //req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    fmt.Println("response Status:", resp.Status)
    fmt.Println("response Headers:", resp.Header)
    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println("response Body:", string(body))
}
