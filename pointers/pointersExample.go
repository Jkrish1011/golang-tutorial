package pointers

import "strings"

/*

x := 5 // stores are a location in RAM
y := x // y gets a new copy of value 5 in RAM
z := &x // z will store the address of x. any updates to z will reflect in x.
*z = 6 // will update the value of z and x both. We have updated the value of x without even accessing x.

Syntax:

```
var p *int
```

A pointer's zero value is nil
The & operator generates a pointer to its operand

```
myString := "hello"
myStringPointer = &myString
```

* dereference a pointer to gain access to the underlying value
```
fmt.Println(*myStringPtr) // reads myString through the pointer
*myStringPtr = "world" // set myString through the pointer
```

*/

// Assignment

/*
Complete the removeProfanity function.

It should use the strings.ReplaceAll function to replace all instances of the following words in the input message with asterisks.

Word	Replacement
fubb	****
shiz	****
witch	*****
It should update the value in the pointer and return nothing.

Do not alter the function signature.

NOTE:
if a pointer points to nothing(the zero value of the pointer type) then dereferencing it will cause a runtime error crashing the program.

There are Pointer receiver and Non-Pointer Receiver in Go

example:
type car struct {
	color string
}

// This will change the color of the car object
func (c *car) setColor(color string) {
	c.color = color
}

// This will NOT change the color of the car object
func (c car) setColor(color string) {
	c.color = color
}
*/

func removeProfanity(message *string) {
	msgValue := *message
	msgValue = strings.ReplaceAll(msgValue, "fubb", "****")
	msgValue = strings.ReplaceAll(msgValue, "shiz", "****")
	msgValue = strings.ReplaceAll(msgValue, "witch", "*****")
	*message = msgValue
	// message = &msgValue
}
