package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args[3]
	for _, i := range arg {
		z01.PrintRune(i)
	}
	z01.PrintRune('\n')
}
