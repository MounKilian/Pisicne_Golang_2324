package piscine

func Join(strs []string, sep string) string {
	var final string
	for _, i := range strs {
		final = final + i + sep
	}
	return final
}
