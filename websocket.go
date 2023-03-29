package main

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"golang.org/x/net/websocket"
)

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

		//msgHandler(msg)
		// err = websocket.Message.Send(ws, msg)
		// if err != nil {
		// 	fmt.Println("Error sending message:", err.Error())
		// 	break
		// }
	}
}
