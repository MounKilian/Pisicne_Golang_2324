package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.Compare("Hello!", "Hello!"))
	fmt.Println(piscine.Compare("Salut!", "B"))
	fmt.Println(piscine.Compare("Ola!", "aaa"))
}
