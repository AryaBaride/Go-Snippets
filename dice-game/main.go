package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Welcome to the dice game")

	dicenumber := rand.Intn(6) + 1
	fmt.Println("The value of the dice is ", dicenumber)

	switch dicenumber {
	case 1:
		fmt.Println("You can start the game!")
	case 2:
		fmt.Println("You can move two steps")
	case 3:
		fmt.Println("You can move three steps")
	case 4:
		fmt.Println("You can move four steps")
	case 5:
		fmt.Println("You can move five steps")
	case 6:
		fmt.Println("You can move six steps")
	default:
		fmt.Println("You can't move")
	}

}
