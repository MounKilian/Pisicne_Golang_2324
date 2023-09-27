// package main

// import (
// 	"fmt"
// 	"os"
// )

// func BasicAtoi(s string) int {
// 	result := 0
// 	num := 0
// 	tab := []rune(s)
// 	for _, i := range tab {
// 		for j := '0'; j < i; j++ {
// 			num++
// 		}
// 		result = result*10 + num
// 		num = 0
// 	}
// 	return result
// }

// func Atoi(s string) int {
// 	result := BasicAtoi(s)
// 	aide := 0
// 	for _, i := range s {
// 		if i >= '0' && i <= '9' {
// 			aide++
// 		} else if i == '-' {
// 			if aide == 0 {
// 				result = -result
// 				aide++
// 			} else {
// 				return 0
// 			}
// 		} else if i == '+' {
// 			if aide == 0 {
// 				aide++
// 			} else {
// 				return 0
// 			}
// 		} else {
// 			return 0
// 		}
// 	}
// 	return result
// }

// func main() {
// 	if os.Args[1] == "-c" {
// 		arg := os.Args[3:]
// 		count := os.Args[2]
// 		help := Atoi(count)
// 		for m, i := range arg {
// 			data, err := os.ReadFile(i)
// 			if err != nil {
// 				fmt.Println(err)
// 			}
// 			if len(data) == 0 {
// 				fmt.Println("")
// 			} else if len(data) < help {
// 				fmt.Println("==>", arg[m], "<==")
// 				for n := range data {
// 					fmt.Print(string(data[n]))
// 				}
// 				os.Exit(1)
// 			} else {
// 				fmt.Println("==>", arg[m], "<==")
// 				for k := len(data) - help + 1; k <= len(data)-1; k++ {
// 					fmt.Print(string(data[k]))
// 				}
// 				fmt.Println()
// 			}
// 		}
// 	}
// }

package main

import (
	"fmt"
	"os"
)

func atoi(s string) int {
	ni := 0
	neg := false
	compt := 1
	valide := 0

	for i := len(s) - 1; i >= 0; i-- {
		if (rune(s[i]) >= 48 && rune(s[i]) <= 57) || rune(s[i]) == 45 || rune(s[i]) == 43 {
			if rune(s[i]) == rune(45) && i == 0 {
				neg = true
				valide++
			} else if rune(s[i]) == rune(43) && i == 0 {
				valide++
			} else if rune(s[i]) == rune(45) && i != 0 {
				return 0
			} else if rune(s[i]) == rune(43) && i != 0 {
				return 0
			}
			for c := 48; c < 58; c++ {
				if rune(s[i]) == rune(c) {
					ni += (c - 48) * compt
					compt = compt * 10
				}
			}
		} else {
			return 0
		}
	}
	if valide >= 2 {
		return 0
	}
	if neg {
		return ni * (-1)
	}
	return ni
}

func main() {
	if len(os.Args) <= 1 {
		return
	}
	for g := 3; g < len(os.Args); g++ {
		file, err := os.ReadFile(os.Args[g])
		if err == nil {
			if g == 3 {
				fmt.Printf("==> %s <==\n", os.Args[g])
				if len(file)-atoi(os.Args[2]) < 0 {
					fmt.Printf("%s", string(file[0:]))
				} else {
					fmt.Printf("%s", string(file[len(file)-atoi(os.Args[2]):]))
				}
			} else {
				fmt.Printf("\n")
				fmt.Printf("==> %s <==\n", os.Args[g])
				if len(file)-atoi(os.Args[2]) < 0 {
					fmt.Printf("%s", string(file[0:]))
				} else {
					fmt.Printf("%s", string(file[len(file)-atoi(os.Args[2]):]))
				}
				if g != len(os.Args)-1 {
					fmt.Printf("\n")
				}
			}
		} else {
			fmt.Printf("%v\n", err)
			defer func() { os.Exit(1) }()
		}
	}
}
