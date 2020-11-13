package main

import "fmt"

type geometry interface {
	area() float64
	perim() float64
}

func area(g geometry) float64 {
	return g.area()
}

func perim(g geometry) float64 {
	return g.perim()
}

type rect struct {
	width, height float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2 * (r.height + r.width)
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return 3.14 * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * 3.14 * c.radius
}

func main() {
	r1 := rect{3, 4}
	c1 := circle{4}

	fmt.Println(area(r1))
	fmt.Println(area(c1))
}
