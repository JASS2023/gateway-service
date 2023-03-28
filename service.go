package main

import (
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetFormatter(&log.TextFormatter{TimestampFormat: "01.01.2006 00:00:00", FullTimestamp: true})

	log.Info("Starting service...")
	LoadConfig()

	log.Infof("%v", Cfg)

	InitMQTTClient()
}
