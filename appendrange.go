package piscine

func AppendRange(min, max int) []int {
	result := []int{}
	if max <= min || max == 0 || min == 0 {
		return result
	}
	for i := min; i <= max; i++ {
		result = append(result, i)
	}
	return result
}
