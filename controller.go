package main

import (
  "fmt"
  "log"
  "net/http"
  "golang.org/x/net/websocket"
  "goHome/home"
  "time"
)

func echoHandler(ws *websocket.Conn) {
  msg := make([]byte, 512)
  for {
    n, err := ws.Read(msg)
    if err != nil {
      log.Println(err)
      return
    }

    fmt.Printf("Receive: %s\n", msg[:n])
    home.ToggleHeatMotor1()
    m, err := ws.Write(msg[:n])
    if err != nil {
      log.Println(err)
      return
    }
    fmt.Printf("Send: %s\n", msg[:m])
  }
}

func schedule(what func(), delay time.Duration) chan bool {
  stop := make(chan bool)

  go func() {
    for {
      what()
      select {
        case <-time.After(delay):
        case <-stop:
        return
      }
    }
    }()

    return stop
}


func main() {

  stop := schedule(home.ReportInternalTemp, 5*time.Second)

  http.Handle("/echo", websocket.Handler(echoHandler))
  http.Handle("/", http.FileServer(http.Dir(".")))
  err := http.ListenAndServe(":1234", nil)
  if err != nil {
    panic("ListenAndServe: " + err.Error())
  }

  stop <- true

}
