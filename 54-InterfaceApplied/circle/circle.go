package circle

type Circle struct {
	radius float32
}

func (c Circle) Area() float32 {
	return 3.14 * c.radius * c.radius
}

func (c Circle) PartialArea(angle float32) float32 {
	return (angle / 360) * 3.14 * c.radius * c.radius
}

func MakeCircle(rad float32) *Circle {
	return &Circle{radius: rad}
}
