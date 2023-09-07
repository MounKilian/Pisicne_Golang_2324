package piscine

func IsUpper(s string) bool {
	count := []rune(s)
	for _, i := range count {
		if i < 65 || i > 90 {
			return false
		}
	}
	return true
}
