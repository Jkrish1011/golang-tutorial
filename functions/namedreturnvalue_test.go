package functions

import "testing"

func TestNamedReturnFunctionValue(t *testing.T) {
	x, y := namedReturnValues(1, 2)

	if x != 2 && y != 1 {
		t.Errorf("Error: numbers were not swapped")
	}
}
