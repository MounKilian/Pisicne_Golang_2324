package piscine

func ToLower(s string) string {
	var fin []rune
	tab := []rune(s)
	for _, i := range tab {
		if i >= 65 && i <= 90 {
			fin = append(fin, i+32)
		} else {
			fin = append(fin, i)
		}
	}
	return string(fin)
}
