package main

import (
	"encoding/json"
	"os"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetFormatter(&log.TextFormatter{TimestampFormat: "01.01.2006 00:00:00", FullTimestamp: true})

	log.Info("Starting service...")
	LoadConfig()

	log.Infof("%v", Cfg)

	//InitMQTTClient()
	ConnectDB()
	dat, err := os.ReadFile("jsonfile1.json")
	if err != nil {
		log.Fatal(err)
	}
	var msg Message
	json.Unmarshal(dat, &msg)
	msgHandler(msg)
	input, err := uuid.Parse("273b62ad-a99d-48be-8d80-ccc55ef688b5")
	statusService(input, false)

	// err := db.Create(&Constraint{StartTime: "01:00:00"}).Error
	// if err != nil {
	// 	log.Error(err)
	// }
	// http.Handle("/echo", websocket.Handler(echoHandler))
	// http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
	// 	s := websocket.Server{Handler: websocket.Handler(echoHandler)}
	// 	s.ServeHTTP(w, r)
	// })
	// err := http.ListenAndServe(":8080", nil)
	// if err != nil {
	// 	fmt.Println("Error starting server:", err.Error())
	// }
}
