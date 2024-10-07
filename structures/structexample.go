package structures

type car struct {
	model string
	make  int
}

func carStruct() car {
	car1 := car{
		model: "VW",
		make:  2023,
	}
	return car1
}
