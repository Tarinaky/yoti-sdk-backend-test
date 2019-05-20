package main

import (
		"fmt"
		"log"
		"net/http"
		"database/sql"
		_ "github.com/lib/pq"
		//"io/ioutil"
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

func (env* Env) PostRoomba (response http.ResponseWriter, request *http.Request) {/*
		// Parse input
		body,err := ioutil.ReadAll(request.Body)
		if err != nil {
				log.Panic(err)
		}

		var room Room
		if err := json.Unmarshal(body, &room); err != nil {
				log.Panic(err)
		}

		// Do roomba calculation
		for _,instruction := range room.Instructions {
			if err := room.Roomba.Move(instruction, room); err != nil {
				log.Panic(err)
			}
		}

		// Update db
		transaction, err := env.db.Begin()
		if err != nil {
			log.Panic(err)
		}
		if _,err := env.db.Exec(
			`INSERT INTO roomba.rooms (
				room_width,
				room_height
			) VALUES (
				$1,
				$2
			)`,
			room.Width,
			room.Height,
		); err != nil {
			log.Panic(err)
		}*/



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

