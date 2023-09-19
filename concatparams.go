package piscine

func ConcatParams(args []string) string {
	result := ""
	for i := 0; i <= len(args)-2; i++ {
		result += args[i] + "\n"
	}
	result += args[len(args)-1]
	return result
}
