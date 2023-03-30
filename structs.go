package main

import (
	"time"

	"github.com/google/uuid"
)

type Message struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}

type Coordinates struct {
	X     *float64 `json:"x"`
	Y     *float64 `json:"y"`
	X_abs *float64 `json:"x_abs"`
	Y_abs *float64 `json:"y_abs"`
	Quadrants []uint `json:"quadrants"`
}

type TrafficLights struct {
	Id1 uuid.UUID `json:"id1"`
	Id2 uuid.UUID `json:"id2"`
}

type ConstructionData struct {
	Message       string        `json:"message"`
	Id            uuid.UUID     `json:"id"`
	Coordinates   []Coordinates `json:"coordinates"`
	StartDateTime time.Time     `json:"startDateTime"`
	EndDateTime   time.Time     `json:"endDateTime"`
	MaximumSpeed  float64       `json:"maximumSpeed"`
	Trafficlights TrafficLights `json:"trafficlights"`
}

type TimeConstr struct {
	Start string `json:"start"`
	End   string `json:"end"`
}

type TimeSensitiveData struct {
	Message         string        `json:"message"`
	Id              uuid.UUID     `json:"id"`
	Coordinates     []Coordinates `json:"coordinates"`
	StartDateTime   time.Time     `json:"startDateTime"`
	EndDateTime     time.Time     `json:"endDateTime"`
	MaximumSpeed    float64       `json:"maximumSpeed"`
	Days            string        `json:"days"`
	TimeConstraints TimeConstr    `json:"time_constraints"`
}
