package advancedfunctions

import "testing"

func TestAdvancedFunctionsAdd(t *testing.T) {
	total := aggregate(4, 5, 6, add)
	t.Logf("Total is : %v", total)
}

func TestAdvancedFunctionsMul(t *testing.T) {
	total := aggregate(4, 5, 6, mul)
	t.Logf("Product is : %v", total)
}
