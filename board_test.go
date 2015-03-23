package main

import "testing"

func TestCreateBoard(T *testing.T) {
	board := NewBoard()

	T.Log(board)
}
