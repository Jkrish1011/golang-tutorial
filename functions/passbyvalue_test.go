package functions

import "testing"

func TestPassByValue(t *testing.T) {
	var x int = 5
	x = passByValueIncrement(x)
	if x != 6 {
		t.Errorf("Error: Failed! x is not 6")
	}

}
