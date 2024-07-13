package main

import "fmt"

func main() {
	fmt.Println("Welcome to pointers")

	var pointer *int

	number := 5

	pointer = &number

	fmt.Println("The value of the address of pointer is", pointer)
	fmt.Println("Actual value of pointer", *pointer)

}
