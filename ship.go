package main

const (
	DIR_UP    = 0
	DIR_DOWN  = 1
	DIR_LEFT  = 2
	DIR_RIGHT = 3
)

type Ship struct {
	startx, starty, direction, length int
	name                              string
}

func NewShip(name string, startx, starty, direction, length int) *Ship {
	return &Ship{startx, starty, direction, length, name}
}

func (s *Ship) CheckHit(x, y int) bool {
	return false
}
