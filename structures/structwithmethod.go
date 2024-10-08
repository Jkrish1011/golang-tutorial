package structures

import "fmt"

type rectangle struct {
	x int
	y int
}

func (r rectangle) area() int {
	return r.x * r.y
}

// Coding question
/*
We store our user's authentication data inside an authenticationInfo struct. We need a method that can take that data and return a basic authorization string.
The format of the string should be: Authorization: Basic USERNAME:PASSWORD
TODO: Create a method on the authenticationInfo struct called getBasicAuth that returns the formatted string.
*/

type authenticationInfo struct {
	username string
	password string
}

// create the method below
func (a authenticationInfo) getBasicAuth() string {
	return fmt.Sprintf("Authorization: Basic %s:%s", a.username, a.password)
}
