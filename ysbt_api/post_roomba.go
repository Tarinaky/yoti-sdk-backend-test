package main

import (
		"io/ioutil"
		"encoding/json"
		"net/http"
		"log"
		"yoti-sdk-backend-test/ysbt_api/roomba"
	)


type output struct {
	Coords []uint `json:coords`
	Patches uint `json:patches`
}

func (env* Env) PostRoomba (response http.ResponseWriter, request *http.Request) {
		if request.Method != "POST" {
			response.Header().Set("Content-Type", "text/plain")
			response.WriteHeader(http.StatusBadRequest)
			response.Write([]byte("400 - This API serves POST only"))
			return
		}

		// Parse input
		body,err := ioutil.ReadAll(request.Body)
		if err != nil {
				log.Panic(err)
		}

		var room roomba.Room
		if err := json.Unmarshal(body, &room); err != nil {
			response.Header().Set("Content-Type", "text/plain")
			response.WriteHeader(http.StatusBadRequest)
			response.Write([]byte("400 - Could not parse JSON"))
			log.Print(err)
			return
		}

		// Do roomba calculation
		if err := room.Process(); err != nil {
			response.Header().Set("Content-Type", "text/plain")
			response.WriteHeader(http.StatusInternalServerError)
			response.Write([]byte("500 - Internal Server Error"))
			log.Print(err)
			return
		}

		// Update db
		if err := room.Store(env.db); err != nil {
			response.Header().Set("Content-Type", "text/plain")
			response.WriteHeader(http.StatusInternalServerError)
			response.Write([]byte("500 - Internal Server Error"))
			log.Print(err)
			return
		}

		// Write response
		var out output
		out.Coords = []uint{room.Roomba.CurrentX, room.Roomba.CurrentY}
		out.Patches = room.Roomba.DirtCollected

		response.Header().Set("Content-Type", "application/json")
		response.WriteHeader(http.StatusCreated)
		json.NewEncoder(response).Encode(out)
}
