package piscine

func FindNextPrime(nb int) int {
	if IsPrime(nb) == false {
		return FindNextPrime(nb + 1)
	} else {
		return nb
	}
}
