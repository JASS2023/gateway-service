package main

import (
	"strings"

	"github.com/mitchellh/mapstructure"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Configuration struct {
	MQTT struct {
		Broker   string   `yaml:"broker"`
		ClientID string   `yaml:"client_id"`
		Topics   []string `yaml:"topics"`
	} `yaml:"mqtt"`
}

var (
	Cfg Configuration
)

func LoadConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		if err == err.(viper.ConfigFileNotFoundError) {
			log.WithError(err).Warn("can't find config file")
		} else {
			log.WithError(err).Panic("fatal error config file: ", err)
		}
	}
	log.Info("Using Config file ", viper.ConfigFileUsed())

	err = viper.Unmarshal(&Cfg, func(c *mapstructure.DecoderConfig) { c.TagName = "yaml" })
	if err != nil {
		log.WithError(err).Fatal("fatal error config file: ", err)
	}
}
