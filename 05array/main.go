package main

import "fmt"

func main() {
	var array = []string{"python", "java", "cpp"}
	fmt.Println(len(array))
	fmt.Printf("%T\n", array)

	array = append(array, "golang")
	fmt.Println(array)

	array = append(array[1:])
	fmt.Println(array)
}
