package shapes

import (
	"fmt"
	"math"
	"testing"
)

type testResult struct {
	Area      float64
	Perimeter float64
	Str       string
}

func (r *testResult) String() string {
	return fmt.Sprintf("Area: %.2f Perimeter: %.2f", r.Area, r.Perimeter)
}

const (
	floatDelta float64 = 0.0001
)

// You should not compare floats using just ==.
// You only compare them with some accuracy
func fNearlyEqual(a, b float64) bool {
	return math.Abs(a-b) <= floatDelta
}

func testImpl(t *testing.T, shape Shape, epectedResult testResult) {
	testname := fmt.Sprintf("%+v, %+v", shape, epectedResult)
	t.Run(testname, func(t *testing.T) {
		isExpectedArea := fNearlyEqual(shape.Area(), epectedResult.Area)
		isExpectedPerimeter := fNearlyEqual(shape.Perimeter(), epectedResult.Perimeter)
		objectStr := fmt.Sprint(shape)
		isExpectedString := objectStr == epectedResult.Str
		if !isExpectedArea || !isExpectedPerimeter || !isExpectedString {
			t.Errorf("Unexpected result: Object %+v Area: %.2f Perimeter %.2f does not corresponds to %+v\n", shape, shape.Area(), shape.Perimeter(), epectedResult)
		} else {

		}
	})
}

func TestCircles(t *testing.T) {
	shapes := map[Shape]testResult{
		Circle{Radius: -1}: {Area: 0, Perimeter: 0, Str: "Circle: radius -1.00"},
		Circle{Radius: 0}:  {Area: 0, Perimeter: 0, Str: "Circle: radius 0.00"},
		Circle{Radius: 1}:  {Area: math.Pi * math.Pow(1, 2), Perimeter: 2 * math.Pi * 1, Str: "Circle: radius 1.00"},
	}

	for testValue, expectedResult := range shapes {
		testImpl(t, testValue, expectedResult)
	}
}

func TestRects(t *testing.T) {
	shapes := map[Shape]testResult{
		Rectangle{Height: -1, Width: -1}: {Area: 0, Perimeter: 0, Str: "Rectangle with height -1.00 and width -1.00"},
		Rectangle{Height: -1, Width: 0}:  {Area: 0, Perimeter: 0, Str: "Rectangle with height -1.00 and width 0.00"},
		Rectangle{Height: 0, Width: -1}:  {Area: 0, Perimeter: 0, Str: "Rectangle with height 0.00 and width -1.00"},
		Rectangle{Height: 0, Width: 1}:   {Area: 0, Perimeter: 2, Str: "Rectangle with height 0.00 and width 1.00"},
		Rectangle{Height: 1, Width: 0}:   {Area: 0, Perimeter: 2, Str: "Rectangle with height 1.00 and width 0.00"},
	}

	for testValue, expectedResult := range shapes {
		testImpl(t, testValue, expectedResult)
	}
}
