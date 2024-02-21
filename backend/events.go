package main

import (
	"time"
)

type Event struct {
	EventId   string
	Timestamp time.Time
	SessionID string
	Program   string
	Type      string
	Data      interface{}
	Context   interface{}
	Error     error
}
