package main

import (
	"fmt"
	"math"
)

type Square struct {
	Length float64
}

type Cube struct {
	Square
}

type Shape interface {
	area() float64
}

type Object interface {
	Shape
	volume() float64
}

func (s *Square) area() float64 {
	return math.Pow(s.Length, 2)
}

func (c *Cube) volume() float64 {
	return math.Pow(c.Length, 3)
}

func areaSum(shapes ...Shape) float64 {
	var sum float64
	for _, s := range shapes {
		sum += s.area()
	}
	return sum
}

func areaVolumeSum(shapes ...Object) float64 {
	var sum float64
	for _, s := range shapes {
		sum += s.area() + s.volume()
	}
	return sum
}

func main() {
	s1 := Square{Length: 5}
	s2 := Square{Length: 6}
	c1 := Cube{Square: Square{Length: 3}}
	c2 := Cube{Square: Square{Length: 4}}

	fmt.Println(areaSum(&s1, &s2, &c1, &c2))
	fmt.Println(areaVolumeSum(&c1, &c2))
}
