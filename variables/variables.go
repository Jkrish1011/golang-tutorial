package variables

import "fmt"

// Initialize the variables from the print statement to int, float64, bool and string with their zero values, respectively.

func variables() {
	// initialize variables here
	var smsSendingLimit int
	var costPerSMS float64
	var hasPermission bool
	var username string
	fmt.Printf("%v %.2f %v %q\n", smsSendingLimit, costPerSMS, hasPermission, username)
}

// Complete the main function. It should print: "Happy birthday! You are now 21 years old!".
// Create a string variable messageStart with the text "Happy birthday! You are now"
// Create an integer variable age set to 21
// Create another string variable messageEnd with the text "years old!"
// The provided fmt.Println statement will print the full message on a single line separated by spaces.

func shortHandVariables() (returnValue string) {
	messageStart := "Happy birthday! You are now"
	age := 21
	messageEnd := "years old!"
	fmt.Println(messageStart, age, messageEnd)
	returnValue = fmt.Sprintf("%s %d %s", messageStart, age, messageEnd)
	return returnValue
}

func sameLineDeclaration() {
	averageOpenRate, displayMessage := 0.23, "is the average open rate of your message"
	fmt.Println(averageOpenRate, displayMessage)
}

func convertions() float64 {
	temperatureInt := 88
	temperatureFloat := float64(temperatureInt)
	return temperatureFloat
}
