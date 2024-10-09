package loops

/*
There is NO WHILE LOOP in GO

Go uses the for loop syntax to execute the while loop.

for CONDITION {
	// Implement your while logic here!
}

example:

plantHeight := 1

for planHeight < 5 {
	fmt.Println("still growing! current height is : %v inches ", plantHeight)
}
fmt.Println("Plant has grown to height of : %v inches", plantHeight)
*/

// Assignment

/*

We have an interesting new cost structure from our SMS vendor.
They charge exponentially more money for each consecutive text we send! Let's write a function that calculates how many messages
 we can send in a given batch given a costMultiplier and a maxCostInPennies.

In a nutshell, the first message costs a penny, and each message after that costs the same as the previous message multiplied
by the costMultiplier.

There is a bug in the code! Add a condition to the for loop to fix the bug. The loop should stop when balance is equal to
or less than 0.

Continue keyword:
The continue keyword stops the current iteration of a loop and continues to the next iteration. contineu is a powerful way to use the
"guard clause" pattern within loops.
example/
for i := 0; i < 10; i++ {
	if i % 2 == 0 {
		continue
	}
	fmt.Println(i)
}

Break Keyword:
the break keyword stop the current iteration of a loop and exists the loop

for i := 0; i < 10; i++ {
	if i == 6 {
		break
	}
	fmt.Println(i)
}
*/

func getMaxMessagesToSend(costMultiplier float64, maxCostInPennies int) int {
	actualCostInPennies := 1.0
	maxMessagesToSend := 0
	// balance := float64(maxCostInPennies) - actualCostInPennies
	for actualCostInPennies <= float64(maxCostInPennies) {
		maxMessagesToSend++
		actualCostInPennies *= costMultiplier
		// balance -= actualCostInPennies

	}
	// if balance <= 0 {
	// 	maxMessagesToSend--
	// }
	return maxMessagesToSend
}
