package main

import (
	"database/sql"
	"encoding/json"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "captrivia"
)

// TODO: instead of using a global, use a closure as shown in
// https://stackoverflow.com/questions/34046194/how-to-pass-arguments-to-router-handlers-in-golang-using-gin-web-framework
var Db *sql.DB

func EventServiceConnect() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	return db
}

func EventServiceClose() {
	Db.Close()
}

func EventServicePost(program string, _type string, data interface{}, context interface{}) error {
	dataJson, err := json.Marshal(data)
	if err != nil{
		return err
	}
	contextJson, err := json.Marshal(context)
	if err != nil{
		return err
	}
	sql := `
	INSERT INTO events (program, type, data, context)
	VALUES ($1, $2, $3, $4)
	`
	_, err = Db.Exec(sql, program, _type, dataJson, contextJson)
	
	fmt.Println(program, _type, data, context)
	return err
}
