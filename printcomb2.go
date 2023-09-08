package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	for n1 := 48; n1 <= 57; n1++ {
		for n2 := 48; n2 <= 55; n2++ {
			for n3 := 48; n3 <= 57; n3++ {
				for n4 := 0; n4 <= 56; n4++ {
					if n1 <= n3 && n2 < n4 {
						z01.PrintRune(rune(n1))
						z01.PrintRune(rune(n2))
						z01.PrintRune(' ')
						z01.PrintRune(rune(n3))
						z01.PrintRune(rune(n4))
						z01.PrintRune(44)
						z01.PrintRune(' ')
					}
				}
			}
		}
	}
	z01.PrintRune((rune(57)))
	z01.PrintRune((rune(56)))
	z01.PrintRune(' ')
	z01.PrintRune((rune(57)))
	z01.PrintRune((rune(57)))
	z01.PrintRune('\n')
}
