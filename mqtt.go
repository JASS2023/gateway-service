package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	log "github.com/sirupsen/logrus"
)

var MQTTClient mqtt.Client

func handle_sub(_ mqtt.Client, msg mqtt.Message) {
	log.Infof("%s", msg.Topic())
	log.Infof("%s", msg.Payload())
}

func handle_pub(_ mqtt.Client, msg mqtt.Message) {
	log.Infof("%s", msg.Topic())
	log.Infof("%s", msg.Payload())
}

func InitMQTTClient() {
	opts := mqtt.NewClientOptions().AddBroker(Cfg.MQTT.Broker).SetClientID(Cfg.MQTT.ClientID)

	opts.SetPingTimeout(time.Second)
	opts.SetConnectTimeout(time.Second)
	opts.SetWriteTimeout(time.Second)
	opts.SetKeepAlive(10)

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

	// Messages will be delivered asynchronously so we just need to wait for a signal to shutdown
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	signal.Notify(sig, syscall.SIGTERM)

	<-sig
	fmt.Println("signal caught - exiting")
	MQTTClient.Disconnect(1000)
	fmt.Println("shutdown complete")
}
