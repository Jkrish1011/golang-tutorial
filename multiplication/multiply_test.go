package multiplication

import "testing"

func TestMultiply(t *testing.T) {
	testCases := []struct {
		a, b, expected int
	}{
		{2, 3, 6},
		{0, 5, 0},
		{-4, 2, -8},
	}

	for _, tc := range testCases {
		result := multiply(tc.a, tc.b)
		if result != tc.expected {
			t.Errorf("multiply(%d, %d) expected %d, got %d", tc.a, tc.b, tc.expected, result)
		}
	}
	t.Logf("Test passed successfully!")
}
