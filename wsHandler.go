package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"goHome/home"
	"log"
	"sync"
	"time"

	"golang.org/x/net/websocket"
)

/*
MsgCommand - simple comm protocol to deliver infor to server
*/
type MsgCommand struct {
	sync.Mutex
	Command string `json:"command, string"`
	Object  string `json:"object, string"`
	Param1  string `json:"param1, string"`
	Param2  string `json:"param2, string"`
}

func echoHandler(ws *websocket.Conn) {

	var id = int32(time.Now().Unix())

	conns.lock.Lock()
	conns.ws[id] = ws
	conns.lock.Unlock()
	defer func() {
		conns.lock.Lock()
		delete(conns.ws, id)
		conns.lock.Unlock()
	}()

	msg := make([]byte, 512)
	for {

		n, err := ws.Read(msg)
		if err != nil {
			log.Println(err)
			return
		}

		x := new(MsgCommand)
		fmt.Printf("Receive: %s\n", msg[:n])
		if err := json.NewDecoder(bytes.NewReader(msg)).Decode(x); err == nil {
			execute(x)
		}

		/*
			m, err := ws.Write(msg[:n])
			if err != nil {
				log.Println(err)
				return
			}
			fmt.Printf("Send: %s\n", msg[:m])
		*/
	}
}

func execute(c *MsgCommand) error {
	defer c.Unlock()
	c.Lock()

	if c.Object == home.CommandOnPumpr1 {
		return home.OnHeatMotor1()
	}

	if c.Object == home.CommandOffPumpr1 {
		return home.OffHeatMotor1()
	}

	return nil
}
