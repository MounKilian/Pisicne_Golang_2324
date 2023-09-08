package piscine

func BasicJoin(elems []string) string {
	var final string
	for _, i := range elems {
		final = final + i
	}
	return final
}
