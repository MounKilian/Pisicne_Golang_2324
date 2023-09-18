package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args[1:]
	for k := len(arg) - 1; k >= 0; k-- {
		for _, i := range arg[k] {
			z01.PrintRune(i)
		}
		z01.PrintRune('\n')
	}
}
