package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "captrivia"
)

type Event struct {
	EventId   string
	Timestamp time.Time
	SessionID string
	Program   string
	Type      string
	Data      interface{}
	Error     error
}

type EventSender struct {
	Db *sql.DB
	// TODO only include in debug build
	// Allows posting events for days that are not today.
	ShiftDateDays int
}

func NewEventSender() *EventSender {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	log.Printf("Successfully connectes to pg server: %s:%d.%s", host, port, dbname)
	return &EventSender{Db: db}
}

func (es *EventSender) Close() {
	es.Db.Close()
}

func (es *EventSender) Send(ev *Event) error {
	ev.EventId = uuid.New().String()
	ev.Timestamp = time.Now().AddDate(0, 0, es.ShiftDateDays)
	ev.Program = "backend"
	sqlFmt := `
	INSERT INTO events (timestamp, event_id, session_id, program, type, data)
	VALUES ($1, '%s', $2, $3, $4, $5)
	`
	sql := fmt.Sprintf(sqlFmt, ev.EventId)

	_, err := es.Db.Exec(sql, ev.Timestamp, ev.SessionID, ev.Program, ev.Type, ev.Data)

	log.Println(ev)
	return err
}
