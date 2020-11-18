package main

import (
	"fmt"

	"../circle"
	"../rect"
)

type Shape interface {
	Area() float32
	PartialArea(float32) float32
}

func Area(s Shape) float32 {
	return s.Area()
}

func PartialArea(s Shape, extra float32) float32 {
	fmt.Println("Im here")
	return s.PartialArea(extra)
}

// type circle struct {
// 	radius float32
// }

// func (c circle) area() float32 {
// 	return 3.14 * c.radius * c.radius
// }

// func (c circle) partialArea(angle float32) float32 {
// 	return (angle / 360) * 3.14 * c.radius * c.radius
// }

// type rectangle struct {
// 	width  float32
// 	height float32
// }

// func (r rectangle) area() float32 {
// 	return r.width * r.height
// }

// func (r rectangle) partialArea(partialHeight float32) float32 {
// 	return (partialHeight * r.width)
// }

func main() {
	c1 := circle.MakeCircle(3.4)
	r1 := rect.MakeRect(5, 6)
	fmt.Println(r1)

	fmt.Println(Area(c1))
	fmt.Println(Area(r1))
	fmt.Println(PartialArea(c1, 180))
	fmt.Println(PartialArea(r1, 2.5))
}
