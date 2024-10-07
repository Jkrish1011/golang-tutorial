package structures

import "testing"

func TestCarStruct(t *testing.T) {
	car := carStruct()
	t.Logf("car is %v and make is %d", car.model, car.make)
	t.Logf("frontWheel is %+v and backWheel is %+v", car.frontWheel, car.backWheel)
}
