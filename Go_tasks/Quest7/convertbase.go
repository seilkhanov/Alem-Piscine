package piscine

func ConvertBase(nbr, baseFrom, baseTo string) string {
	initialbase := len(baseFrom)
	finalbase := len(baseTo)
	result := 0
	newnbr := []byte{}
	output := ""

	for i := 0; i < len(nbr); i++ {
		digitValue := Index(baseFrom, string(nbr[i]))
		result += digitValue * Power(initialbase, len(nbr)-1-i)
	}

	for result > 0 {
		newnbr = append(newnbr, byte(result%finalbase))
		result = result / finalbase
	}

	for i, j := 0, len(newnbr)-1; i < j; i, j = i+1, j-1 {
		newnbr[i], newnbr[j] = newnbr[j], newnbr[i]
	}

	for _, digit := range newnbr {
		output += string(baseTo[digit])
	}
	return output
}

func Power(nb int, power int) int {
	result := 1
	for power > 0 {
		result *= nb
		power--
	}
	return result
}

func Index(s string, toFind string) int {
	for i := 0; i <= (len(s) - len(toFind)); i++ {
		if s[i:i+len(toFind)] == toFind {
			return i
		}
	}
	return -1
}
