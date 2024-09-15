package main

import (
	"fmt"
	"math"
)

type Circle struct {
	Radius float64
}

type Rectangle struct {
	length float64
	width  float64
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}

func (r Rectangle) Area() float64 {
	return r.length * r.width
}

func PrintArea(s Shape) string {
	return fmt.Sprintf("%v", s.Area())
}

type Shape interface {
	Area() float64
}

func main() {
	rectangle := Rectangle{10, 12}
	circle := Circle{7}

	fmt.Println(rectangle.Area())
	fmt.Println(circle.Area())
}
