package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	arg := os.Args[1:]
	if len(arg) == 1 {
		data, err := ioutil.ReadFile(arg[0])
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(data))
	} else if len(arg) > 1 {
		fmt.Print("Too many arguments")
		fmt.Print("\n")
	} else {
		fmt.Print("File name missing")
		fmt.Print("\n")
	}
}
