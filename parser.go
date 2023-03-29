package main

import (
	"errors"
	"fmt"

	log "github.com/sirupsen/logrus"
)

func msgHandler(msg Message) error {
	switch msg.Type {
	case "plan_construction_site":
		log.Info("Received plan_construction_site")
		construction, ok := msg.Data.(ConstructionData)
		if !ok {
			fmt.Println("Error parsing construction data")
			return errors.New("error parsing construction data")
		}
		err := planConstructionSite(construction)
		if err != nil {
			log.Error(err)
			return err
		}
	case "plan_service":
		log.Info("Received plan_service")
		service, ok := msg.Data.(TimeSensitiveData)
		if !ok {
			fmt.Println("Error parsing time sensitive data")
			return errors.New("error parsing time sensitive data")
		}
		err := planService(service)
		if err != nil {
			log.Error(err)
			return err
		}
	default:
		log.Info("Received unknown message")
	}
	return nil
}

func planService(service TimeSensitiveData) error {
	id := service.id
	startDateTime := service.startDateTime
	endDateTime := service.endDateTime
	maximumSpeed := service.maximumSpeed
	description := service.description
	days := service.days
	startTime := service.timeConstraints.start
	endTime := service.timeConstraints.end
	for i := 0; i < len(service.coordinates); i++ {
		x := service.coordinates[i].x
		y := service.coordinates[i].y
		x_abs := service.coordinates[i].x_abs
		y_abs := service.coordinates[i].y_abs
		constraint := &Constraint{
			CityId:      id,
			Type:        2,
			X:           x,
			Y:           y,
			X_Abs:       x_abs,
			Y_Abs:       y_abs,
			Days:        days,
			IssueDate:   startDateTime,
			ExpiryDate:  endDateTime,
			MaxSpeed:    maximumSpeed,
			StartTime:   startTime,
			EndTime:     endTime,
			Description: description,
		}
		err := DB.Create(constraint).Error
		if err != nil {
			log.Error(err)
			return err
		}
	}
	return nil
}

func planConstructionSite(construction ConstructionData) error {
	id := construction.id
	startDateTime := construction.startDateTime
	endDateTime := construction.endDateTime
	maximumSpeed := construction.maximumSpeed
	light1 := construction.trafficlights.id1
	light2 := construction.trafficlights.id2
	for i := 0; i < len(construction.coordinates); i++ {
		x := construction.coordinates[i].x
		y := construction.coordinates[i].y
		x_abs := construction.coordinates[i].x_abs
		y_abs := construction.coordinates[i].y_abs
		constraint := &Constraint{
			CityId:      id,
			Type:        1,
			X:           x,
			Y:           y,
			X_Abs:       x_abs,
			Y_Abs:       y_abs,
			IssueDate:   startDateTime,
			ExpiryDate:  endDateTime,
			MaxSpeed:    maximumSpeed,
			Light1:      light1,
			Light2:      light2,
			Description: "Construction Site",
		}
		err := DB.Create(constraint).Error
		if err != nil {
			log.Error(err)
			return err
		}
	}
	return nil
}
