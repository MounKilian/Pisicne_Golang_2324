package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	result := true
	etat := 1
	count := 0
	if f(a[0], a[1]) < 0 {
		etat = 0
	}
	for j := 0; j <= len(a)-2; j++ {
		if f(a[j], a[j+1]) == 0 {
			count++
		}
	}
	if count == len(a)-1 {
		return true
	}
	for i := 1; i <= len(a)-2; i++ {
		if etat == 0 {
			if f(a[i], a[i+1]) < 0 {
				result = true
			} else {
				return false
			}
		} else if etat == 1 {
			if f(a[i], a[i+1]) > 0 {
				result = true
			} else {
				return false
			}
		}
	}
	return result
}
