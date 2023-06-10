package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(a []string) {
	for _, word := range a {
		for _, letter := range word {
			z01.PrintRune(letter)
		}
		z01.PrintRune('\n')
	}
}

// func SplitWhiteSpaces(s string) []string {
// 	answer := []string{}
// 	inWord := false
// 	start := 0

// 	for i, letter := range s {
// 		seps := letter == ' ' || letter == '\t' || letter == '\n'
// 		if seps {
// 			if inWord {
// 				answer = append(answer, s[start:i])
// 				inWord = false
// 			}
// 		} else {
// 			if !inWord {
// 				start = i
// 				inWord = true
// 			}
// 		}
// 	}

// 	if inWord {
// 		answer = append(answer, s[start:])
// 	}

// 	return answer
// }
