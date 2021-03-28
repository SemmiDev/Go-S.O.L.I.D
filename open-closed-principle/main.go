package main

import (
	"fmt"
	"math"
)

type (
	circle struct {
		radius float64
	}
	square struct {
		length float64
	}
	triangle struct {
		height float64
		base   float64
	}
	shape interface {
		area() float64
	}
	calculator struct {}
)

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (s square) area() float64 {
	return s.length * s.length
}

func (t triangle) area() float64 {
	return t.base * t.height / 2
}

func (a calculator) areaSum(shapes ...shape) {
	for _, shape := range shapes {
		switch shape.(type) {
			case circle:
				fmt.Println("circle : ", shape.area())
			case square:
				fmt.Println("square : ",shape.area())
			case triangle:
				fmt.Println("rectangle : ",shape.area())
			default:
				fmt.Println("not found")
		}
	}
}

func main(){

	c := circle{radius: 20}
	s := square{length: 37}
	t := triangle{height: 2, base: 10}

	calc := calculator{}
	calc.areaSum(c, s, t)
}
