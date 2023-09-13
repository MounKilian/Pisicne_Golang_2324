package piscine

func FindNextPrime(nb int) int {
	n := nb + 1
	if !IsPrime(nb) {
		return FindNextPrime(n)
	} else {
		return nb
	}
}
