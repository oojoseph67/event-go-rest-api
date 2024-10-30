package models

import "time"

type Event struct {
	ID          int
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int
}

var events []Event = []Event{}

func (e *Event) CreateEvent() { // method of struct Event to create an event
	// later add it to the database
	events = append(events, *e)
}

// this wont be a method because we dont call it on an instance of Event, it will be a function
func GetAllEvents() []Event {
	return events
}
