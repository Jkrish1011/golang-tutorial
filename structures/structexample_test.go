package structures

import "testing"

func TestCarStruct(t *testing.T) {
	car := carStruct()
	t.Logf("car is %v and make is %d", car.model, car.make)
}
