package maps

import "errors"

/*
To Insert an element

```
m[key] = elem
```


To Get an element

```
elem = m[key]
```


To Delete an element

```
delete(m, key)
```


To Check if a key exists

```
elem, ok := m[key]
```

If key is in m, then ok is true and elem is the value as expected.
If key is not in the map, then ok is false and elem is the zero value for the map's element type.

IMPORTANT:
---------
Note on passing maps
Like slices, maps are also passed by reference into functions.
This means that when a map is passed into a function we write, we can make changes to the original, we don't have a copy.

NOTE:
----

map keys may be of any type that is comparable. The language spec defines this precisely, but in short,
comparable types are boolean, numeric, string, pointer, channel, and interface types, and structs or arrays that
contain only those types. Notably absent from the list are slices, maps, and functions; these types cannot be compared using ==,
and may not be used as map keys.

It's obvious that strings, ints, and other basic types should be available as map keys, but perhaps unexpected are struct keys.
Struct can be used to key data by multiple dimensions. For example, this map of maps could be used to tally web page hits by country:

```
hits := make(map[string]map[string]int)
```

This is a map of string to (map of string to int). Each key of the outer map is the path to a web page with its own inner map.
Each inner map key is a two-letter country code.
This expression retrieves the number of times an Australian has loaded the documentation page:
*/

// Assignment

/*
Complete the deleteIfNecessary function.

Check the scheduledForDeletion bool to determine if they are scheduled for deletion or not.

- If the user doesn't exist in the map, return the error not found.
- If they exist but aren't scheduled for deletion, return deleted as false with no errors.
- If they exist and are scheduled for deletion, return deleted as true with no errors and delete their record from the map.

*/

type user2 struct {
	name                 string
	number               int
	scheduledForDeletion bool
}

func deleteIfNecessary(users map[string]user2, name string) (deleted bool, err error) {
	user, ok := users[name]

	if !ok {
		return false, errors.New("not found")
	}

	if user.scheduledForDeletion {
		delete(users, name)
		return true, nil
	}

	return false, nil

}
