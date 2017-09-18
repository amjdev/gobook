package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x float64
	y float64
	r float64
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

type Shape interface {
	area() float64
	perimeter() float64
}

type Multishape struct {
	shapes []Shape
}

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}

func (r *Rectangle) perimeter() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return (2*l + 2*w)
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func (c *Circle) perimeter() float64 {
	return math.Pi * 2 * c.r
}

func (m *Multishape) area() float64 {
	var area float64
	for _, s := range m.shapes {
		area += s.area()
	}
	return area
}

func (m *Multishape) perimeter() float64 {
	var perimeter float64
	for _, s := range m.shapes {
		perimeter += s.perimeter()
	}
	return perimeter
}

func main() {
	c := Circle{x: 0, y: 0, r: 5}
	fmt.Println("Área del círculo = ", c.area())
	fmt.Println("Perímetro del círculo = ", c.perimeter())

	r := Rectangle{0, 0, 10, 10}
	fmt.Println("Área del rectángulo = ", r.area())
	fmt.Println("Perímetro del rectángulo = ", r.perimeter())

	s := []Shape{&c, &r}
	m := Multishape{s}
	fmt.Println("Área total = ", m.area())
	fmt.Println("Perímetro total = ", m.perimeter())
}
