package piscine

func Capitalize(s string) string {
	var result string
	tab := []rune(s)
	result = string(tab[0])
	for i := 1; i <= len(tab)-1; i++ {
		if tab[i] >= 97 && tab[i] <= 122 {
			if tab[i-1] <= 97 || tab[i-1] >= 122 {
				result = result + string(tab[i])
			}
		} else {
			result = result + string(tab[i]-32)
		}
	}
	return result
}
