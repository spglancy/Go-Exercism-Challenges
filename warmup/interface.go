package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perimeter() float64
}

type circle struct {
	radius float64
}

type rectangle struct {
	height, width float64
}

func (c circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

func (c circle) perimeter() float64 {
	return 2 * c.radius * math.Pi
}

func (r rectangle) area() float64 {
	return r.height * r.width
}

func (r rectangle) perimeter() float64 {
	return (2 * r.height) + (2 * r.width)
}

func measure(g geometry) {
	fmt.Println(g.area())
	fmt.Println(g.perimeter())
}

func main() {
	circle := circle{radius: 4}
	rectangle := rectangle{height: 10, width: 10}
	measure(circle)
	measure(rectangle)
}
