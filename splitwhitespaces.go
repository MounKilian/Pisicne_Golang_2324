package piscine

func SplitWhiteSpaces(s string) []string {
	result := []string{}
	aide := ""
	for j, i := range s {
		if string(i) == " " {
			if string(s[j-1]) == " " {
				result = append(result, "")
			} else {
				result = append(result, aide)
			}
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
