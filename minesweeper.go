package main

import (
	"math/rand"
)

type boardCells []int8

// Cell values
// N from 0 to 8 represents the amount of nearby mines
const (
	bomb    = -1
	flagged = -126
	covered = -127
)

type game struct {
	board     boardCells // Underlying board with bombs and numbers placed
	boardView boardCells // Board as shown to the user (includes flags and covered cells)
	height    uint
	width     uint
	bombs     uint
}

func makeGame(width, height, bombs uint) *game {
	game := &game{
		width:     width,
		height:    height,
		board:     make([]int8, height*width),
		boardView: make([]int8, height*width),
		bombs:     bombs,
	}
	placeBombs(game)
	coverBoard(game)
	return game
}

func coverBoard(game *game) {
	for i := range game.boardView {
		game.boardView[i] = covered
	}
}

func placeBombs(game *game) {
	for i := uint(0); i < game.bombs; i++ {
		placeBombAtRandom(game)
	}
}

// Finds an empty cell in the game board
// and places a bomb on it
func placeBombAtRandom(game *game) {
	cellsCount := int(game.width * game.height)
	for {
		cell := uint(rand.Intn(cellsCount))
		if placeBombAt(game, cell) {
			return
		}
	}
}

// Places a bomb at cell.
// Returns true if succesful, otherwise false
func placeBombAt(game *game, cell uint) bool {
	if game.board[cell] == bomb {
		return false
	}
	game.board[cell] = bomb
	incrementNeighbors(game, cell)
	return true

}

// Given a cell N, increments the value of
// all nearby cells by 1
func incrementNeighbors(game *game, cell uint) {
	for _, cell := range cellNeighbors(game, cell) {
		if game.board[cell] != bomb {
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
	col := int(cell % game.width)
	row := int(cell / game.width)

	for _, offset := range neighborOffsets {
		newRow := row + offset.rows
		newCol := col + offset.cols
		if 0 <= newCol && uint(newCol) < game.width && 0 <= newRow && uint(newRow) < game.height {
			neighbors = append(neighbors, uint(newCol)+uint(newRow)*game.width)
		}
	}

	return neighbors
}

// Toggle cell flag.
// Return true if sucessful, false otherwise.
// Fails when the cell is not covered.
func toggleFlag(game *game, cell uint) bool {
	cells := game.boardView
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
// If the cell has no nearby bombs, uncovers
// al nearby cells too in cascade.
// Return true if sucessful, false otherwise.
// Fails when the cell is not covered or there
// is a flag on it.
func uncover(game *game, cell uint) bool {
	if game.boardView[cell] != covered {
		return false
	}
	cellValue := game.board[cell]
	game.boardView[cell] = cellValue
	if cellValue == 0 { // No bombs near, cascade
		for _, otherCell := range cellNeighbors(game, cell) {
			uncover(game, otherCell)
		}
	}

	return true
}

// Checks if the game is finished already
// either by a bomb being uncovered
// or the game being won.
func isGameFinished(game *game) bool {
	// To win, player must uncover all cells
	// that do not contain bombs
	// If a bomb is uncovered, player loses
	pendingCells := game.width*game.height - game.bombs
	for _, cellValue := range game.boardView {
		if cellValue == bomb {
			return true // Bomb uncovered, game finished
		}
		pendingCells--
	}
	if pendingCells == 0 {
		return true // No pending cells to uncover, game finished
	}
	return false
}
