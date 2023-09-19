package piscine

func SplitWhiteSpaces(s string) []string {
	result := []string{}
	aide := ""
	for _, i := range s {
		if string(i) == " " {
			result = append(result, aide)
			aide = ""
		} else {
			aide += string(i)
		}
	}
	result = append(result, aide)
	if string(s[len(s)-1]) == " " {
		result = append(result, "")
	}
	return result
}
