package main

import (
		"fmt"
		"log"
		"net/http"
		"database/sql"
		_ "github.com/lib/pq"
		"io/ioutil"
		"encoding/json"
		"yoti-sdk-backend-test/ysbt_api/roomba"
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

type output struct {
	Coords []uint `json:coords`
	Patches uint `json:patches`
}

func (env* Env) PostRoomba (response http.ResponseWriter, request *http.Request) {
		// Parse input
		body,err := ioutil.ReadAll(request.Body)
		if err != nil {
				log.Panic(err)
		}

		var room roomba.Room
		if err := json.Unmarshal(body, &room); err != nil {
				log.Panic(err)
		}

		// Do roomba calculation
		if err := room.Process(); err != nil {
			log.Panic(err)
		}

		// Update db
		if err := room.Store(env.db); err != nil {
			log.Panic(err)
		}

		// Write response
		var out output
		out.Coords = []uint{room.Roomba.CurrentX, room.Roomba.CurrentY}
		out.Patches = room.Roomba.DirtCollected

		response.Header().Set("Content-Type", "application/json")
		response.WriteHeader(http.StatusCreated)
		json.NewEncoder(response).Encode(out)
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

		var env Env
		env.db = db

		http.HandleFunc("/roomba", env.PostRoomba)
		log.Fatal(http.ListenAndServe(":8080",nil))
}

