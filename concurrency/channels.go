package concurrency

/*

Channels
--------
The way we can declare new goroutines and allow them to compute some tasks and share the computated data back.

Channels are typed, thread-safe queue. It allows different goroutines to communicate with each other.

Important: Like maps, and slices channels must be created before USE. They also use the same keyword `make`.

Syntax:
------

```
ch := make(chan int)
```

Send data into the channel:
--------------------------
```
ch <- 69
```

`<-` is called the channel operator. the data flows in the direction of the arrow.
The operation will block until a new channel is ready to receive the value.

Receive data from a channel:
--------------------------

```
v :=  <-ch
```

This reads and receives a value from the channel and saves it into the variable `v`.
The operation blocks until there is a value from the channel to be received.

Deadlock
--------
A deadlock is when a group of goroutines are blocking and none of them are able to continue execution.

Read-Only Channels:
-------------------
A channel can be marked as read-only by casting it from a `chan` to a <- chan type. For Example.

```
func main() {
	ch := make(chan int)
	readCh(ch)
}

func readCh(ch <-chan int) {
	// ch can only be read from in this function
	...
}

```

Write-Only Channels:
-------------------
Everything remains the same. Just that the arrow position moves in the opposite direction.

```
func writeCh(ch chan<- int) {
	// ch can only be written to in this function
	...
}
```

*/

// Assignment

/*

Run the program. You'll see that it deadlocks and never exits.
The sendIsOld function is trying to send on a channel, but no other goroutines are running that can accept the value from the
channel.

Fix the deadlock by spawning a goroutine to send the "is old" values.

Closing go channel:
-------------------
Channels can be explicitly close by the sender

```
close(ch)
```

Check if channel is close is similar to checking value exists in maps.

```
value, ok := <-ch
```

ok is `false`, channel is empty and closed.
value contains the value that is being sent from the channel

*/
import (
	"time"
)

type email struct {
	body string
	date time.Time
}

func checkEmailAge(emails [3]email) [3]bool {
	isOldChan := make(chan bool)

	go sendIsOld(isOldChan, emails)

	isOld := [3]bool{}
	isOld[0] = <-isOldChan
	isOld[1] = <-isOldChan
	isOld[2] = <-isOldChan
	return isOld
}

// don't touch below this line

func sendIsOld(isOldChan chan<- bool, emails [3]email) {
	for _, e := range emails {
		if e.date.Before(time.Date(2020, 0, 0, 0, 0, 0, 0, time.UTC)) {
			isOldChan <- true
			continue
		}
		isOldChan <- false
	}
}

// Assignment 2

/*
The sendReports function sends out a batch of reports to our clients and reports back how many were sent across a channel.
It closes the channel when it's done.

Complete the countReports function. It should:

Use an infinite for loop to read from the channel:
If the channel is closed, break out of the loop
Otherwise, keep a running total of the number of reports sent
Return the total number of reports sent
*/

func countReports(numSentCh chan int) int {
	total := 0
	for {
		value, ok := <-numSentCh
		if !ok {
			break
		}
		total += value
	}
	return total
}

// don't touch below this line

func sendReports(numBatches int, ch chan int) {
	for i := 0; i < numBatches; i++ {
		numReports := i*23 + 32%17
		ch <- numReports
	}
	close(ch)
}
