package piscine

func IsAlpha(s string) bool {
	for i := range s {
		nonalpha := (s[i] < 65 || s[i] > 90) && (s[i] < 97 || s[i] > 122)
		nonnumeric := (s[i] < 48 || s[i] > 57)
		if nonalpha && nonnumeric {
			return false
		}
	}
	return true
}
