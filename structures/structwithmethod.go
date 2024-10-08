package structures

type rectangle struct {
	x int
	y int
}

func (r rectangle) area() int {
	return r.x * r.y
}
