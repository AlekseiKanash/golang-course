package maps

import (
	"fmt"
	"sort"
)

func sortData(data map[int]string) []string {
	out := []string{}
	keys := []int{}
	for k := range data {
		keys = append(keys, k)
	}

	sort.Ints(keys)

	for _, k := range keys {
		out = append(out, data[k])
	}
	return out
}

func printSorted(data map[int]string) {
	out := sortData(data)
	fmt.Println(out)
}

func TestMaps() {
	fmt.Println("Maps Test")

	cnt := map[int]string{2: "a", 0: "b", 1: "c"}
	printSorted(cnt)

	cnt = map[int]string{10: "aa", 0: "bb", 500: "cc"}
	printSorted(cnt)
}
