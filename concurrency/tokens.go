package concurrency

/*
Token:
-----
Token is a unary value. We don't care what is passed through the channel, we care when and if it is passed through channels.
Empty Struct are usually treated as `tokens`.

We can block and wait until something is sent on a channel using the below syntax:

```
<-ch
```

This will blocks until something is sent into the channel and once it is done, it discards the value and continue the execution.
*/

// Assignment

/*

Run the code. It never exits! The channel passed to waitForDBs stays blocked, because it's only popping the first value off the channel.

Fix the waitForDBs function. It should pause execution until it receives a token for every database from the dbChan channel.
Each time waitForDBs reads a token, the getDBsChannel goroutine will print a message to the console for you.

The succinctly named numDBs input is the total number of databases. Look at the test code to see how these functions are used
so you can understand the control flow.

*/

import "fmt"

func waitForDBs(numDBs int, dbChan chan struct{}) {
	for idxNumDB := 0; idxNumDB < numDBs; idxNumDB++ {
		<-dbChan
	}
}

// don't touch below this line
/*
This function returns the pointer variable which points to the address(reference of count variable)
This will allow us to receive the same address reference, and use pointer to get the value and go ahead. Memory Optimization.

--------------------------------------------------------------------------------------------------------------------------------------------------------------------
| &count: This expression takes the address of the count variable. It essentially returns a pointer to the memory location where the count variable is stored.      |
| *int: This type represents a pointer to an integer value. It indicates that the variable will hold the memory address of an integer.								|
--------------------------------------------------------------------------------------------------------------------------------------------------------------------
*/
func getDBsChannel(numDBs int) (chan struct{}, *int) {
	fmt.Printf("getDBsChannel::numDBs : %v\n", numDBs)
	count := 0
	ch := make(chan struct{})

	go func() {
		for i := 0; i < numDBs; i++ {
			ch <- struct{}{}
			fmt.Printf("Database %v is online\n", i+1)
			count++
		}
	}()

	return ch, &count
}
