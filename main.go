package main

import (
	"fmt"
)

func add(a, b int) int {
	return a + b
}

func main() {
	result := add(2, 3)
	fmt.Println("The sum is:", result)
}
