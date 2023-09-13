package piscine

func FindNextPrime(nb int) int {
	n := nb + 1
	if IsPrime(nb) == false {
		return FindNextPrime(n)
	} else {
		return nb
	}
}
