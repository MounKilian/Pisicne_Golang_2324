package piscine

func Capitalize(s string) string {
	var result string
	tab := []rune(s)
	result = result + string(tab[0])
	for i := 1; i <= len(tab)-1; i++ {
		if tab[i] >= 'a' && tab[i] <= 'z' {
			if (tab[i-1] >= 'A' && tab[i-1] <= 'Z') || (tab[i-1] >= 'a' && tab[i-1] <= 'z') {
				result = result + string(tab[i])
			} else if tab[i-1] == ' ' || tab[i-1] == '+' || tab[i-1] == '!' || tab[i-1] == '?' || (tab[i-1] >= 48 && tab[i-1] <= 57) || tab[i-1] == '_' || tab[i-1] == '@' || tab[i-1] == '%' {
				result = result + string(rune(tab[i]-32))
			}
		} else {
			result = result + string(tab[i])
		}
	}
	return result
}
