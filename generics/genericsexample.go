package generics

/*

Generics are used to produce repetable code for various different types, i.e., Generics reduce repetitive code.

It is be much more common to see generics used in libraries and packages.
Libraries and packages contain importable code intended to be used in many applications, so it makes sense to write them in a more abstract way.
Generics are often the way to do just that!

example:

func FunctionName[T any](s []T) ([]T, []T) {
	// Here it accepts a type T of any type.
	// The input arguement should be a [](slice) of type T
	// The return arguments should be 2 slices of type T
	// Any kind of similar signatures and return types can be generated according to the logic which we wish to implement.
}

*/

// Assignment - 1

/*

Complete the getLast() function. It should be a generic function that returns the last element from a slice,
no matter the types stored in the slice. If the slice is empty, it should return the zero value of the type.

*/

func getLast[T any](s []T) T {
	if len(s) == 0 {
		var zeroVal T
		return zeroVal
	}
	return s[len(s)-1]
}
