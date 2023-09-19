package piscine

func Split(s, sep string) []string {
	result := []string{}
	aide := ""
	for i := 0; i <= len(s)-len(sep); i++ {
		if s[i:i+len(sep)] == sep {
			result = append(result, aide)
			aide = ""
			i = i + len(sep) - 1
		} else {
			aide += string(s[i])
		}
	}
	for j := len(s) - len(sep) + 1; j <= len(s)-1; j++ {
		aide += string(s[j])
	}
	result = append(result, aide)
	return result
}
