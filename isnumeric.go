package piscine

func IsNumeric(s string) bool {
	count := []rune(s)
	for _, i := range count {
		if i < 48 || i > 57 {
			return false
		}
	}
	return true
}
