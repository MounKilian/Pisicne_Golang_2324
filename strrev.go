package piscine

func StrRev(s string) string {
	tab := []rune(s)
	var result []rune
	for i := len(s) - 1; i >= 0; i-- {
		result = append(result, rune(tab[i]))
	}
	return string(result)
}
