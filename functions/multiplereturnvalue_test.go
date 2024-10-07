package functions

import "testing"

func TestMultipleReturnValue(t *testing.T) {
	val1, val2 := multipleReturnValue()
	// val1 , _ := multipleReturnValue() in this case val2 is ignored!

	if val1+val2 != "helloworld" {
		t.Errorf("Error: not correct")
	}

}
