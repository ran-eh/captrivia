package main

import (
	"database/sql"
	"encoding/json"
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
	Context   interface{}
	Error     error
}

// TODO: instead of using a global, use a closure as shown in
// https://stackoverflow.com/questions/34046194/how-to-pass-arguments-to-router-handlers-in-golang-using-gin-web-framework
var Db *sql.DB

// Change to get events logged in the past for testing
var ShiftDateDays int = 0

func EventServiceConnect() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	log.Printf("Successfully connectes to pg server: %s:%d.%s", host, port, dbname)

	return db
}

func EventServiceClose() {
	Db.Close()
}

func EventServicePost(ev *Event) error {
	ev.EventId = uuid.New().String()
	ev.Timestamp = time.Now().AddDate(0, 0, ShiftDateDays)
	ev.Program = "backend"
	contextJson, err := json.Marshal(ev.Context)
	if err != nil {
		return err
	}

	sqlFmt := `
	INSERT INTO events (timestamp, event_id, session_id, program, type, data, context)
	VALUES ($1, '%s', $2, $3, $4, $5, '%s')
	`
	sql := fmt.Sprintf(sqlFmt, ev.EventId, string(contextJson))

	_, err = Db.Exec(sql, ev.Timestamp, ev.SessionID, ev.Program, ev.Type, ev.Data)

	log.Println(ev)
	return err
}
