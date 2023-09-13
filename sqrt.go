package piscine

func Sqrt(nb int) int {
	result := 0
	for i := nb; i >= 0; i-- {
		if i*i == nb {
			result = i
		}
	}
	return result
}
