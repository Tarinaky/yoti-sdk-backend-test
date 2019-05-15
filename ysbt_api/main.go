package main

import (
		"fmt"
		"log"
		//"net/http"
		"database/sql"
		_ "github.com/lib/pq"
		//"github.com/gorilla/mux"
)

const (
		port = "5432"
		server = "192.168.99.100"
		password = "password"
		username = "ysbt_db"
		database = "ysbt_db"
		sslmode = "disable"
)

type Env struct {
		db *sql.DB
}

func (env* Env) PostRoomba (response http.ResponseWriter, request *http.Request) {
		// Parse input
		// Do roomba calculation
		// Update db
		// Write response
}

func main() {
		psqlInfo := fmt.Sprintf(
				"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
				server,
				port,
				username,
				password,
				database,
				sslmode,
		)

		db,err := sql.Open("postgres", psqlInfo)
		if err != nil {
				log.Panic(err)
		}

		if err := db.Ping(); err != nil {
				log.Panic(err)
		}

		log.Print("Database connection open")
}

