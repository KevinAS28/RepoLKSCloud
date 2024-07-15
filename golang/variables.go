package main

import (
	"fmt"
)

func main() {
	// Create variables
	name := "Alice"
	age := 30
	pi := 3.14159

	// Convert to different data types
	age_str := fmt.Sprint(age) // Convert int to string
	pi_int := int(pi)          // Convert float to int (truncates)

	// Print the results
	fmt.Printf("Name: %s (string)\n", name)
	fmt.Printf("Age: %d (int), %s (string)\n", age, age_str)
	fmt.Printf("Pi: %f (float), %d (int)\n", pi, pi_int)
}
