package piscine

func AppendRange(min, max int) []int {
	result := []int{}
	if max <= min {
		return result
	}
	for i := min; i <= max; i++ {
		result = append(result, i)
	}
	return result
}
