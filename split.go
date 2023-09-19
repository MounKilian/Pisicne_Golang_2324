package piscine

func Split(s, sep string) []string {
	result := []string{}
	aide := ""
	for i := 0; i <= len(s)-2; i++ {
		if s[i:i+len(sep)] == sep {
			result = append(result, aide)
			aide = ""
			i = i + len(sep) - 1
		} else {
			aide += string(s[i])
		}
	}
	result = append(result, aide)
	return result
}
