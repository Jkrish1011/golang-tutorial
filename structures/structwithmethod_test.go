package structures

import "testing"

func TestStructWithMethod(t *testing.T) {
	rect := rectangle{
		x: 5,
		y: 10,
	}
	if rect.area() != 50 {
		t.Errorf("Error: Area calculation is not correct!")
	}
}
