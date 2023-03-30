package main

// import (
// 	"time"

// 	mqtt "github.com/eclipse/paho.mqtt.golang"
// 	"github.com/google/uuid"
// 	log "github.com/sirupsen/logrus"
// )

// func startConstraint(client mqtt.Client, id uuid.UUID) {
// 	var c Constraint
// 	DB.Where("city_id =?", id).First(&c)
// 	if c.IssueDate.Before(time.Now()) {
// 		log.Error("constraint has already expired")
// 	}
// 	duration := c.ExpiryDate.Sub(c.IssueDate)
// 	time.Sleep(time.Duration(duration.Seconds()))
// 	if c.Type == 1 {
// 		// publish_construction(client, id, true)
// 	}
// 	if c.Type == 2 {
// 		// publish_service(client, id, true)
// 	}
// }
