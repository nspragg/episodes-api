package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nspragg/episodes-api/db"
	"github.com/nspragg/episodes-api/handlers"
)

const port = ":8181"

func main() {
	client, err := db.Dial()
	if err != nil {
		log.Fatal("Failed to connect to DB: " + err.Error())
	}

	defer client.Close()

	r := mux.NewRouter()
	r.HandleFunc("/episodes/{id}", handlers.GetEpisodes(r, client)).Methods("GET")

	log.Fatal(http.ListenAndServe(port, r))
}
