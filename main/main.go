package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.IsPrintable("1MTRY0A#|89ji"))
	fmt.Println(piscine.IsPrintable("Hello\n"))
}
