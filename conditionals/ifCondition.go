package conditionals

import "fmt"

func getLength(email string) int {
	return len(email)
}

func ifCondition() bool {
	var email string = "jk@jk.com"
	// this if method will not allow `length` to be used after the scope of the if block
	if length := getLength(email); length < 1 {
		fmt.Println("email is invalid")
		return false
	}
	return true
}
