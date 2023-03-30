package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

func msgHandler(msg Message) error {
	switch msg.Type {
	case "plan_construction_site":
		log.Info("Received plan_construction_site")
		construction, ok := msg.Data.(map[string]interface{})
		coordinates := construction["coordinates"].([]interface{})
		if len(coordinates) == 0 {
			return errors.New("No coordinates found in plan_construction_site")
		}
		if !ok {
			str := fmt.Sprintf("Invalid construction data: %#v", msg.Data)
			log.Info(str)
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
		service, ok := msg.Data.(map[string]interface{})
		coordinates := service["coordinates"].([]interface{})
		if len(coordinates) == 0 {
			return errors.New("No coordinates found in plan_service")
		}
		if !ok {
			str := fmt.Sprintf("Invalid time sensitive data: %#v", msg.Data)
			log.Info(str)
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

func planService(service map[string]interface{}) error {
	id, err := uuid.Parse(service["id"].(string))
	if err != nil {
		return err
	}
	startDateTime, err := time.Parse(time.RFC3339, service["startDateTime"].(string))
	if err != nil {
		return err
	}
	endDateTime, err := time.Parse(time.RFC3339, service["endDateTime"].(string))
	if err != nil {
		return err
	}
	maximumSpeed, ok := service["maximumSpeed"].(float64)
	if !ok {
		return errors.New("Invalid maximum speed")
	}
	days, ok := service["days"].(string)
	if !ok || len(days) != 7 {
		return errors.New("Invalid days")
	}
	timeConstraints := service["time_constraints"].(interface{})
	timeConstraints1, ok := timeConstraints.(map[string]interface{})
	if !ok {
		return errors.New("Invalid time constraints")
	}
	start, ok := timeConstraints1["start"].(string)
	if !ok {
		return errors.New("Invalid time constraints")
	}
	end, ok := timeConstraints1["end"].(string)
	if !ok {
		return errors.New("Invalid time constraints")
	}
	coordinates := service["coordinates"].([]interface{})

	for i := 0; i < len(coordinates); i++ {
		log.Info(coordinates[i])
		coord, ok := coordinates[i].(map[string]interface{})
		if !ok {
			return errors.New("Error parsing coordinates")
		}
		x := coord["x"].(float64)
		y := coord["y"].(float64)
		x_abs := coord["x_abs"].(float64)
		y_abs := coord["y_abs"].(float64)
		newIssueDate := time.Now()
		newExpiryDate := endDateTime.Add((newIssueDate.Sub(startDateTime)))
		constraint := &Constraint{
			CityId:      id,
			Type:        2,
			X:           &x,
			Y:           &y,
			X_Abs:       &x_abs,
			Y_Abs:       &y_abs,
			Days:        days,
			IssueDate:   newIssueDate,
			ExpiryDate:  newExpiryDate,
			MaxSpeed:    maximumSpeed,
			StartTime:   start,
			EndTime:     end,
			Description: "Service",
		}
		//go startConstraint(*constraint) //start counting time
		err := DB.Create(constraint).Error
		if err != nil {
			log.Error(err)
			return err
		}
	}
	return nil
}

func planConstructionSite(construction map[string]interface{}) error {
	id, err := uuid.Parse(construction["id"].(string))
	if err != nil {
		return err
	}
	startDateTime, err := time.Parse(time.RFC3339, construction["startDateTime"].(string))
	if err != nil {
		return err
	}
	endDateTime, err := time.Parse(time.RFC3339, construction["endDateTime"].(string))
	if err != nil {
		return err
	}
	maximumSpeed, ok := construction["maximumSpeed"].(float64)
	if !ok {
		return errors.New("Invalid maximum speed")
	}
	traffic := construction["traffic_lights"].(interface{})
	traffic1, ok := traffic.(map[string]interface{})
	if !ok {
		return errors.New("error parsing construction data")
	}
	light1, err := uuid.Parse(traffic1["id1"].(string))
	if err != nil {
		return err
	}
	light2, err := uuid.Parse(traffic1["id2"].(string))
	if err != nil {
		return err
	}

	coordinates := construction["coordinates"].([]interface{})

	for i := 0; i < len(coordinates); i++ {
		log.Info(coordinates[i])
		coord, ok := coordinates[i].(map[string]interface{})
		if !ok {
			return errors.New("Error parsing coordinates")
		}
		x := coord["x"].(float64)
		y := coord["y"].(float64)
		x_abs := coord["x_abs"].(float64)
		y_abs := coord["y_abs"].(float64)
		newIssueDate := time.Now()
		newExpiryDate := endDateTime.Add((newIssueDate.Sub(startDateTime)))
		constraint := &Constraint{
			CityId:      id,
			Type:        1,
			X:           &x,
			Y:           &y,
			X_Abs:       &x_abs,
			Y_Abs:       &y_abs,
			IssueDate:   newIssueDate,
			ExpiryDate:  newExpiryDate,
			MaxSpeed:    maximumSpeed,
			Description: "Construction Site",
			Light1:      light1,
			Light2:      light2,
		}
		//go startConstraint(*constraint) //start counting time
		err := DB.Create(constraint).Error
		if err != nil {
			log.Error(err)
			return err
		}
	}
	return nil
}

func statusConstruction(id uuid.UUID, finished bool) (string, error) {
	var data []Constraint
	err := DB.Where("city_id =? AND type=?", id, 1).Find(&data).Error
	if err != nil {
		log.Error(err)
		return "", err
	}
	typ := "status_construction_site"
	var cd ConstructionData
	var message string
	if finished {
		message = "removed_construction_site"
	} else {
		message = "built_construction_site"
	}
	cd.Message = message
	cd.Id = id
	cd.Id = id
	for i := 0; i < len(data); i++ {
		var c Coordinates
		c.X = data[i].X
		c.Y = data[i].Y
		c.X_abs = data[i].X_Abs
		c.Y_abs = data[i].Y_Abs
		cd.Coordinates = append(cd.Coordinates, c)
	}
	cd.StartDateTime = data[0].IssueDate
	cd.EndDateTime = data[0].ExpiryDate
	cd.MaximumSpeed = data[0].MaxSpeed
	var tl TrafficLights
	tl.Id1 = data[0].Light1
	tl.Id2 = data[0].Light2
	cd.Trafficlights = tl
	msg := struct {
		Type string           `json:"type"`
		Data ConstructionData `json:"data"`
	}{
		typ, cd,
	}
	output, err := json.Marshal(msg)
	os.WriteFile("output.json", output, 0644)
	str := fmt.Sprintf("%v", data)
	log.Info(output)
	return str, nil
}

func statusService(id uuid.UUID, finished bool) (string, error) {
	var data []Constraint
	err := DB.Where("city_id =? AND type=?", id, 2).Find(&data).Error
	if err != nil {
		log.Error(err)
		return "", err
	}
	typ := "status_service"
	var cd TimeSensitiveData
	var message string
	if finished {
		message = "removed_service"
	} else {
		message = "built_service"
	}
	cd.Message = message
	cd.Id = id
	cd.Id = id
	for i := 0; i < len(data); i++ {
		var c Coordinates
		c.X = data[i].X
		c.Y = data[i].Y
		c.X_abs = data[i].X_Abs
		c.Y_abs = data[i].Y_Abs
		cd.Coordinates = append(cd.Coordinates, c)
	}
	cd.StartDateTime = data[0].IssueDate
	cd.EndDateTime = data[0].ExpiryDate
	cd.MaximumSpeed = data[0].MaxSpeed
	cd.Days = data[0].Days
	var tl TimeConstr
	tl.Start = data[0].StartTime
	tl.End = data[0].EndTime
	cd.TimeConstraints = tl
	msg := struct {
		Type string            `json:"type"`
		Data TimeSensitiveData `json:"data"`
	}{
		typ, cd,
	}
	output, err := json.Marshal(msg)
	os.WriteFile("output.json", output, 0644)
	str := fmt.Sprintf("%v", data)
	log.Info(output)
	return str, nil
}
