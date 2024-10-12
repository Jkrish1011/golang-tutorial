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
)

func logMessages(chEmails, chSms chan string) {
	// select {
	// case em, ok := <-
	// }

}

// don't touch below this line

func logSms(sms string) {
	fmt.Println("SMS:", sms)
}

func logEmail(email string) {
	fmt.Println("Email:", email)
}
