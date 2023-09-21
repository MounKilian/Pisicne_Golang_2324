package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	result := true
	etat := 1
	if f(a[0], a[1]) < 0 {
		etat = 0
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
