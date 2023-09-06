package main

import (
	"piscine"

	"github.com/01-edu/z01"
)

func IsNegative(nb int) {
	if nb < 0 {
		z01.PrintRune(54)
	} else {
		z01.PrintRune(70)
	}
}

func main() {
	piscine.IsNegative(1)
	piscine.IsNegative(0)
	piscine.IsNegative(-1)
}
