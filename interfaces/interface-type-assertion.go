package interfaces

/*
	To access the underlying type of an interface value.
	we can cast an interface to its underlying type using a type assertion
*/

type shape interface {
	area() float64
}

type circle struct {
	radium float64
}

func typeAssertionExample(s shape) {
	// c is a new circle cast from s
	// ok is a bool that is true if s was a circle or false otherwise
	// c, ok := s.(circle)

}
