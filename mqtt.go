package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

var MQTTClient mqtt.Client

const (
	AT_MOST_ONCE  = 0
	AT_LEAST_ONCE = 1
	EXACTLY_ONCE  = 2
)

func InitMQTTClient() {
	opts := mqtt.NewClientOptions().AddBroker(Cfg.MQTT.Broker).SetClientID(Cfg.MQTT.ClientID)

	opts.SetPingTimeout(time.Second)
	opts.SetConnectTimeout(time.Second)
	opts.SetWriteTimeout(time.Second)
	opts.SetKeepAlive(10 * time.Second)

	opts.SetAutoReconnect(true)
	opts.SetConnectRetry(true)

	opts.OnConnect = func(c mqtt.Client) {
		log.Info("connection established")
	}

	opts.OnConnectionLost = func(c mqtt.Client, err error) {
		log.WithError(err).Warn("connection lost")
	}

	opts.OnReconnecting = func(c mqtt.Client, opts *mqtt.ClientOptions) {
		log.Warn("attempting to reconnect")
	}

	MQTTClient := mqtt.NewClient(opts)
	if token := MQTTClient.Connect(); token.Wait() && token.Error() != nil {
		log.WithError(token.Error()).Panic(token.Error())
	}

	// Listen to construction planning
	MQTTClient.Subscribe("construction/+/plan", EXACTLY_ONCE, handle_construction_plan)
	MQTTClient.Subscribe("service/+/plan", EXACTLY_ONCE, handle_service_plan)
	// Duckie information
	MQTTClient.Subscribe("duckie/+/status", EXACTLY_ONCE, handle_duckie_status)
	MQTTClient.Subscribe("obstruction/+/status", EXACTLY_ONCE, handle_obstacle_status)

	log.Info("MQTT client initialized")

	// Messages will be delivered asynchronously so we just need to wait for a signal to shutdown
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	signal.Notify(sig, syscall.SIGTERM)

	<-sig
	fmt.Println("signal caught - exiting")
	MQTTClient.Disconnect(1000)
	fmt.Println("shutdown complete")
}

func handle_construction_plan(client mqtt.Client, msg mqtt.Message) {
	log.Info("construction plan received")
	// TODO - handle construction planning and store it in the database
	log.Infof("%s", msg.Topic())
	log.Infof("%s", msg.Payload())

	time.Sleep(10 * time.Second)
	publish_construction(client, uuid.New())
}

func handle_service_plan(client mqtt.Client, msg mqtt.Message) {
	log.Info("service plan received")
	// TODO - handle construction planning and store it in the database
	log.Infof("%s", msg.Topic())
	log.Infof("%s", msg.Payload())

	// Simulate the placement of a service
	time.Sleep(10 * time.Second)
	publish_service(client, uuid.New())
}

func handle_duckie_status(client mqtt.Client, msg mqtt.Message) {
	log.Info("duckie status received")
	// TODO - handle construction planning and store it in the database
	log.Infof("%s", msg.Topic())
	log.Infof("%s", msg.Payload())
}

func handle_obstacle_status(client mqtt.Client, msg mqtt.Message) {
	log.Info("obstacle status received")
	// TODO - handle construction planning and store it in the database
	log.Infof("%s", msg.Topic())
	log.Infof("%s", msg.Payload())
}

func publish_construction(client mqtt.Client, id uuid.UUID) {
	topic := fmt.Sprintf("construction/%s/plan", id)
	// TODO add JSON payload
	client.Publish(topic, EXACTLY_ONCE, true, "test")
}

func publish_service(client mqtt.Client, id uuid.UUID) {
	topic := fmt.Sprintf("service/%s/plan", id)
	// TODO add JSON payload
	client.Publish(topic, EXACTLY_ONCE, true, "test")
}
