package main

import (
	"fmt"
	"piscine"
)

func f(a, b int) int {
	if a > b {
		return 1
	} else if a == b {
		return 0
	} else {
		return -1
	}
}

func main() {
	a1 := []int{893348, 750835, 381939, 231675, -338395, -417710, -435701, -944974}
	a2 := []int{0, 0, 0, 0, 0, 0, 0}

	result1 := piscine.IsSorted(f, a1)
	result2 := piscine.IsSorted(f, a2)

	fmt.Println(result1)
	fmt.Println(result2)
}
