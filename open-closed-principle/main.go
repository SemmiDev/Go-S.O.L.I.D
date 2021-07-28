package main

import (
	"fmt"
	"math"
)

type circle struct {
	Radius float64
}

type square struct {
	Length float64
}

type triangle struct {
	Height float64
	Base   float64
}

type calculator struct{}

type shape interface {
	area() float64
}

func (c circle) area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (s square) area() float64 {
	return s.Length * s.Length
}

func (t triangle) area() float64 {
	return t.Base * t.Height / 2
}

func (a calculator) areaSum(shapes ...shape) {
	for _, shape := range shapes {
		switch shape.(type) {
		case *circle:
			fmt.Println("circle : ", shape.area())
		case *square:
			fmt.Println("square : ", shape.area())
		case *triangle:
			fmt.Println("rectangle : ", shape.area())
		default:
			fmt.Println("not found")
		}
	}
}

func main() {

	c := &circle{Radius: 20}
	s := &square{Length: 37}
	t := &triangle{Height: 2, Base: 10}

	calc := calculator{}
	calc.areaSum(c, s, t)
}
