package main

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
	// TODO
	// - find random cell without a bomb
	// - place bomb at cell
	// - increment nearby-mines count for all neighbor cells without bombs
}

func incrementNeighbors(game *game, cell uint) {
	// TODO
	// - Find neighbors without bombs
	// - Increase value by 1 on each
}
