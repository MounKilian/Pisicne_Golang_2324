package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args[1:]
	min := 0
	var indice string
	for j := 0; j <= len(arg)-1; j++ {
		min = j
		for m := 0; m <= len(arg)-1; m++ {
			if arg[m] > arg[min] {
				min = j
				indice = arg[j]
				arg[j] = arg[m]
				arg[m] = indice
			}
		}
	}
	for _, i := range arg {
		for _, k := range i {
			z01.PrintRune(k)
		}
		z01.PrintRune('\n')
	}
}
