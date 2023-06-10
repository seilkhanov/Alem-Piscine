package piscine

func FindNextPrime(nb int) int {
	if nb <= 2 {
		return 2
	}

	for i := nb; ; i++ {
		check_prime := true

		for j := 2; j*j <= i; j++ {
			if i%j == 0 {
				check_prime = false
				break
			}
		}

		if check_prime {
			return i
		}
	}
}
