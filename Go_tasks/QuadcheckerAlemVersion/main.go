package piscine

// package main

// import (
// 	"io"
// 	"os"

// 	"github.com/01-edu/z01"
// )

// func main() {
// 	nameBytes, _ := io.ReadAll(os.Stdin)

// 	y := 0
// 	x := 0

// 	for _, r := range nameBytes {
// 		if r == '\n' {
// 			y++
// 		} else {
// 			x++
// 		}
// 	}

// 	if y <= 0 || x <= 0 {
// 		printError()
// 		return
// 	}

// 	x /= y

// 	str := []string{}

// 	if quadA(x, y) == string(nameBytes) {
// 		str = append(str, "[quadA] ["+itoa(x)+"] ["+itoa(y)+"]")
// 	}
// 	if quadB(x, y) == string(nameBytes) {
// 		str = append(str, "[quadB] ["+itoa(x)+"] ["+itoa(y)+"]")
// 	}
// 	if quadC(x, y) == string(nameBytes) {
// 		str = append(str, "[quadC] ["+itoa(x)+"] ["+itoa(y)+"]")
// 	}
// 	if quadD(x, y) == string(nameBytes) {
// 		str = append(str, "[quadD] ["+itoa(x)+"] ["+itoa(y)+"]")
// 	}
// 	if quadE(x, y) == string(nameBytes) {
// 		str = append(str, "[quadE] ["+itoa(x)+"] ["+itoa(y)+"]")
// 	}
// 	if len(str) == 0 {
// 		printError()
// 		return
// 	}
// 	for i := 0; i <= len(str)-1; i++ {
// 		for _, args := range str[i] {
// 			z01.PrintRune(args)
// 		}
// 		if i != len(str)-1 {
// 			z01.PrintRune(' ')
// 			z01.PrintRune('|')
// 			z01.PrintRune('|')
// 			z01.PrintRune(' ')
// 		} else {
// 			z01.PrintRune('\n')
// 		}
// 	}
// }

// func itoa(n int) string {
// 	result := ""
// 	for n > 0 {
// 		result = string(n%10+'0') + result
// 		n /= 10
// 	}
// 	return result
// }

// func printError() {
// 	err := "Not a quad function"
// 	for _, arg := range err {
// 		z01.PrintRune(arg)
// 	}
// 	z01.PrintRune('\n')
// }

// func quadA(x, y int) string {
// 	result := ""

// 	if x > 0 && y > 0 {
// 		for i := 1; i <= y; i++ {
// 			if i == 1 || i == y {
// 				for j := 1; j <= x; j++ {
// 					if j == 1 || j == x {
// 						result += "o"
// 					} else {
// 						result += "-"
// 					}
// 				}
// 				result += "\n"
// 			} else {
// 				for j := 1; j <= x; j++ {
// 					if j == x || j == 1 {
// 						result += "|"
// 					} else {
// 						result += " "
// 					}
// 				}
// 				result += "\n"
// 			}
// 		}
// 	}
// 	return result
// }

// func quadB(x, y int) string {
// 	result := ""

// 	if x > 0 && y > 0 {
// 		for i := 1; i <= y; i++ {
// 			if i == 1 {
// 				for j := 1; j <= x; j++ {
// 					if j == 1 {
// 						result += "/"
// 					} else if j == x {
// 						result += "\\"
// 					} else {
// 						result += "*"
// 					}
// 				}
// 				result += "\n"
// 			} else if i == y {
// 				for j := 1; j <= x; j++ {
// 					if j == 1 {
// 						result += "\\"
// 					} else if j == x {
// 						result += "/"
// 					} else {
// 						result += "*"
// 					}
// 				}
// 				result += "\n"
// 			} else {
// 				for j := 1; j <= x; j++ {
// 					if j == x || j == 1 {
// 						result += "*"
// 					} else {
// 						result += " "
// 					}
// 				}
// 				result += "\n"
// 			}
// 		}
// 	}
// 	return result
// }

// func quadC(x, y int) string {
// 	result := ""

// 	if x > 0 && y > 0 {
// 		for i := 1; i <= y; i++ {
// 			if i == 1 {
// 				for j := 1; j <= x; j++ {
// 					if j == 1 || j == x {
// 						result += "A"
// 					} else {
// 						result += "B"
// 					}
// 				}
// 				result += "\n"
// 			} else if i == y {
// 				for j := 1; j <= x; j++ {
// 					if j == 1 || j == x {
// 						result += "C"
// 					} else {
// 						result += "B"
// 					}
// 				}
// 				result += "\n"
// 			} else {
// 				for j := 1; j <= x; j++ {
// 					if j == x || j == 1 {
// 						result += "B"
// 					} else {
// 						result += " "
// 					}
// 				}
// 				result += "\n"
// 			}
// 		}
// 	}
// 	return result
// }

// func quadD(x, y int) string {
// 	result := ""

// 	if x > 0 && y > 0 {
// 		for i := 1; i <= y; i++ {
// 			if i == 1 {
// 				for j := 1; j <= x; j++ {
// 					if j == 1 {
// 						result += "A"
// 					} else if j == x {
// 						result += "C"
// 					} else {
// 						result += "B"
// 					}
// 				}
// 				result += "\n"
// 			} else if i == y {
// 				for j := 1; j <= x; j++ {
// 					if j == 1 {
// 						result += "A"
// 					} else if j == x {
// 						result += "C"
// 					} else {
// 						result += "B"
// 					}
// 				}
// 				result += "\n"
// 			} else {
// 				for j := 1; j <= x; j++ {
// 					if j == x || j == 1 {
// 						result += "B"
// 					} else {
// 						result += " "
// 					}
// 				}
// 				result += "\n"
// 			}
// 		}
// 	}
// 	return result
// }

// func quadE(x, y int) string {
// 	result := ""

// 	if x > 0 && y > 0 {
// 		for i := 1; i <= y; i++ {
// 			if i == 1 {
// 				for j := 1; j <= x; j++ {
// 					if j == 1 {
// 						result += "A"
// 					} else if j == x {
// 						result += "C"
// 					} else {
// 						result += "B"
// 					}
// 				}
// 				result += "\n"
// 			} else if i == y {
// 				for j := 1; j <= x; j++ {
// 					if j == 1 {
// 						result += "C"
// 					} else if j == x {
// 						result += "A"
// 					} else {
// 						result += "B"
// 					}
// 				}
// 				result += "\n"
// 			} else {
// 				for j := 1; j <= x; j++ {
// 					if j == x || j == 1 {
// 						result += "B"
// 					} else {
// 						result += " "
// 					}
// 				}
// 				result += "\n"
// 			}
// 		}
// 	}
// 	return result
// }
