package main

import (
	"fmt"
)

func main() {
	fruits := []string{"apple", "banana", "cherry"}

	// Access elements using for loop
	for index, fruit := range fruits {
		fmt.Println(index, fruit)
	}
}
