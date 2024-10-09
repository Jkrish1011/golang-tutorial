package slices

import (
	"errors"
	"fmt"
)

/*
Slices are built on top of arrays with the feature of flexiblity in sizes.
Arrays are fixed in size. Once you make an array like [10]int you can't add an 11th element.
A slice is a dynamically-sized, flexible view of the elements of an array.

Slices wrap arrays to give a more general, powerful, and convenient interface to sequences of data.
Except for items with explicit dimensions such as transformation matrices, most array programming in Go is done with slices rather than simple arrays.

NOTE: The zero value of slice is nil.

The syntax of slices is/are:

arrayname[lowIndex:highIndex]
arrayname[lowIndex:]
arrayname[:highIndex]
arrayname[:]

Where lowIndex is inclusive and highIndex is exclusive.

consider an array of integers
a := [6, 3, 2, 6, 5]
	 0, 1, 2, 3, 4

b is a slice having acess to [3, 2, 6]. So to declare that

b:= a[1:4] - 4th index is not included.

NOTE: Slices hold references to an underlying array, and if you assign one slice to another, both refer to the same array.
If a function takes a slice argument, any changes it makes to the elements of the slice will be visible to the caller,
 analogous to passing a pointer to the underlying array.


*/

// Assignment

/*
Complete the getMessageWithRetriesForPlan function. It takes a plan variable as input as well as an array of 3 messages.
You've been provided with constants representing the plan types at the top of the file.

If the plan is a pro plan, return all the strings from the messages input in a slice.
If the plan is a free plan, return the first 2 strings from the messages input in a slice.
If the plan isn't either of those, return a nil slice and an error that says unsupported plan.
*/
const (
	planFree = "free"
	planPro  = "pro"
)

func getMessageWithRetriesForPlan(plan string, messages [3]string) ([]string, error) {
	switch plan {
	case planFree:
		return messages[:2], nil
	case planPro:
		return messages[:], nil
	default:
		fmt.Printf("No Matching plan found. Error!")
	}
	return nil, errors.New("No matching plan found")
}
