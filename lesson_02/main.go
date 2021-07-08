package main

import (
	"fmt"

	"github.com/AlekseiKanash/golang-course/lesson_02/fibonacchi"
)

func main() {
	defer fmt.Print("Done.\n")
	fmt.Printf("Fibonacchi numbers are calculating...\n")
	fibonacchi.PrintFibonacciSequence(10)
}
