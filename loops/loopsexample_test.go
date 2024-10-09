package loops

import "testing"

func TestLoopsExample(t *testing.T) {
	loopExample()
}

func TestAssignment1(t *testing.T) {
	numOfMessage := 50
	cost := bulkSend(numOfMessage)
	t.Logf("The cost for sending %d message is %v", numOfMessage, cost)
}
