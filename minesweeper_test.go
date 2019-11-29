package main

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

// Utility function for debugging purposes
func showBoard(game *game) string {
	var output strings.Builder
	for row := 0; row < int(game.height); row++ {
		for col := 0; col < int(game.width); col++ {
			cell := row*int(game.height) + col
			value := game.boardView[cell]
			switch value {
			case bomb:
				output.WriteString(" * ")
			case covered:
				output.WriteString(" - ")
			case flagged:
				output.WriteString(" F ")
			case 0:
				output.WriteString(" . ")
			default:
				output.WriteString(fmt.Sprintf(" %d ", value))
			}
		}
		output.WriteString("\n")
	}
	return output.String()
}

func Test_cellNeighbors(t *testing.T) {
	type args struct {
		game *game
		cell uint
	}
	tests := []struct {
		name string
		args args
		want []uint
	}{
		{
			name: "cellNeighbors at corner cell 1",
			args: args{
				game: makeGame(5, 6, 5),
				cell: 0,
			},
			want: []uint{1, 5, 6},
		},
		{
			name: "cellNeighbors at corner cell 2",
			args: args{
				game: makeGame(5, 6, 5),
				cell: 29,
			},
			want: []uint{23, 24, 28},
		},
		{
			name: "cellNeighbors at corner cell in the middle",
			args: args{
				game: makeGame(5, 6, 5),
				cell: 12,
			},
			want: []uint{
				6, 7, 8,
				11 /**/, 13,
				16, 17, 18,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cellNeighbors(tt.args.game, tt.args.cell); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("cellNeighbors(game, %d) = %v, want %v", tt.args.cell, got, tt.want)
			}
		})
	}
}

func Test_uncover(t *testing.T) {
	emptyboard := boardCells{
		0, 0, 0, 0,
		0, 0, 0, 0,
		0, 0, 0, 0,
		0, 0, 0, 0,
	}

	t.Run("uncover() cascading on 0-value cell", func(t *testing.T) {
		game := makeGame(4, 4, 1)
		copy(game.board, emptyboard)
		placeBombAt(game, 6)
		uncover(game, 0)
		c := int8(covered)
		expected := boardCells{
			0, 1, c, c,
			0, 1, c, c,
			0, 1, 1, 1,
			0, 0, 0, 0,
		}
		if !reflect.DeepEqual(expected, game.boardView) {
			gameTmp := makeGame(4, 4, 1)
			gameTmp.boardView = expected
			t.Errorf("uncover() did not cascade as expected, wanted:\n%s\n got:\n%s\n", showBoard(gameTmp), showBoard(game))
		}
	})

	t.Run("uncover() not cascading on non 0-value cell", func(t *testing.T) {
		game := makeGame(4, 4, 1)
		copy(game.board, emptyboard)
		placeBombAt(game, 6)
		uncover(game, 2)
		c := int8(covered)
		expected := boardCells{
			c, c, 1, c,
			c, c, c, c,
			c, c, c, c,
			c, c, c, c,
		}
		if !reflect.DeepEqual(expected, game.boardView) {
			gameTmp := makeGame(4, 4, 1)
			gameTmp.boardView = expected
			t.Errorf("uncover() did cascade when it should had not, wanted:\n%s\n got:\n%s\n", showBoard(gameTmp), showBoard(game))
		}
	})

}
