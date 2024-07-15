package main

import (
	"fmt"
)

func main() {
	// age := 20
	var age int = 20

	if age >= 18 {
		fmt.Println("You are eligible to vote.")
	} else {
		fmt.Println("You are not eligible to vote.")
	}

	// ================================================

	marks := 93
	// var marks int = 93

	if marks >= 90 {
		fmt.Println("Excellent! You got an A.")
	} else if marks >= 80 {
		fmt.Println("Great job! You got a B.")
	} else {
		fmt.Println("Keep practicing! You got a C or lower.")
	}
}
