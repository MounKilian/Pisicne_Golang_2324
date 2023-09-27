package main

import (
	"fmt"
	"os"
)

func BasicAtoi(s string) int {
	result := 0
	num := 0
	tab := []rune(s)
	for _, i := range tab {
		for j := '0'; j < i; j++ {
			num++
		}
		result = result*10 + num
		num = 0
	}
	return result
}

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

func main() {
	if os.Args[1] == "-c" {
		arg := os.Args[3:]
		count := os.Args[2]
		help := Atoi(count)
		for m, i := range arg {
			data, err := os.ReadFile(i)
			if err != nil {
				fmt.Println(err)
			}
			if len(data) == 0 {
				fmt.Print("")
			} else if len(data) <= help {
				fmt.Println("==>", arg[m], "<==")
				for n := range data {
					fmt.Print(string(data[n]))
				}
				fmt.Print("\n")
				fmt.Print("exit status 1")
			} else {
				fmt.Println("==>", arg[m], "<==")
				for k := len(data) - help + 1; k <= len(data)-1; k++ {
					fmt.Print(string(data[k]))
				}
				fmt.Print("\n")
			}
		}
	}
}
