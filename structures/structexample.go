package structures

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
	return car1
}
