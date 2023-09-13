package piscine

func FindNextPrime(nb int) int {
	if IsPrime(nb) == true {
		return nb
	} else {
		return FindNextPrime(nb + 1)
	}
}
