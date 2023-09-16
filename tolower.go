package piscine

func ToLower(s string) string {
	var result []rune
	tab := []rune(s)
	for _, i := range tab {
		if i >= 'A' && i <= 'Z' {
			result = append(result, i+32)
		} else {
			result = append(result, i)
		}
	}
	return string(result)
}
