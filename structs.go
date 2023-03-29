package main

import (
	"time"

	"github.com/google/uuid"
)

type Message struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}

type ConstructionData struct {
	id          uuid.UUID
	coordinates []struct {
		x     *float64
		y     *float64
		x_abs *float64
		y_abs *float64
	}
	startDateTime time.Time
	endDateTime   time.Time
	maximumSpeed  float64
	trafficlights struct {
		id1 uuid.UUID
		id2 uuid.UUID
	}
}

type TimeSensitiveData struct {
	id          uuid.UUID
	coordinates []struct {
		x     *float64
		y     *float64
		x_abs *float64
		y_abs *float64
	}
	startDateTime   time.Time
	endDateTime     time.Time
	maximumSpeed    float64
	days            string
	timeConstraints struct {
		start string
		end   string
	}
	description string
}
