package piscine

func Join(strs []string, sep string) string {
	var final string
	for i := 0; i <= len(strs)-2; i++ {
		final = final + strs[i] + sep
	}
	final = final + strs[len(strs)-1]
	return final
}
