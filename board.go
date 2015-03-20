package main

import "math/rand"

const (
	BOARD_X         = 16
	BOARD_Y         = 16
	SHIPS_TO_DEPLOY = 5
)

type ShipMeta struct {
	Name   string
	Length int
}

var shipNames = []ShipMeta{
	ShipMeta{"Aircraft Carrier", 5},
	ShipMeta{"Battleship", 4},
	ShipMeta{"Submarine", 3},
	ShipMeta{"Destroyer", 3},
	ShipMeta{"Cruiser", 3},
	ShipMeta{"Patrol Boat", 2},
}

type Board struct {
	Ships []Ship
}

func NewBoard() (board *Board) {
	i := 0
	for i < SHIPS_TO_DEPLOY {
		sp := rand.Intn(len(shipNames) - 1)
		shipmeta := shipNames[sp]
		direction := rand.Intn(DIR_RIGHT)

		// board.Ships = append(board.Ships, )
		i++
	}
	return
}
