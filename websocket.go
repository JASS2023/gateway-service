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
			id = constraction.id;
			startDateTime = constraction.startDateTime;
			endDateTime = constraction.endDateTime;
			maximumSpeed = constraction.maximumSpeed;
			description = constraction.description;
			db := ConnectDB()
			for i := 0; i < len(constraction.coordinates); i++ {
				x = constraction.coordinates[i].x;
				y = constraction.coordinates[i].y;
				quadrant = constraction.coordinates[i].quadrant;
				constraint = &Constraint{
					Type=1, 
					Quadrant=quadrant, 
					X=x, 
					Y=y, 
					IssueDate = startDateTime, 
					ExpiryDate = endDateTime, 
					MaxSpeed = maximumSpeed, 
					Description = "Construction Site", 
					};
				err := db.Create(constraint).Error;
				if err != nil {
					log.Error(err)
				}
				b, err := json.Marshal(constraint)
				if err != nil {
					fmt.Println(err)
					return
				}
				fmt.Println(string(b))
			}
		case "plan_service":
			log.Info("Received plan_service")
			service,ok := msg.Data.(TimeSensitiveData)
			if!ok {
                fmt.Println("Error parsing time sensitive data")
                break
            }
			fmt.Println("Time sensitive data id:", service.id)
			id = service.id;
			startDateTime = service.startDateTime;
			endDateTime = service.endDateTime;
			maximumSpeed = service.maximumSpeed;
			description = service.description;
			startTime = service.timeConstraints.start;
			endTime = service.timeConstraints.end;
			db := ConnectDB()
			for i := 0; i < len(service.coordinates); i++ {
				x = service.coordinates[i].x;
				y = service.coordinates[i].y;
				quadrant = service.coordinates[i].quadrant;
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
				b, err := json.Marshal(constraint)
				if err != nil {
					fmt.Println(err)
					return
				}
				fmt.Println(string(b))
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
