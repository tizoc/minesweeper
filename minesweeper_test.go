package main

import (
	"reflect"
	"testing"
)

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
