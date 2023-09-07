package piscine

func IsPrintable(s string) bool {
	ref := []rune(s)
	for _, i := range ref {
		if (i < 97 || i > 123) && (i < 65 || i > 91) {
			return false
		}
	}
	return true
}
