package slices

/*
Many functions, especially those in the standard library, can take an arbitrary number of final arguments.
This is accomplished by using the "..." syntax in the function signature.

A variadic function receives the variadic arguments as a slice.

func concat(strs ...string) string {
	// strs variable is just like a slice. We use it as if it's a slice
    final := ""
    // strs is just a slice of strings
    for i := 0; i < len(strs); i++ {
        final += strs[i]
    }
    return final
}

func main() {
    final := concat("Hello ", "there ", "friend!")
    fmt.Println(final)
    // Output: Hello there friend!
}

Spread Operator
---------------
The spread operator allows us to pass a slice into a variadic function.
The spread operator consists of three dots following the slice in the function call.

func printStrings(strings ...string) {
	for i := 0; i < len(strings); i++ {
		fmt.Println(strings[i])
	}
}

func main() {
    names := []string{"bob", "sue", "alice"}
    printStrings(names...)
}

*/

// Assignment

/*
Complete the sum function to return the sum of all inputs.

*/

func sum(nums ...int) int {
	total := 0
	for i := 0; i < len(nums); i++ {
		total += nums[i]
	}
	return total
}
