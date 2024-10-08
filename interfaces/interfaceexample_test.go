package interfaces

import (
	"testing"
	"time"
)

func TestInterfaces(t *testing.T) {
	reply, cost := sendMessage(birthdayMessage{
		birthdayTime:  time.Date(1994, 03, 21, 0, 0, 0, 0, time.UTC),
		recipientName: "Vebi",
	})
	t.Logf("Success: the reply : \"%v\" and the cost was : \"%v\"", reply, cost)

	reply2, cost2 := sendMessage(sendingReport{
		reportName:    "ANNUAL-Revenue-Report",
		numberOfSends: 30,
	})
	t.Logf("Success: the reply : \"%v\" and the cost was : \"%v\"", reply2, cost2)
}
