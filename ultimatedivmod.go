package piscine

func UltimateDivMod(a *int, b *int) {
	c := 0
	c = *a
	*a = *a / *b
	*b = c % *b
}
