package interfaces

import (
	"fmt"
	"time"
)

/*
Info:
Interfaces are just collections of method signatures.
A type "implements" an interface if it has methods that match the interface's method signatures.
you don't need to explicitly mention that a struct implements it, rather the go compiler infers that information own it's own.

Interfaces are not classes
1. Interfaces are not classes, they are slimmer
2. Interfaces don't have constructors or deconstructors that require that data is created or destroyed
3. Interfaces aren't hierarchical by nature, though there is synatic sugar to create interfaces that happen
	To be supersets of other interfaces.
4. Interfaces define function signatures, but not underlying behavior.
	Making an interface often won't DRY up your code in regards to struct methods.

Question:

The birthdayMessage and sendingReport structs already have implementations of the getMessage method.
The getMessage method returns a string, and any type that implements the method can be considered a message
(meaning it implements the message interface).

1/ Add the getMessage() method signature as a requirement on the message interface.
2/ Complete the sendMessage function. It should return:
	2.1/ The content of the message.
	2.2/ The cost of the message, which is the length of the message multiplied by 3.

*/

func sendMessage(msg message) (string, int) {
	var message string = msg.getMessage()
	return message, len(message) * 3
}

type message interface {
	getMessage() string
}

// don't edit below this line

type birthdayMessage struct {
	birthdayTime  time.Time
	recipientName string
}

func (bm birthdayMessage) getMessage() string {
	return fmt.Sprintf("Hi %s, it is your birthday on %s", bm.recipientName, bm.birthdayTime.Format(time.RFC3339))
}

type sendingReport struct {
	reportName    string
	numberOfSends int
}

func (sr sendingReport) getMessage() string {
	return fmt.Sprintf(`Your "%s" report is ready. You've sent %v messages.`, sr.reportName, sr.numberOfSends)
}
