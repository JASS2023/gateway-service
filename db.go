package main

import (
	"time"

	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Constraint struct {
	// TODO is the ID really a integer or an UUID?
	// TODO which geometry should we support for the coordinates?
	ID          *uint `gorm:"primaryKey"`
	Type        uint
	Quadrant    string `gorm:"default: 1111"`
	X           *uint  `gorm:"not null"`
	Y           *uint  `gorm:"not null"`
	MaxSpeed    float64
	Days        string    `gorm:"default:'1111111'"`
	StartTime   string    `gorm:"default:'00:00:00'"`
	EndTime     string    `gorm:"default:'23:59:59'"`
	IssueDate   time.Time `gorm:"default:current_timestamp"`
	ExpiryDate  time.Time
	Description string `gorm:"default:'N/A'"`
}

var DB *gorm.DB

func ConnectDB() error {
	dsn := "host=localhost user=jass2023 password=jass2023 dbname=jass2023 port=5432 sslmode=disable"
	gormDb, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.WithError(err).Fatal("Could not connect to the database")
	}

	log.Info("Connected to the database")

	// Migrate the schema
	gormDb.AutoMigrate(&Constraint{})

	DB = gormDb

	return nil
}
