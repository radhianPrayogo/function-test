package main

import (
	"fmt"
	"sort"
	"strings"
)

func sortNumb(theNums ...float64) []float64 {
	sort.Float64s(theNums)
	result := strings.Join(theNums, ",")
	return result
}

func main() {
	fmt.Println(sortNumb(5, 4, 3, 7, 5, 8, 6, 9))
}
