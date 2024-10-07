package structures

import "testing"

func TestEmbeddedStructures(t *testing.T) {
	laneTruck := truck{
		car: car{
			make:  20203,
			model: "Benz",
		},
		bedSize: 1000,
	}
	if laneTruck.make != 20203 {
		t.Errorf("Error!")
	}
}
