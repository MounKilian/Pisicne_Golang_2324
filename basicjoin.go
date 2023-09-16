package piscine

func BasicJoin(elems []string) string {
	var result string
	for _, i := range elems {
		result += i
	}
	return result
}
