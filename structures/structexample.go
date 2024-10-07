package structures

import "fmt"

type car struct {
	model      string
	make       int
	frontWheel wheel
	backWheel  wheel
}

type wheel struct {
	radius   int
	material string
}

func carStruct() car {

	car1 := car{
		model: "VW",
		make:  2023,
		frontWheel: wheel{
			radius:   10,
			material: "alloy",
		},
		backWheel: wheel{
			radius:   12,
			material: "steel",
		},
	}

	// we can initialize an empty structure as well
	car2 := car{}
	fmt.Printf("%+v", car2)
	// We can define the keys later on as well.
	return car1
}
