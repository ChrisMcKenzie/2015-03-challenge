package main

const (
	BOARD_X         = 16
	BOARD_Y         = 16
	NUM_ROUNDS      = 6
	SHOTS_PER_ROUND = 5
)

type ship interface {
	CheckHit(x, y int) bool
}

type Board [BOARD_X][BOARD_Y]int

func (b *Board) AddShip(x, y int) {

}

var (
	board Board
)
