package piscine

func Atoi(s string) int {
	result := BasicAtoi(s)
	aide := 0
	for _, i := range s {
		if i >= '0' && i <= '9' {
			aide++
		} else if i == '-' {
			if aide == 0 {
				result = -result
				aide++
			} else {
				return 0
			}
		} else if i == '+' {
			if aide == 0 {
				aide++
			} else {
				return 0
			}
		} else {
			return 0
		}
	}
	return result
}
