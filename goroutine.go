package main

import (
	"errors"
	"time"
)

func startConstraint(c Constraint) error {
	if c.IssueDate.Before(time.Now()) {
		return errors.New("constraint has already expired yet")
	}
	if c.Type == 1 {
		statusConstruction(c.CityId, false)
	}
	if c.Type == 2 {
		statusService(c.CityId, false)
	}
	duration := c.ExpiryDate.Sub(c.IssueDate)
	time.Sleep(time.Duration(duration.Seconds()))
	if c.Type == 1 {
		statusConstruction(c.CityId, true)
	}
	if c.Type == 2 {
		statusService(c.CityId, true)
	}
	return nil
}
