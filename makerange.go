package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		return []int(nil)
	}
	i := make([]int, max-1, min)
	return i
}
