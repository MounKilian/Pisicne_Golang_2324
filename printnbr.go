package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbr(n int) {
	var tab []int
	j := 0

	if n == 0 {
		z01.PrintRune('0')
	}

	if n == -9223372036854775808 {
		n = n / 10
		tab = append(tab, 8)
	}

	if n < 0 {
		z01.PrintRune('-')
		n = -1 * n
	}

	for n != 0 {
		for i := 1; i <= n%10; i++ {
			j++
		}
		n = n - (n % 10)
		n = n / 10

		tab = append(tab, j)
		j = 0
	}

	for o := len(tab) - 1; o >= 0; o-- {
		z01.PrintRune(rune(tab[o] + 48))
	}
}
