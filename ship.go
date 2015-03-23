package main

const (
	DIR_UP         = 0
	DIR_DOWN       = 1
	DIR_LEFT       = 2
	DIR_RIGHT      = 3
	LEN_CARRIER    = 5
	LEN_BATTLESHIP = 4
	LEN_SUBMARINE  = 3
	LEN_DESTROYER  = 3
	LEN_CRUISER    = 3
	LEN_PATROL     = 2
)

var SHIP_CLASSES = []struct {
	Name   string
	Length int
}{
	{"Aircraft Carrier", LEN_CARRIER},
	{"Battleship", LEN_BATTLESHIP},
	{"Submarine", LEN_SUBMARINE},
	{"Destroyer", LEN_DESTROYER},
	{"Cruiser", LEN_CRUISER},
	{"Patrol Boat", LEN_PATROL},
}

type Ship struct {
	x, y, direction, length int
	name, class             string
}

func NewShip(name, class string, x, y, direction, length int) *Ship {
	return &Ship{x, y, direction, length, name, class}
}

func (s *Ship) calcShip() [][]int {
	//return a slice of coords
	coords := make([][]int, s.length)

	for i := 0; i < s.length; i++ {
		if i == 0 {
			coords[i] = []int{s.x, s.y}
			continue
		}

		// determine direction
		switch s.direction {
		case DIR_UP:
			coords[i] = []int{s.x, coords[i-1][1] - 1}
		case DIR_DOWN:
			coords[i] = []int{s.x, coords[i-1][1] + 1}
		case DIR_LEFT:
			coords[i] = []int{coords[i-1][0] - 1, s.y}
		case DIR_RIGHT:
			coords[i] = []int{coords[i-1][0] + 1, s.y}
		}
	}

	return coords
}

func (s *Ship) IsHere(x, y int) bool {
	coords := s.calcShip()

	for _, v := range coords {
		if v[0] == x && v[1] == y {
			return true
		}
	}

	return false
}
