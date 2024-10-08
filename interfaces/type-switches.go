package interfaces

import "fmt"

func printNumericValue(num interface{}) {
	switch v := num.(type) {
	case int:
		fmt.Printf("%T\n", v)
	case string:
		fmt.Printf("%T\n", v)
	default:
		fmt.Printf("%T\n", v)
	}
}

type cost struct {
	toAddress string
	cost      int
}

type sms struct {
	toPhoneNumber int
	cost          int
}

// Assume expense interface is implemented on cost and sms
// func sample(e expense) {
// switch v:= e.(type) {
// case cost:
// 	return v.toAddress
// case sms:
// 	return v.toPhoneNumber
// }
// }
