package piscine

func StrLen(s string) int {
	count := 0
	for i := 0; i <= len(s)-1; i++ {
		if s[i] != ' ' {
			count++
		}
	}
	return count
}
