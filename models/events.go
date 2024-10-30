package models

import "time"

type Event struct {
	ID          int
	Name        string
	Description string
	Location    string
	DateTime    time.Time
	UserID      int
}

var events []Event = []Event{}

func (e *Event) Save() { // method to save the event
	// later add it to the database
	events = append(events, *e)
}

// this wont be a method because we dont call it on an instance of Event, it will be a function
func GetAllEvents() []Event {
	return events
}
