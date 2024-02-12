package main

import (
	"fmt"
	"math"
)

type Square struct {
	Length float64
	Width  float64
}

type Circle struct {
	Radius float64
}

func (s Square) Area() float64 {
	return s.Length * s.Width
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}

type Shape interface {
	Area() float64
}

func Info(s Shape) {
	fmt.Println(s.Area())
}

func main() {
	square := Square{
		Length: 10.0,
		Width:  10.0,
	}
	circle := Circle{
		Radius: 10.0,
	}

	Info(square)
	Info(circle)
}
