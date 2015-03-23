package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	BOARD_X         = 16
	BOARD_Y         = 16
	SHIPS_TO_DEPLOY = 5
)

type Board struct {
	Ships []*Ship
}

func (b *Board) isConflict(x, y int) bool {
	// TODO: check conflict with whole ship
	// not just start
	for _, v := range b.Ships {
		if v.IsHere(x, y) {
			return true
		}
	}

	return false
}

func NewBoard() (board *Board) {
	rand.Seed(time.Now().Unix())
	board = &Board{}
	// var ships = shipNames
	i := 0
	// Generate randow board with ships
	for i < SHIPS_TO_DEPLOY {
		classID := rand.Intn(len(SHIP_CLASSES) - 1)
		class := SHIP_CLASSES[classID]
		direction := rand.Intn(DIR_RIGHT)
		// calc x for direction

		// TODO: break out into func and
		// check for conflict
		var x, y int
		switch direction {
		case DIR_UP:
			x = rand.Intn(BOARD_X)
			y = rand.Intn(BOARD_Y-class.Length) + class.Length
		case DIR_DOWN:
			x = rand.Intn(BOARD_X)
			y = rand.Intn(BOARD_Y - class.Length)
		case DIR_LEFT:
			x = rand.Intn(BOARD_X-class.Length) + class.Length
			y = rand.Intn(BOARD_Y)
		case DIR_RIGHT:
			x = rand.Intn(BOARD_X - class.Length)
			y = rand.Intn(BOARD_Y)
		}

		ship := NewShip(class.Name, class.Name, x, y, direction, class.Length)
		fmt.Println(ship)
		board.Ships = append(board.Ships, ship)
		i++
	}
	return
}
