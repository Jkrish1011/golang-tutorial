package advancedfunctions

import "fmt"

/*

We pass functions as input values to other functions when handling scenarios like:
1. HTTP/API handlers
2. Onclick Callbacks

Any time we need to run some code at runtime according to a logic, this might be handy.
We call them as First-Class functions
A function is a first-class function that can be treated like any other value.
A functions type is dependent on the types of it parameters and return values.

Example:
--------
func() int
func(string) int


Currying
--------
Function currying is the process of writing a function that takes a function as input and returns a new function.

*/

func add(x, y int) int {
	return x + y
}

func mul(x, y int) int {
	return x * y
}

// Here the arithmetic is an input variable which is of type function with signature (int, int) and returns an int
func aggregate(a, b, c int, arithmetic func(int, int) int) int {
	return arithmetic(arithmetic(a, b), c)
}

// Function Currying

func functionCurrying(x int) {
	// When passing mul into the selfMath, it creates a new function and embeds the passed one into it
	// and calls the embedded function with the same input value.
	// This makes mul behave like a square function and add like a doubling function.
	squareFunc := selfMath(mul)
	doubleFunc := selfMath(add)

	fmt.Printf("Square of %d : is %d\n", x, squareFunc(x))
	fmt.Printf("Double of %d : is %d", x, doubleFunc(x))
}

// here selfMath accepts a function with signature (int, int) and returns an int as input and
// returns a function with signature (int) and returns an int
func selfMath(mathFunc func(int, int) int) func(int) int {
	return func(x int) int {
		return mathFunc(x, x)
	}
}
