package piscine

func TrimAtoi(s string) int {
	if len(s) == 0 {
		return 0
	}
	var result int
	for _, i := range s {
		if i >= '0' && i <= '9' {
			if i == '1' {
				result += 1
				result *= 10
			} else if i == '2' {
				result += 2
				result *= 10
			} else if i == '3' {
				result += 3
				result *= 10
			} else if i == '4' {
				result += 4
				result *= 10
			} else if i == '5' {
				result += 5
				result *= 10
			} else if i == '6' {
				result += 6
				result *= 10
			} else if i == '7' {
				result += 7
				result *= 10
			} else if i == '8' {
				result += 8
				result *= 10
			} else if i == '9' {
				result += 9
				result *= 10
			} else if i == '0' {
				result *= 10
			}
		}
	}
	result /= 10
	aide := 0
	for _, j := range s {
		if j >= '0' && j <= '9' {
			aide++
		} else if j == '-' {
			if aide == 0 {
				result = -result
			}
		}
	}
	return result
}
