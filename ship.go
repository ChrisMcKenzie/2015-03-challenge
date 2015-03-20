package main

const (
	DIR_UP    = 1
	DIR_DOWN  = 2
	DIR_LEFT  = 3
	DIR_RIGHT = 4
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
