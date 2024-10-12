package concurrency

import (
	"fmt"
	"slices"
	"testing"
)

func TestRangeInChannel(t *testing.T) {
	type testCase struct {
		n        int
		expected []int
	}
	tests := []testCase{
		{5, []int{0, 1, 1, 2, 3}},
		{3, []int{0, 1, 1}},
	}
	if withSubmit {
		tests = append(tests, []testCase{
			{0, []int{}},
			{1, []int{0}},
			{7, []int{0, 1, 1, 2, 3, 5, 8}},
		}...)
	}

	passCount := 0
	failCount := 0

	for _, test := range tests {
		actual := concurrentFib(test.n)
		if !slices.Equal(actual, test.expected) {
			failCount++
			t.Errorf(`
---------------------------------
Test Failed:
  n:        %v
  expected: %v
  actual:   %v
`, test.n, test.expected, actual)
		} else {
			passCount++
			fmt.Printf(`
---------------------------------
Test Passed:
  n:        %v
  expected: %v
  actual:   %v
`, test.n, test.expected, actual)
		}
	}

	fmt.Println("---------------------------------")
	fmt.Printf("%d passed, %d failed\n", passCount, failCount)
}

func TestBaristaExample(t *testing.T) {
	numberOfOrders := 5
	orderChan := make(chan string, numberOfOrders)

	go func() {
		for idxOrder := 0; idxOrder < numberOfOrders; idxOrder++ {
			order := fmt.Sprintf("Coffee Order #%d", idxOrder)
			orderChan <- order
			fmt.Println("Order Placed : ", order)
		}
		close(orderChan)
	}()

	// go baristaExample(orderChan) - if you do this, then the order's won't be served because the main goroutine would finish compiling and
	// won't be around to read from the channel.
	baristaExample(orderChan)
}
