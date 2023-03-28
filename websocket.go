package main

import (
	"fmt"
	"time"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/websocket"
)

type Message struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}


type ConstructionData struct {
	id int
	coordinates [] struct {
		x int
		y int
		quadrant int
	}
	startDateTime time.Time
	endDateTime time.Time
	maximumSpeed float64
	trafficlights struct{
		id1 int
		id2 int
	}
}


type TimeSensitiveData struct {
	id int
	coordinates []struct {
		x int
		y int
		quadrant int
	}
	startDateTime time.Time
	endDateTime time.Time
	maximumSpeed float64
	days string
	timeConstraints struct {
		start string
		end string
	}
	description string
}

func echoHandler(ws *websocket.Conn) {
	defer ws.Close()
	log.Info("New connection")
	for {
		var msg Message
		err := websocket.JSON.Receive(ws, &msg)
		if err != nil {
			fmt.Println("Error receiving message:", err.Error())
			break
		}
		fmt.Println("Received message:", msg)

		switch msg.Type {
		case "plan_construction_site":
			log.Info("Received plan_construction_site")
			constraction,ok := msg.Data.(ConstructionData)
			if!ok {
                fmt.Println("Error parsing construction data")
                break
            }
			fmt.Println("Construction data id:", constraction.id)
			var id int = constraction.id;
			var startDateTime time.Time = constraction.startDateTime;
			var endDateTime time.Time = constraction.endDateTime;
			var maximumSpeed float64 = constraction.maximumSpeed;
			var description string = constraction.description;
			db := ConnectDB()
			for i := 0; i < len(constraction.coordinates); i++ {
				var x int = constraction.coordinates[i].x;
				var y int = constraction.coordinates[i].y;
				var quadrant int = constraction.coordinates[i].quadrant;
				err := db.Create(&Constraint{
					Type=1, 
					Quadrant=quadrant, 
					X=x, 
					Y=y, 
					IssueDate = startDateTime, 
					ExpiryDate = endDateTime, 
					MaxSpeed = maximumSpeed, 
					Description = "Construction Site", 
					}).Error;
				if err != nil {
					log.Error(err)
				}
			}
		case "plan_service":
			log.Info("Received plan_service")
			service,ok := msg.Data.(TimeSensitiveData)
			if!ok {
                fmt.Println("Error parsing time sensitive data")
                break
            }
			fmt.Println("Time sensitive data id:", service.id)
			var id int = service.id;
			var startDateTime time.Time = service.startDateTime;
			var endDateTime time.Time = service.endDateTime;
			var maximumSpeed float64 = service.maximumSpeed;
			var description string = service.description;
			var startTime string = service.timeConstraints.start;
			var endTime string = service.timeConstraints.end;
			db := ConnectDB()
			for i := 0; i < len(service.coordinates); i++ {
				var x int = service.coordinates[i].x;
				var y int = service.coordinates[i].y;
				var quadrant int = service.coordinates[i].quadrant;
				err := db.Create(&Constraint{
					Type=2, 
                    Quadrant=quadrant, 
                    X=x, 
                    Y=y, 
                    IssueDate = startDateTime, 
                    ExpiryDate = endDateTime, 
                    MaxSpeed = maximumSpeed, 
					StartTime = startTime, 
					EndTime = endTime, 
                    Description = description, 
                    }).Error;
                if err!= nil {
					log.Error(err)
				}
			}
		default:
			log.Info("Received unknown message")
		}
		// err = websocket.Message.Send(ws, msg)
		// if err != nil {
		// 	fmt.Println("Error sending message:", err.Error())
		// 	break
		// }
	}
}
