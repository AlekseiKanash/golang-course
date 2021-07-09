package main

import (
	"fmt"

	"github.com/AlekseiKanash/golang-course/lesson_02/fibonacci"
)

func main() {
	fmt.Printf("Fibonacci numbers are calculating...\n")
	defer fmt.Print("Done.\n")
	fibonacci.PrintFibonacciSequence(10)
}
