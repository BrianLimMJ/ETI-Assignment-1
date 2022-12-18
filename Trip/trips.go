package Trips

import (
	_ "github.com/go-sql-driver/mysql"
)

type Trips struct {
	tripID      int
	passengerId string
	driverId    string
	startPostal string
	endPostal   string
	isStarted   bool
	isCompleted bool
}
