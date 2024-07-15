package main

import (
	"fmt"
)

func sumNumbers(intArg int, floatArg float64) float64 {
	hasil := floatArg + float64(intArg)
	return hasil
}

func main() {
	total := sumNumbers(5, 3.14)
	fmt.Println("The sum is: ", total)
}
