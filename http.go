package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func gamesCollectionHandler(w http.ResponseWriter, r *http.Request) {
	// TODO
}

func gameHandler(w http.ResponseWriter, r *http.Request) {
	// TODO
}

func gamePlayHandler(w http.ResponseWriter, r *http.Request) {
	// TODO
}

func serveMineSweeper() {
	router := mux.NewRouter()
	router.HandleFunc("/minesweeper/games", gamesCollectionHandler).Methods("POST")
	router.HandleFunc("/minesweeper/games/{id}", gameHandler).Methods("GET")
	router.HandleFunc("/minesweeper/games/{id}/plays", gamePlayHandler).Methods("POST")
	log.Print("Starting server on port 8080")
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", router))
}
