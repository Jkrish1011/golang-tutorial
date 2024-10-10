package advancedfunctions

import "fmt"

/*
Closures are functions which reference variables outside of it's own function body.
The function may access and assign to the referenced variables

Note:
They just return a function and the logic inside that function is what is important.
the signature of the returned function, what you do with it, and what are the other variables you declare outside the
scope of the return function everything counts.! It's important to think through the requirement before implementing closures.

NOTE:
The function that we are returning has no name, making them anonymous.
*/

// This function will keep concatinating the strings on which it is being passed into
func concatter() func(string) string {
	doc := ""
	return func(s string) string {
		doc += s + " "
		return doc
	}
}

func closureExample() {
	topGearExample := concatter()
	topGearExample("James May")
	topGearExample("and")
	topGearExample("Jeremy Clarkson")
	topGearExample("Was the best")

	fmt.Printf("%s\n", topGearExample("in top gear's entire seasons!"))
}
