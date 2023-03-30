package main

import (
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetFormatter(&log.TextFormatter{TimestampFormat: "2006-01-02 15:04:05", FullTimestamp: true})

	log.Info("Starting service...")
	LoadConfig()

	log.Infof("%v", Cfg)

	err := ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	InitMQTTClient()
}
