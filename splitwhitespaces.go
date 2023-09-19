package piscine

func SplitWhiteSpaces(s string) []string {
	result := []string{}
	aide := ""
	for j, i := range s {
		if string(i) == " " {
			if string(s[j-1]) == " " {
				result = append(result, "")
			}
			result = append(result, aide)
			aide = ""
		} else {
			aide += string(i)
		}
	}
	result = append(result, aide)
	return result
}
