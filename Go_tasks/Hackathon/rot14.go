package piscine

func Rot14(s string) string {
	newsentence := ""
	for _, arg := range s {
		isUalpha := (arg > 64 && arg < 91)
		isLalpha := (arg > 96 && arg < 123)
		if isUalpha || isLalpha {
			if isUalpha {
				if arg+14 < 91 {
					newsentence += string(arg + 14)
				} else {
					newsentence += string(arg - 12)
				}
			} else if isLalpha {
				if arg+14 < 123 {
					newsentence += string(arg + 14)
				} else {
					newsentence += string(arg - 12)
				}
			}
		} else {
			newsentence += string(arg)
		}
	}
	return newsentence
}
