package piscine

func Capitalize(s string) string {
	newstring := []byte{}
	previousIsLetter := false

	for i := range s {
		ualpha := s[i] > 64 && s[i] < 91
		lalpha := s[i] > 96 && s[i] < 123
		numeric := s[i] >= 48 && s[i] <= 57
		if ualpha || lalpha || numeric {
			if !previousIsLetter {
				if ualpha || numeric {
					newstring = append(newstring, s[i])
				} else {
					newstring = append(newstring, s[i]-32)
				}
				previousIsLetter = true
			} else {
				if ualpha {
					newstring = append(newstring, s[i]+32)
				} else {
					newstring = append(newstring, s[i])
				}
			}
		} else {
			newstring = append(newstring, s[i])
			previousIsLetter = false
		}
	}
	return string(newstring)
}
