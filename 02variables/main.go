package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to the Golang"
	fmt.Println(welcome)

	name := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter your name")

	input, _ := name.ReadString('\n')
	fmt.Println("Welcome,", input)

}
