package concurrency

/*
Go was designed to be concurrent, which is a trait fairly unique to go.

keyword `go`
concurrency is as simple as using the keyword `go` while calling a function.

```
go doSomething()
```

here doSomething() will be executed concurrently with the rest of the code in the function.
they go keyword spawns a new `goroutine`


*/

// Assignment

/*
Edit the sendEmail() function to execute its anonymous function concurrently so that the "received" message
prints after the "sent" message.
*/

import (
	"fmt"
	"time"
)

func sendEmail(message string) {
	// Adding go here would spawn a new gorountine and it will be executed separately. thus Email send will be printed first and then this
	// function because it has a sleep timer inside it(for the purpose of the assignment)
	go func() {
		time.Sleep(time.Millisecond * 250)
		fmt.Printf("Email received: '%s'\n", message)
	}()
	fmt.Printf("Email sent: '%s'\n", message)
}

// Don't touch below this line

func test(message string) {
	sendEmail(message)
	time.Sleep(time.Millisecond * 500)
	fmt.Println("========================")
}
