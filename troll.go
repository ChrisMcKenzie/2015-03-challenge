package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

func main() {
	fmt.Println("Shall we play a game?")
	fmt.Println("Here's six rounds. You can get more Stella Artois in the kitchen.")

	score := 0

	i := 0

	for i < 30 {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter a spot.")
		text, _ := reader.ReadString('\n')

		if len(text) > 2 {
			fmt.Println("You hit a spot! Enter another.")
			score := 1
		} else {
			fmt.Println("Missed, try again.")
			score := score - 1
		}

		i++
	}

	if score > 52 {
		fmt.Println("you win!")
	} else {
		fmt.Println("you lost!")
	}

	var direction = []string{"horizontal", "vertical"}

	type ship struct {
		name      string
		spots     int
		worth     int
		placement string
	}

	var ships = []ship{
		ship{"Aircraft Carrier", 5, 20, direction[rand.Intn(1)]},
		ship{"Battleship", 4, 12, direction[rand.Intn(1)]},
		ship{"Submarine", 3, 6, direction[rand.Intn(1)]},
		ship{"Destroyer", 3, 6, direction[rand.Intn(1)]},
		ship{"Cruiser", 3, 6, direction[rand.Intn(1)]},
		ship{"Patrol Boat", 2, 2, direction[rand.Intn(1)]},
	}

	type Board struct {
		description string
		Ships       []ship
	}

	board := Board{"I am a 16x16 board.", ships}

	fmt.Println(board)
}
