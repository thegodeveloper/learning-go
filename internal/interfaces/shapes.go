package interfaces

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perimeter() float64
}

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perimeter() float64 {
	return 2 * (r.width + r.height)
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (c circle) diameter() float64 {
	return 2 * c.radius
}

func Shapes(show bool) {
	if show {
		fmt.Println("--- Shapes Interface")

		r := rect{width: 14, height: 7}
		c := circle{radius: 7}

		measure(r)
		measure(c)

		fmt.Println("diameter:", c.diameter())
	}
}

func measure(g geometry) {
	fmt.Println("geometry:", g)
	fmt.Println("area:", g.area())
	fmt.Println("perimeter:", g.perimeter())
}
