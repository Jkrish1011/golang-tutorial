package advancedfunctions

import "fmt"

/*
Closures are functions which reference variables outside of it's own function body.
The function may access and assign to the referenced variables
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
