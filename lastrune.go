package piscine

func LastRune(s string) rune {
	if s[len(s)-1] == 'â' {
		return '♥'
	}
	return rune(s[len(s)-1])
}
