package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args[1:]
	for _, k := range arg {
		for _, i := range k {
			z01.PrintRune(i)
		}
		z01.PrintRune('\n')
	}
}
