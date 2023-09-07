package piscine

func IsPrintable(s string) bool {
	ref := []rune(s)
	for _, i := range ref {
		if i < 20 || i > 127 {
			return false
		}
	}
	return true
}
