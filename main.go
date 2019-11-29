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
	return &game{
		width:  width,
		height: height,
		board:  make([]int8, height*width),
		bombs:  bombs,
	}
}

func main() {
}
