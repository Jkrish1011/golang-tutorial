package errors

import (
	"fmt"
)

/*
Complete the sendSMSToCouple function. It should send 2 messages, first to the customer and then to the customer's spouse.

1/ Use sendSMS() to send the msgToCustomer. If an error is encountered, return 0 and the error.
2/ Do the same for the msgToSpouse
3/ If both messages are sent successfully, return the total cost of the messages added together.

Note; When you return a non-nil error in Go, it's conventional to return the "zero" values of all other return values.
*/

func sendSMSToCouple(msgToCustomer, msgToSpouse string) (int, error) {
	costToCustomer, errorToCustomer := sendSMS(msgToCustomer)
	if errorToCustomer != nil {
		return 0, errorToCustomer
	}

	costToSpouse, errorToSpouse := sendSMS(msgToSpouse)
	if errorToSpouse != nil {
		return 0, errorToCustomer
	}
	return costToCustomer + costToSpouse, nil

}

func sendSMS(message string) (int, error) {
	const maxTextLen = 25
	const costPerChar = 2
	if len(message) > maxTextLen {
		return 0, fmt.Errorf("can't send texts over %v characters", maxTextLen)
	}
	return costPerChar * len(message), nil
}
