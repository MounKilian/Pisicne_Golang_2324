package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	result := true
	for i := 0; i <= len(a)-2; i++ {
		if f(a[i], a[i+1]) < 0 {
			result = true
		} else {
			return false
		}
	}
	return result
}
