package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.FindNextPrime(-10000))
	fmt.Println(piscine.FindNextPrime(1000000086))
	fmt.Println(piscine.FindNextPrime(1000000087))
	fmt.Println(piscine.FindNextPrime(1000000088))

}
