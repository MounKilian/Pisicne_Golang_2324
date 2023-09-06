package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	for n1 := 48; n1 <= 55; n1++ {
		for n2 := 49; n2 <= 56; n2++ {
			if n1 >= n2 {
				n2 = n1 + 1
			}
			for n3 := 50; n3 <= 57; n3++ {
				if n2 >= n3 {
					n3 = n2 + 1
				}
				z01.PrintRune(rune(n1))
				z01.PrintRune(rune(n2))
				z01.PrintRune(rune(n3))
				z01.PrintRune(59)
				z01.PrintRune(' ')
			}
		}
	}
}
