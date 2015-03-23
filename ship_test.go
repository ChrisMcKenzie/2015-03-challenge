package main

import "testing"

var ship = &Ship{6, 1, DIR_LEFT, LEN_CARRIER, "USS RONALD REAGAN", "asdasdd"}

func TestCalcLocation(T *testing.T) {
	shipCoords := ship.calcShip()
	if len(shipCoords) != LEN_CARRIER {
		T.Errorf("Expected Result: %d\nActual Result: %d", LEN_CARRIER, len(shipCoords))
	}
	T.Logf("%#v", shipCoords)
}

func TestIsHere(T *testing.T) {
	hit := ship.IsHere(5, 1)
	if hit == false {
		T.Error("Supposed to be hit")
	}

	nothit := ship.IsHere(5, 2)
	if nothit == true {
		T.Error("Supposed to be not hit")
	}
}
