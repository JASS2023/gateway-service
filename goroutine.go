package main

import (
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/google/uuid"
)

func startConstraint(client mqtt.Client, id uuid.UUID, typ int) {
	time.Sleep(time.Duration(15))
	if typ == 1 {
		publish_construction(client, id, true)
	}
	if typ == 2 {
		publish_service(client, id, true)
	}
}
