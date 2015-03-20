package main

const (
	BOARD_X         = 16
	BOARD_Y         = 16
	NUM_ROUNDS      = 6
	SHOTS_PER_ROUND = 5
)

type (
	Board struct {
		Ships []Ship
	}
)

func (b *Board) AddShip(ship Ship) {

}

var (
	board Board
)
