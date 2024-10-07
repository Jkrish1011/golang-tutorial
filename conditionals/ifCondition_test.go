package conditionals

import "testing"

func TestEmail(t *testing.T) {
	result := ifCondition()
	if !result {
		t.Errorf("Error: Invalid email")
	}
}
