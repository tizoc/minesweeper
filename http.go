package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var games = make(map[string]*game)

type gameSettings struct {
	Width  uint `json:"width"`
	Height uint `json:"height"`
	Bombs  uint `json:"bombs"`
}

func gamesCollectionHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: validate content-type of request
	// 9x9 with 10 bombs is the default
	config := gameSettings{Width: 9, Height: 9, Bombs: 10}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&config)

	if err != nil {
		// TODO: add error message
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// TODO: limit board size
	// check that the amount of bombs makes sense
	// for the board size and is > 0
	game := makeGame(config.Width, config.Height, config.Bombs)
	gameID := makeGameID()
	games[gameID] = game

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	encoder := json.NewEncoder(w)
	// TODO: all fields in game are private, either make
	// the exportable ones public or define a new struct
	encoder.Encode(game)
}

func makeGameID() string {
	return "TODO"
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
