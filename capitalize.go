package piscine

func Capitalize(s string) string {
	var result []rune
	tab := []rune(s)
	if tab[0] >= 'a' && tab[0] <= 'z' {
		result = append(result, tab[0]-32)
	} else {
		result = append(result, tab[0])
	}
	for i := 1; i <= len(tab)-1; i++ {
		if tab[i] >= 'a' && tab[i] <= 'z' {
			if (tab[i-1] >= 'a' && tab[i-1] <= 'z') || (tab[i-1] >= '0' && tab[i-1] <= '9') || (tab[i-1] >= 'A' && tab[i-1] <= 'Z') {
				result = append(result, tab[i])
			} else {
				result = append(result, tab[i]-32)
			}
		} else if tab[i] >= 'A' && tab[i] <= 'Z' {
			if (tab[i-1] >= 'a' && tab[i-1] <= 'z') || (tab[i-1] >= '0' && tab[i-1] <= '9') || (tab[i-1] >= 'A' && tab[i-1] <= 'Z') {
				result = append(result, tab[i]+32)
			} else {
				result = append(result, tab[i])
			}
		} else {
			result = append(result, tab[i])
		}
	}
	return string(result)
}
