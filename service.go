package main

import (
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
	"golang.org/x/net/websocket"
)

func main() {
	log.SetFormatter(&log.TextFormatter{TimestampFormat: "01.01.2006 00:00:00", FullTimestamp: true})

	log.Info("Starting service...")
	LoadConfig()

	log.Infof("%v", Cfg)

	// InitMQTTClient()
	// db := ConnectDB()

	// err := db.Create(&Constraint{StartTime: "01:00:00"}).Error
	// if err != nil {
	// 	log.Error(err)
	// }
	// http.Handle("/echo", websocket.Handler(echoHandler))
	http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		s := websocket.Server{Handler: websocket.Handler(echoHandler)}
		s.ServeHTTP(w, r)
	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err.Error())
	}
}
