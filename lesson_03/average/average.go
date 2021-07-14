package average

import (
	"fmt"
	"reflect"
)

type ReturnCode int8

const (
	Success ReturnCode = iota
	Error
)

func isNumericKind(kind reflect.Kind) bool {
	switch kind {
	case reflect.Complex128, reflect.Complex64,
		reflect.Float32, reflect.Float64,
		reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return true

	default:
		return false
	}
}

// GetAverage takes any kind of slice or array
// for numeric types it calculates an average value of all items
// Returns 0, Error in case of unsupported data types
// Returns average value, Success in other cases
func GetAverage(inData interface{}) (float64, ReturnCode) {
	if inData == nil {
		fmt.Printf("Error! Not initialized container!\n")
		return -1, Error
	}

	switch reflect.TypeOf(inData).Kind() {
	case reflect.Slice, reflect.Ptr, reflect.Array:

		isNumeric := isNumericKind(reflect.TypeOf(inData).Elem().Kind())
		if !isNumeric {
			fmt.Printf("Error! Not supported type %s !\n", reflect.TypeOf(inData).Elem().Kind().String())
			return .0, Error
		}

		values := reflect.Indirect(reflect.ValueOf(inData))

		if values.Len() == 0 {
			return .0, Success
		}

		sum := .0

		lenght := values.Len()
		for i := 0; i < lenght; i++ {
			sum += values.Index(i).Convert(reflect.TypeOf(sum)).Float()
		}

		return sum / float64(lenght), Success
	}

	return .0, Error
}

func TestSlice() {
	fmt.Println("Slices Test")

	// Supported Type
	int_arr := [2]int{1, 2}
	avg_arr, err_arr := GetAverage(int_arr)
	fmt.Println(err_arr, avg_arr)

	// Unsupported Type
	str_arr := [2]string{"1", "2"}
	avg_str, err_str := GetAverage(str_arr)
	fmt.Println(err_str, avg_str)

	// Supported Type
	int_slice := []int32{1, 2}
	avg, err := GetAverage(int_slice)
	fmt.Println(err, avg)

	// Supported Type
	float_slice := []float64{1, 2}
	float_avg, float_err := GetAverage(float_slice)
	fmt.Println(float_err, float_avg)

	//unsupported type
	str_slice := []string{"1", "2"}
	s_avg, s_err := GetAverage(str_slice)
	fmt.Println(s_avg, s_err)
}
