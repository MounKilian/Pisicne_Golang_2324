package piscine

func SplitWhiteSpaces(s string) []string {
	result := []string{}
	word := ""
	for _, i := range s {
		if i == ' ' {
			if word == "" {
				continue
			} else {
				result = append(result, word)
				word = ""
			}
		} else {
			word += string(i)
		}
	}
	result = append(result, word)
	return result
}
