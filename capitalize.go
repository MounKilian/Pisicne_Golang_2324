package piscine

func Capitalize(s string) string {
	var result string
	tab := []rune(s)
	if tab[0] >= 'a' && tab[0] <= 'z' {
		result = result + string(tab[0]-32)
	} else {
		result = result + string(tab[0])
	}
	for i := 1; i <= len(tab)-1; i++ {
		if tab[i] >= 'a' && tab[i] <= 'z' {
			if (tab[i-1] >= 'A' && tab[i-1] <= 'Z') || (tab[i-1] >= 'a' && tab[i-1] <= 'z') || (tab[i-1] >= 48 && tab[i-1] <= 57) {
				result = result + string(tab[i])
			} else {
				result = result + string(rune(tab[i]-32))
			}
		} else if tab[i] >= 'A' && tab[i] <= 'Z' {
			if (tab[i-1] >= 'a' && tab[i-1] <= 'z') || (tab[i-1] >= 'A' && tab[i-1] <= 'Z') || (tab[i-1] >= 48 && tab[i-1] <= 57) {
				result = result + string(tab[i]+32)
			} else {
				result = result + string(tab[i])
			}
		} else {
			result = result + string(tab[i])
		}
	}
	return result
}