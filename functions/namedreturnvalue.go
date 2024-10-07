package functions

func namedReturnValues(a, b int) (x, y int) { // if two variable of same type are adjecent we can use the shorthand method to indicate their type
	// here x and y are assigned zero as default value
	x = b
	y = a
	return x, y
}
