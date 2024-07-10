package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to Pizza Hut")
	fmt.Println("Please rate our pizza")

	input := bufio.NewReader(os.Stdin)

	rating, _ := input.ReadString('\n')

	fmt.Println("Thanks for the rating", rating)

	numrating, err := strconv.ParseFloat(strings.TrimSpace(rating), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Your rating is", numrating+1)
	}

}
