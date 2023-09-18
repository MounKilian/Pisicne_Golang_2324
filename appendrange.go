package piscine

func AppendRange(min, max int) []int {
	result := []int{}
	if max <= min {
		return []int(nil)
	}
	for i := min; i < max; i++ {
		result = append(result, i)
	}
	return result
}
