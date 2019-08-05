package main

import (
	"fmt"
	"sort"
)

func sortNumb(theNums ...float64) []float64 {
	sort.Float64s(theNums)
	return theNums
}

func main() {
	fmt.Println(sortNumb(5, 4, 3, 7, 5, 8, 6, 9))
}
