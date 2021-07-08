package fibonacci

import "fmt"

// getFibonacciSequence calculates a sequence of Fibonacchi
// numbers.
// [in] uint8 sequence_len - how many numbers in the sequence
// [ret] []uint64 - array of results
func getFibonacciSequence(sequence_len uint8) []uint64 {
	ret_array := []uint64{0, 1}

	var prev_number uint64 = 1
	var current_number uint64 = 1
	var sum uint64 = 1

	var i uint8 = 2
	for i < sequence_len {
		sum = current_number + prev_number
		prev_number = current_number
		current_number = sum
		ret_array = append(ret_array, current_number)
		i++
	}

	return ret_array
}

// PrintFibonacciSequence prints a sequence of Fibonacchi numbers
// [in] the sequence lenght
func PrintFibonacciSequence(sequence_len uint8) {
	fibonacchi_string := getFibonacciSequence(sequence_len)
	fmt.Printf("The 1st %d numbers are: %d\n", sequence_len, fibonacchi_string)
}
