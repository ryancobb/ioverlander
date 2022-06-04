package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
  "gorm.io/gorm/clause"
)

type OverlanderPoint struct {
	gorm.Model
	Id          int
	Name        string
	Description string
	Category    string
	Latitude    string
	Longitude   string
}

func main() {
	dsn := "host=localhost user=postgres dbname=ioverlander port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	db.AutoMigrate(&OverlanderPoint{})

  points := FetchPoints()

  var processed_points []OverlanderPoint
  for _, point := range points {
    processed_points = append(processed_points, point.toDb())
  }

  db.Clauses(clause.OnConflict{
    UpdateAll: true,
  }).Create(&processed_points)
}
