package errors

import "testing"

func TestTryCatch(t *testing.T) {

}

func TestAssignment1(t *testing.T) {
	totalCost, error := sendSMSToCouple("This", "couple!")
	if error != nil {
		t.Errorf("Assignment 1 issues : %v", error)
	}
	t.Logf("The total cost is %v", totalCost)
}
