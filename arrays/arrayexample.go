package arrays

import "fmt"

/*
Arrays are fixed-size groups of variables of the same type
the type ([n]T) is an array of n values of type T.

declaration:

var myInts [10]int

or

primes := [6]int {2, 3, 5, 7, 11, 13}

primes := [6]int {2, 3, 5, 7, 11} - in this case, 6th index will be zero, i.e., default value for the type declared
*/

func arrayExample() {
	primes := [6]int{2, 3, 5, 7, 11}
	fmt.Printf("%+v", primes)
}

/*

ASSIGNMENT:

Complete the getMessageWithRetries function. It takes three strings and returns:

An array of 3 strings
An array of 3 integers
The returned string array contains the original messages. The first is the primary message, the second is the first reminder, and the third is the last reminder.
The integers represent the cost of sending each message.

The cost of each message is equal to the length of the message, plus the length of any previous messages. For example:

"hello" costs 5
"world" costs 5, adding "hello" makes total cost 10 (5 + 5)
"!" costs 1, adding the previous messages makes total cost 11 (5 + 5 + 1)

*/

func getMessageWithRetries(primary, secondary, tertiary string) ([3]string, [3]int) {
	tmpArray := [3]string{primary, secondary, tertiary}
	tmpCost := [3]int{}
	for i := 0; i < 3; i++ {
		tmpCost[i] = len(tmpArray[i])
	}
	return tmpArray, tmpCost
}
