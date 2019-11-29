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
	board  boardCells
	height uint
	width  uint
	bombs  uint
}

func makeGame(width, height, bombs uint) *game {
	game := &game{
		width:  width,
		height: height,
		board:  make([]int8, height*width),
		bombs:  bombs,
	}
	placeBombs(game)
	return game
}

func placeBombs(game *game) {
	for i := uint(0); i < game.bombs; i++ {
		placeBombAtRandom(game)
	}
}

func placeBombAtRandom(game *game) {
	cellsCount := game.width * game.height
	for {
		cell := uint(rand.Intn(int(cellsCount)))
		if game.board[cell] != bomb {
			game.board[cell] = bomb
			incrementNeighbors(game, cell)
			return
		}
	}
}

func incrementNeighbors(game *game, cell uint) {
	for _, cell := range cellNeighbors(game, cell) {
		if game.board[cell] != bomb {
			game.board[cell]++
		}
	}
}

func cellNeighbors(game *game, cell uint) []uint {
	// TODO: find list of empty neighbors and return
	// the list of cells
	return make([]uint, 0)
}
