package slices

/*
RANGE
-----

Go provides syntactic sugar to iterate easily over elements of a slice;

```
for INDEX, ELEMENT := range SLICE {
	// This is like enumerate. Get index and the element present at that Index readily.
}
```

example:

fruits := []string {"banana", "grape", "orange"}

for idx, fruit := range fruits {
	fmt.Println("Idx: %d, Value: %s", idx, fruit)
	// Output will be as shown below
	0 banana
	1 grape
	2 orange
}
*/

// Assignment

/*
Complete the indexOfFirstBadWord function. If it finds any bad words in the message it should return the index of the
first bad word in the msg slice. This will help us filter out naughty words from our messaging system.
If no bad words are found, return -1 instead.

Use the range keyword.
*/
func indexOfFirstBadWord(msg []string, badWords []string) int {
	for idxMsg, word := range msg {
		for _, badWord := range badWords {
			if word == badWord {
				return idxMsg
			}
		}
	}
	return -1
}
