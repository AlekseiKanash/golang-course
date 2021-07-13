package strings

import "fmt"

func Max(inStrings []string) string {
	if 0 == len(inStrings) || nil == inStrings {
		fmt.Printf("Error, Empty slice.\n")
		return ""
	}
	maxIndex := -1
	maxLen := -1
	for index, str := range inStrings {
		strLen := len(str)
		if strLen > maxLen {
			maxLen = strLen
			maxIndex = index
		}
	}

	return inStrings[maxIndex]
}

func Reverse(inStrings []string) []string {
	var out []string
	for i := len(inStrings) - 1; i >= 0; i-- {
		out = append(out, inStrings[i])
	}

	return out
}

func TestStrings() {
	fmt.Println("Strings Test")
	// valid
	longestString := Max([]string{"one", "two", "three"})
	fmt.Println(longestString)

	// valid
	longestString = Max([]string{"one", "two"})
	fmt.Println(longestString)

	// invalid
	longestString = Max([]string{})
	fmt.Println(longestString)

	inverted := Reverse([]string{"one", "two"})
	fmt.Println(inverted)
}
