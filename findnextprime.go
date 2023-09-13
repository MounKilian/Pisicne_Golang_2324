package piscine

func FindNextPrime(nb int) int {
	if nb < 2 {
		return 2
	}
	if IsPrime(nb) == true {
		return nb
	} else {
		return FindNextPrime(nb + 1)
	}
}
