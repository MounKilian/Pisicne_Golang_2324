package main

import (
	"fmt"
	"os"
	"piscine"
)

func main() {
	if os.Args[1] == "-c" {
		arg := os.Args[3:]
		count := os.Args[2]
		help := piscine.Atoi(count)
		for m, i := range arg {
			data, err := os.ReadFile(i)
			if err != nil {
				fmt.Println(err)
			}
			if len(data) == 0 {
				fmt.Print("")
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
