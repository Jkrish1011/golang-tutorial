package division

import "testing"

func TestDivision(t *testing.T) {
	testCases := []struct {
		a, b, expected int
	}{
		{6, 3, 2},
		{0, 5, 0},
		{-4, 2, -2},
	}

	for _, tc := range testCases {
		result := divide(tc.a, tc.b)
		if result != tc.expected {
			t.Errorf("Division(%d, %d) expected %d, got %d", tc.a, tc.b, tc.expected, result)
		}
	}
	t.Logf("Test passed successfully!")
}
