package variables

import "testing"

func TestVariables(t *testing.T) {
	variables()
	t.Logf("Test passed successfully!")
}

func TestShortHandVariables(t *testing.T) {
	retValue := shortHandVariables()
	if retValue != "Happy birthday! You are now 21 years old!" {
		t.Errorf("Error: Did not get the expected answer!")
	}
}
