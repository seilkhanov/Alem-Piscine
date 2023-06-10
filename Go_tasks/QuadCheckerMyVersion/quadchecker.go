package piscine

// package main

// import (
// 	"fmt"
// 	"io"
// 	"os"
// 	"os/exec"
// 	"sort"
// 	"strconv"
// 	"strings"
// )

// func main() {
// 	nameBytes, _ := io.ReadAll(os.Stdin)
// 	name := strings.TrimSpace(string(nameBytes))

// 	y := 0
// 	x := 0

// 	for _, r := range nameBytes {
// 		if r == '\n' {
// 			y++
// 		} else {
// 			x++
// 		}
// 	}

// 	if y == 0 || x == 0 {
// 		fmt.Println("Not a quad function")
// 		return
// 	}

// 	x /= y

// 	var matchingQuads []string
// 	quads := []string{"quadA", "quadB", "quadC", "quadD", "quadE"}

// 	for _, quad := range quads {
// 		if output, err := exec.Command("./"+quad, strconv.Itoa(x), strconv.Itoa(y)).CombinedOutput(); err == nil {
// 			if strings.TrimSpace(string(output)) == string(name) {
// 				matchingQuads = append(matchingQuads, "["+quad+"] ["+strconv.Itoa(x)+"] ["+strconv.Itoa(y)+"]")
// 			}
// 		}
// 	}

// 	if len(matchingQuads) == 0 {
// 		fmt.Println("Not a quad function")
// 		return
// 	}

// 	sort.Strings(matchingQuads)

// 	fmt.Println(strings.Join(matchingQuads, " || "))
// }
