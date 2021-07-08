package main

import (
	"fmt"

	"github.com/AlekseiKanash/golang-course/lesson_02/fibonacci"
)

func main() {
	defer fmt.Print("Done.\n")
	fmt.Printf("Fibonacci numbers are calculating...\n")
	fibonacci.PrintFibonacciSequence(10)
}
