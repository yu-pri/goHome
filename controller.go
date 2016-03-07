package main

import (
  "fmt"
  "log"
  "net/http"
  "golang.org/x/net/websocket"
  "goHome/home"
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

func main() {
  http.Handle("/echo", websocket.Handler(echoHandler))
  http.Handle("/", http.FileServer(http.Dir(".")))
  err := http.ListenAndServe(":1234", nil)
  if err != nil {
    panic("ListenAndServe: " + err.Error())
  }
}
