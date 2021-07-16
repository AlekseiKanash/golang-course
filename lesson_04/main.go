package main

import (
	"github.com/AlekseiKanash/golang-course/lesson_04/shapes"
)

func main() {
	// choose your own dimensions
	c := shapes.Circle{Radius: 8}
	r := shapes.Rectangle{
		Height: 9,
		Width:  3,
	}

	shapes.DescribeShape(c)
	shapes.DescribeShape(r)
}
