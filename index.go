package piscine

func Index(s string, toFind string) int {
	if len(s) == 0 {
		return 0
	}
	plus := 0
	count := 0
	tab := []rune(toFind)
	tab2 := []rune(s)
	for _, i := range s {
		if i == tab[0] {
			break
		} else {
			count++
			plus++
		}
	}
	if count == len(s) {
		return -1
	}
	if count != 0 {
		for _, j := range tab {
			if tab2[plus] == j {
				plus++
			} else {
				return -1
			}
		}
	}
	return count
}
