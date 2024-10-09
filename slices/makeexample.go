package slices

/*

Most of the time we don't need to think about the underlying array of a slice.
We can create a new slice using the make function:

// func make([]T, len, cap) []T
mySlice := make([]int, 5, 10)

// the capacity argument is usually omitted and defaults to the length
mySlice := make([]int, 5)

NOTE: Slices created with make will be filled with the zero value of the type.

Capacity:
The capacity of a slice is the number of elements in the underlying array,
counting from the first element in the slice. It is accessed using the built-in cap() function:

```
mySlice := []string{"I", "love", "go"}
fmt.Println(cap(mySlice)) // 3
```

*/

// Assignment

/*
	```
	If we know the rough size of a slice before we fill it up, we can make our program faster by creating the slice
	with that size ahead of time so that the Go runtime doesn't need to continuously allocate new underlying arrays of
	larger and larger sizes. By setting the length, the slice can still be resized later,
	but it means we can avoid all the expensive resizing since we know what we'll need.
	```

	Complete the getMessageCosts() function. It takes a slice of messages and returns a slice of message costs.

Preallocate a slice for the message costs of the same length as the messages slice.

Fill the costs slice with costs for each message. The cost in the cost slice should correspond to the message in the messages slice at the same index.
The cost of a message is the length of the message multiplied by 0.01.
*/

func getMessageCosts(messages []string) []float64 {
	costs := make([]float64, len(messages))
	for i := 0; i < len(messages); i++ {
		costs[i] = float64(len(messages[i])) * 0.01
	}
	return costs
}
