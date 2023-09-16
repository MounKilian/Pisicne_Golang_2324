package piscine

func Join(strs []string, sep string) string {
	var result string
	for i := 0; i <= len(strs)-2; i++ {
		result += strs[i] + sep
	}
	return result + strs[len(strs)-1]
}
