package functions

import "testing"

func TestEarlyReturns(t *testing.T) {
	quotient, error := dividingFunction(2, 1)
	// val1 , _ := multipleReturnValue() in this case val2 is ignored!
	if error != nil {
		t.Errorf("Error!")
	}

	t.Logf("Quotient is %v", quotient)

}
