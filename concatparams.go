package piscine

func ConcatParams(args []string) string {
	result := ""
	for _, i := range args {
		if i == args[len(args)-1] {
			result += i
		} else {
			result += i + string('\n')
		}
	}
	return result
}
