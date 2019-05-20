package main

import (
		"fmt"
		"log"
		"net/http"
		"database/sql"
		_ "github.com/lib/pq"
)

const (
		port = "5432"
		server = "db"
		password = "password"
		username = "ysbt_db"
		database = "ysbt_db"
		sslmode = "disable"
)

type Env struct {
		db *sql.DB
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

