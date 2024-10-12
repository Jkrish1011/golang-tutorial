package concurrency

/*
keyword `select`

Used in cases when a single goroutine is listening to multiple channels and wants to process the return values from these
channels in the sequential order in which they are received.

NOTE:
It is similar to `switch` but it's for channels

Example:
--------

```

select {
	case i, ok := <- chanInt:
		fmt.Println(i)

	case s, ok := <- chanString:
		fmt.Println(s)
}

```

In the above example, if the chanInt is ready to send a message, the corresponding case gets executed.
Similarly for chanString is returning values, that particular case gets executed.

Info::
If both have value returned, one will be selected randomly.

Select Default Case:
-------------------

The default case in a select statement executes immediately if no other channel has a value ready.
Note: A default case stops the select statement from blocking.


*/

// Assignment

/*
Complete the logMessages function.

Use an infinite for loop and a select statement to log the emails and sms messages as they come in order across the two channels.
Add a condition to return from the function when one of the two channels closes, whichever is first.

Use the logSms and logEmail functions to log the messages.

*/

import (
	"fmt"
	"time"
)

func logMessages(chEmails chan string, chSms chan string) {
	// The for loop will go on for until on of the channels closes(close logic is inside the sendToLogger())
	for {
		select {
		case s, ok := <-chSms:
			if !ok {
				return
			}
			logSms(s)
		case e, ok := <-chEmails:
			if !ok {
				return
			}
			logEmail(e)
		}
	}
}

// don't touch below this line

func logSms(sms string) {
	fmt.Println("SMS:", sms)
}

func logEmail(email string) {
	fmt.Println("Email:", email)
}

// Assignment 2 - default case

/*

Complete the saveBackups function.

It should read values from the snapshotTicker and saveAfter channels simultaneously and continuously.

If a value is received from snapshotTicker, call takeSnapshot()
If a value is received from saveAfter, call saveSnapshot() and return from the function: you're done.
If neither channel has a value ready, call waitForData() and then time.Sleep() for 500 milliseconds. After all, we want to show in the logs that the snapshot service is running.

*/

func saveBackups(snapshotTicker, saveAfter <-chan time.Time, logChan chan string) {
	for {
		select {
		case <-snapshotTicker:
			takeSnapshot(logChan)
		case <-saveAfter:
			saveSnapshot(logChan)
			return
		default:
			waitForData(logChan)
			time.Sleep(time.Millisecond * 500)
		}
	}
}

// don't touch below this line

func takeSnapshot(logChan chan string) {
	logChan <- "Taking a backup snapshot..."
}

func saveSnapshot(logChan chan string) {
	logChan <- "All backups saved!"
	close(logChan)
}

func waitForData(logChan chan string) {
	logChan <- "Nothing to do, waiting..."
}
