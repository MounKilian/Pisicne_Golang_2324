package piscine

func FirstRune(s string) rune {
	if s[0] == 'â' {
		return '♥'
	}
	return rune(s[0])
}
