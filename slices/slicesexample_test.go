package slices

import "testing"

func TestSlices(t *testing.T) {
	planName := "pro"
	messages := [3]string{
		"first",
		"second",
		"third",
	}
	retMessages, err := getMessageWithRetriesForPlan(planName, messages)
	if err != nil {
		t.Errorf("error: %v", err)
	}
	t.Logf("Returned results are : %+v", retMessages)
}
