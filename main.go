package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"hwCalendar/model/event"
	"log"
	"time"
)

func main() {
	err := godotenv.Load(".env.example")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	newEvent := event.Event{Id: 1, Name: "Simple", Description: "Very simple", Timestamp: time.Now()}
	_, err = newEvent.Add()
	if err != nil {
		log.Println(err)
	}
	e, err := event.ById(1)
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("Just added event: %+v\n", e)
	all, err := event.All()
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("All events after adding one: %+v\n", all)

	e.Name = "Yet another name"
	_ = e.Update()
	e, err = event.ById(1)
	fmt.Printf("Same event after updating its name: %+v\n", e)

	awesome := event.Event{Id: 2, Name: "Awesome", Description: "Very awesome", Timestamp: time.Now()}
	_, _ = awesome.Add()
	all, _ = event.All()
	fmt.Printf("Added new event: %+v\n", all)

	_ = e.Delete()
	all, _ = event.All()
	fmt.Printf("Deleted previous event: %+v\n", all)
}
