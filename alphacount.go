package piscine

func AlphaCount(s string) int {
	compt := 0
	ref := []rune(s)
	for _, i := range ref {
		if i >= 97 && i < 123 {
			compt++
		}
		if i >= 65 && i < 91 {
			compt++
		}
	}
	return compt
}
