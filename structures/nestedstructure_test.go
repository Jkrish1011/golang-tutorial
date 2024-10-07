package structures

import "testing"

func TestNestedStructures(t *testing.T) {
	toSend := messageToSend{
		message: "hello",
		sender: user{
			name:   "jk",
			number: 8129,
		},
		recipient: user{
			name:   "jk2",
			number: 8129,
		},
	}
	var result bool = canSendMessage(toSend)
	if result == false {
		t.Errorf("Error: Cannot send message")
	}
	t.Logf("Can send message")
}
