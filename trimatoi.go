package piscine

func TrimAtoi(s string) int {
	i := 0
	for _, j := range s {
		if j >= '0' && j <= '9' {
			if j == '0' {
				i += 0
				i *= 10
			} else if j == '1' {
				i += 1
				i *= 10
			} else if j == '2' {
				i += 2
				i *= 10
			} else if j == '3' {
				i += 3
				i *= 10
			} else if j == '4' {
				i += 4
				i *= 10
			} else if j == '5' {
				i += 5
				i *= 10
			} else if j == '6' {
				i += 6
				i *= 10
			} else if j == '7' {
				i += 7
				i *= 10
			} else if j == '8' {
				i += 8
				i *= 10
			} else if j == '9' {
				i += 9
				i *= 10
			}
		}
	}
	result := i / 10
	if s[0] == '-' {
		return -result
	}
	aide := 0
	for _, j := range s {
		if j >= '0' && j <= '9' {
			aide++
		}
		if j == '-' {
			if aide == 0 {
				result = -result
			}
		}
	}
	return result
}
