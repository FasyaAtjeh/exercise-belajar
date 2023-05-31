package main

import (
	"fmt"
	"strings"
)

func CountingLetter(text string) int {
	// unreadable letters = R, S, T, Z
	counter := 0
	for _, letters := range text {
		capital := strings.ToUpper(string(letters))
		if capital == "R" || capital == "S" || capital == "T" || capital == "Z" {
           counter++ 
        }
	}
	return counter
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(CountingLetter("Semangat"))
}
