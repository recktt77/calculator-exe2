package calculator_exe2

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Length float64
	Width  float64
}

type Circle struct {
	Radius float64
}

type Square struct {
	Length float64
}

type Triangle struct {
	SideA float64
	SideB float64
	SideC float64
}

func (r Rectangle) Area() float64 {
	return r.Length * r.Width
}
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Length)
}

func (c Circle) Area() float64 {
	return c.Radius * c.Radius * math.Pi
}
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func (s Square) Area() float64 {
	return s.Length * s.Length
}

func (s Square) Perimeter() float64 {
	return s.Length * 4
}

func (t Triangle) Area() float64 {
	semiPerimeter := (t.SideA + t.SideB + t.SideC) / 2
	area := math.Sqrt(semiPerimeter * (semiPerimeter - t.SideA) * (semiPerimeter - t.SideB) * (semiPerimeter - t.SideC))
	return area
}
func (t Triangle) Perimeter() float64 {
	return t.SideA + t.SideB + t.SideC
}

func PrintShapeDetails(s Shape) {
	fmt.Println(s.Area())
	fmt.Println(s.Perimeter())
}
