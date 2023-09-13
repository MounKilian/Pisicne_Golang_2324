package piscine

func IterativePower(nb int, power int) int {
	result := nb
	if nb <= 0 {
		return 0
	} else if power <= 0 {
		return 0
	} else {
		for i := 1; i <= power-1; i++ {
			nb *= result
		}
		return nb
	}
}
