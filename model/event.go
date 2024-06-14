package model

import "time"

type Event struct {
	Id          int
	Name        string
	Description string
	Timestamp   time.Time
}
