package piscine

func BasicAtoi(s string) int {
	result := 0
	num := 0
	tab := []rune(s)
	for _, i := range tab {
		for j := '0'; j < i; j++ {
			num++
		}
		result = (result + num) * 10
		num = 0
	}
	return result / 10
}
