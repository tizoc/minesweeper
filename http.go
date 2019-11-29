package main

import (
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
)

var games = make(map[string]*game)

type gameSettings struct {
	Width  uint `json:"width"`
	Height uint `json:"height"`
	Mines  uint `json:"mines"`
}

func gamesCollectionHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: validate content-type of request
	// 9x9 with 10 mines is the default
	config := gameSettings{Width: 9, Height: 9, Mines: 10}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&config)

	if err != nil {
		// TODO: add error message
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// TODO: limit board size
	// check that the amount of mines makes sense
	// for the board size and is > 0
	game := makeGame(config.Width, config.Height, config.Mines)
	gameID := makeGameID()
	games[gameID] = game

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	encoder := json.NewEncoder(w)
	encoder.Encode(map[string]string{
		"ID": gameID,
	})
}

// Makes a random identifier string for games
// Doesn't check for collisions
func makeGameID() string {
	hasher := sha1.New()
	randomBytes := make([]byte, 4)
	rand.Read(randomBytes)
	hash := hasher.Sum(randomBytes)
	return base64.URLEncoding.EncodeToString(hash)
}

func gameHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	gameID := vars["id"]
	game, ok := games[gameID]

	if !ok {
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		encoder := json.NewEncoder(w)
		encoder.Encode(game)
	}
}

type play struct {
	Action   string `json:"action"`
	Location uint   `json:"location"`
}

func gamePlayHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: validate content-type of request
	// TODO: verify that game has not been finished already
	vars := mux.Vars(r)
	gameID := vars["id"]
	game, ok := games[gameID]

	if !ok {
		w.WriteHeader(http.StatusNotFound)
	} else {
		play := play{}
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&play)

		if err != nil {
			// TODO: add error message
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if play.Location >= uint(len(game.BoardView)) {
			// TODO: add error message
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		switch play.Action {
		case "flag":
			// TODO: notify if flagging not valid or ignore?
			toggleFlag(game, play.Location)
		case "uncover":
			// TODO: notify if already uncovered or ignore?
			uncover(game, play.Location)
		default:
			// TODO: add error message
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		encoder := json.NewEncoder(w)
		encoder.Encode(game)
	}
}

func serveMineSweeper() {
	router := mux.NewRouter()
	router.HandleFunc("/minesweeper/games", gamesCollectionHandler).Methods("POST")
	router.HandleFunc("/minesweeper/games/{id}", gameHandler).Methods("GET")
	router.HandleFunc("/minesweeper/games/{id}/plays", gamePlayHandler).Methods("POST")
	log.Print("Starting server on port 8080")
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", router))
}
