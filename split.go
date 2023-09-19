package piscine

func Split(s, sep string) []string {
	result := []string{}
	aide := ""
	for i := 0; i <= len(s)-2; i++ {
		if s[i:i+len(sep)] == sep {
			result = append(result, aide)
			aide = ""
			i = i + len(sep) - 1
			if i >= len(s)-2 {
				i = len(s) - 2
			}
		} else {
			aide += string(s[i])
		}
	}
	result = append(result, aide+string(s[len(s)-1]))
	return result
}
