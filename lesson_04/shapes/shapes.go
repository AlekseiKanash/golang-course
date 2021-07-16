package shapes

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func (c Circle) String() string {
	return fmt.Sprintf("Circle: radius: %.2f\nArea: %.2f\nPerimeter %.2f", c.Radius, c.Area(), c.Perimeter())
}

type Rectangle struct {
	Height float64
	Width  float64
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

func (r Rectangle) Perimeter() float64 {
	return 2*r.Height + 2*r.Width
}

func (r Rectangle) String() string {
	return fmt.Sprintf("Rectangle with height %.2f and width %.2f\nArea: %.2f\nPerimeter: %.2f", r.Height, r.Width, r.Area(), r.Perimeter())
}

func DescribeShape(s Shape) {
	fmt.Println(s)
}
