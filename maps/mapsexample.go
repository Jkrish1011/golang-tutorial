package maps

import (
	"errors"
)

/*
Maps are similar to JavaScript objects, Python dictionaries, and Ruby hashes. Maps are a data structure that provides key->value mapping.

The zero value of a map is nil.

We can create a map by using a literal or by using the make() function:

syntax:

make(map[key-type]value-type)

example 1:
// declare a map , and then later add the values into it.
ages := make(map[string]int)
ages["John"] = 37
ages["Mary"] = 24
ages["Mary"] = 21 // overwrites 24

example 2:
// create maps when declaring it.
ages = map[string]int{
  "John": 37,
  "Mary": 21,
}

The len() function works on a map, it returns the total number of key/value pairs.

ages = map[string]int{
  "John": 37,
  "Mary": 21,
}
fmt.Println(len(ages)) // 2

// Assignment

Question:

Complete the getUserMap function. It takes a slice of names and a slice of phone numbers, and returns a map
of name -> user structs and an error. A user struct just contains a user's name and phone number.
The first name in the names slice pairs with the first phone number, and so on.

If the length of names and phoneNumbers is not equal, return an error with the string "invalid sizes".
*/

type user struct {
	name        string
	phoneNumber int
}

func getUserMap(names []string, phoneNumbers []int) (map[string]user, error) {
	if len(names) != len(phoneNumbers) {
		return nil, errors.New("invalid sizes")
	}

	returnMap := make(map[string]user)

	for idx := 0; idx < len(names); idx++ {
		returnMap[names[idx]] = user{
			name:        names[idx],
			phoneNumber: phoneNumbers[idx],
		}
	}
	return returnMap, nil
}
