package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args[1:]
	j := 0
	for _, k := range arg {
		k = k
		for _, i := range arg[j] {
			z01.PrintRune(i)
		}
		z01.PrintRune('\n')
		j++
	}
}
