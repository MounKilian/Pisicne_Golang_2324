package piscine

import (
	"fmt"

	"github.com/01-edu/z01"
)

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune('0')
	} else {
		j := 0
		result := []int{}
		for n != 0 {
			j = n % 10
			result = append(result, j)
			n = (n / 10) - (j / 10)
			j = 0
		}
		i_mini := 0
		indice := 0
		for i := 0; i <= len(result)-1; i++ {
			i_mini = i
			for j := i + 1; j <= len(result)-1; j++ {
				if result[j] < result[i_mini] {
					i_mini = j
					indice = result[i]
					result[i] = result[i_mini]
					result[i_mini] = indice
				}
			}
		}
		for _, k := range result {
			fmt.Print(k)
		}
	}
}
