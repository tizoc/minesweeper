package main

import (
	"math/rand"
	"time"
)

type boardCells []int8

// Cell values
// N from 0 to 8 represents the amount of nearby mines
const (
	mine    = -1
	flagged = -126
	covered = -127
)

const (
	notStartedStatus = "NotStarted"
	startedStatus    = "Started"
	lostStatus       = "Lost"
	wonStatus        = "Won"
)

type game struct {
	board     boardCells // Underlying board with mines and numbers placed
	BoardView boardCells `json:"board"` // Board as shown to the user (includes flags and covered cells)
	Height    uint       `json:"height"`
	Width     uint       `json:"width"`
	Mines     uint       `json:"mines"`
	StartedAt time.Time  `json:"startedAt"`
	EndedAt   time.Time  `json:"endedAt"`
	Status    string     `json:"status"`
}

func makeGame(width, height, mines uint) *game {
	game := &game{
		Width:     width,
		Height:    height,
		board:     make([]int8, height*width),
		BoardView: make([]int8, height*width),
		Mines:     mines,
		Status:    notStartedStatus,
	}
	placeMines(game)
	coverBoard(game)
	return game
}

func startGame(game *game) {
	game.StartedAt = time.Now().UTC()
	game.Status = startedStatus
}

func coverBoard(game *game) {
	for i := range game.BoardView {
		game.BoardView[i] = covered
	}
}

func placeMines(game *game) {
	for i := uint(0); i < game.Mines; i++ {
		placeMineAtRandom(game)
	}
}

// Finds an empty cell in the game board
// and places a mine on it
func placeMineAtRandom(game *game) {
	cellsCount := int(game.Width * game.Height)
	for {
		cell := uint(rand.Intn(cellsCount))
		if placeMineAt(game, cell) {
			return
		}
	}
}

// Places a mine at cell.
// Returns true if succesful, otherwise false
func placeMineAt(game *game, cell uint) bool {
	if game.board[cell] == mine {
		return false
	}
	game.board[cell] = mine
	incrementNeighbors(game, cell)
	return true

}

// Given a cell N, increments the value of
// all nearby cells by 1
func incrementNeighbors(game *game, cell uint) {
	for _, cell := range cellNeighbors(game, cell) {
		if game.board[cell] != mine {
			game.board[cell]++
		}
	}
}

// Represents an offset from a given cell, positive and negative
type offset struct {
	rows int
	cols int
}

// Offsets for neighbors of a given cell (cell itself would be {0, 0})
var neighborOffsets = []offset{
	{-1, -1}, {-1, +0}, {-1, +1},
	{+0, -1} /* ... */, {+0, +1},
	{+1, -1}, {+1, +0}, {+1, +1},
}

// Given a cell position N, returns a slice containing
// the locations of all neighboring cells.
// Board edges are properly handled so that only valid
// cell locations are returned, and nothing outside the
// board is.
func cellNeighbors(game *game, cell uint) []uint {
	neighbors := make([]uint, 0, 8) // 8 is max valid amount of neighbors
	col := int(cell % game.Width)
	row := int(cell / game.Width)

	for _, offset := range neighborOffsets {
		newRow := row + offset.rows
		newCol := col + offset.cols
		if 0 <= newCol && uint(newCol) < game.Width && 0 <= newRow && uint(newRow) < game.Height {
			neighbors = append(neighbors, uint(newCol)+uint(newRow)*game.Width)
		}
	}

	return neighbors
}

// Toggle cell flag.
// Return true if sucessful, false otherwise.
// Fails when the cell is not covered.
func toggleFlag(game *game, cell uint) bool {
	cells := game.BoardView
	switch cells[cell] {
	case flagged:
		cells[cell] = covered
		return true
	case covered:
		cells[cell] = flagged
		return true
	default:
		return false
	}
}

// Uncovers a cell.
// If the cell has no nearby mines, uncovers
// al nearby cells too in cascade.
// Return true if sucessful, false otherwise.
// Fails when the cell is not covered or there
// is a flag on it.
func uncover(game *game, cell uint) bool {
	if game.BoardView[cell] != covered {
		return false
	}
	cellValue := game.board[cell]
	game.BoardView[cell] = cellValue
	if cellValue == 0 { // No mines near, cascade
		for _, otherCell := range cellNeighbors(game, cell) {
			uncover(game, otherCell)
		}
	}

	return true
}

// Checks if the game is finished already
// either by a mine being uncovered
// or the game being won.
func checkGameFinished(game *game) bool {
	// To win, player must uncover all cells
	// that do not contain mines
	// If a mine is uncovered, player loses
	pendingCells := game.Width*game.Height - game.Mines
	for _, cellValue := range game.BoardView {
		if cellValue == mine {
			game.Status = lostStatus
			game.EndedAt = time.Now().UTC()
			return true // Mine uncovered, game finished
		}
		pendingCells--
	}
	if pendingCells == 0 {
		game.Status = wonStatus
		game.EndedAt = time.Now().UTC()
		return true // No pending cells to uncover, game finished
	}
	return false
}
