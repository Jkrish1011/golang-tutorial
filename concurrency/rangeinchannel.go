package concurrency

import (
	"fmt"
	"time"
)

/*
We can iterate over the channels using for loop and it will block and wait to receive an item.

```
for item := chan {
	...
}
```

This example will receive values over the channel (blocking at each iteration if nothing new is there)
and will exit only when the channel is closed.
*/

// Assignment

/*
Complete the concurrentFib function. It should:

Create a new channel of ints
Call fibonacci concurrently
Use a range loop to read from the channel and append the values to a slice
Return the slice

*/

func concurrentFib(n int) []int {
	fibChan := make(chan int)
	retSlice := []int{}
	go fibonacci(n, fibChan)

	for item := range fibChan {
		retSlice = append(retSlice, item)
	}

	return retSlice
}

// don't touch below this line

func fibonacci(n int, ch chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x+y
	}
	close(ch)
}

// Example Barista Coffee example

func baristaExample(orderChan chan string) {
	for order := range orderChan {
		fmt.Printf("Preparing %s \n", order)
		time.Sleep(2 * time.Second)
		fmt.Printf("Served %s \n", order)
	}
}
