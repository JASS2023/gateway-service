package main

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"golang.org/x/net/websocket"
)

type Message struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}

func echoHandler(ws *websocket.Conn) {
	defer ws.Close()
	log.Info("New connection")
	for {
		var msg Message
		err := websocket.JSON.Receive(ws, &msg)
		if err != nil {
			fmt.Println("Error receiving message:", err.Error())
			break
		}
		fmt.Println("Received message:", msg)

		switch msg.Type {
		case "plan_construction_site":
			log.Info("Received plan_construction_site")
		case "plan_service":
			log.Info("Received plan_service")
		default:
			log.Info("Received unknown message")
		}
		// err = websocket.Message.Send(ws, msg)
		// if err != nil {
		// 	fmt.Println("Error sending message:", err.Error())
		// 	break
		// }
	}
}
