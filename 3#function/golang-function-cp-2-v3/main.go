package main

import (
	"fmt"
	"strings"
)

func CountVowelConsonant(str string) (int, int, bool) {
	vocalCount := 0
	consonantCount := 0
	for _, c := range str {
		cc := strings.ToLower(string(c))
		if cc == "a" || cc == "i" || cc == "u" || cc == "e" || cc == "o" {
			vocalCount += 1
		} else if cc != " " && cc != ","{
			consonantCount +=1
		}
	}
	fmt.Println(vocalCount, consonantCount)
	isValid := false
	if vocalCount == 0 || consonantCount == 0 {
		isValid = true
	} else {
		isValid = false
	}
	return vocalCount, consonantCount, isValid // TODO: replace this
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(CountVowelConsonant("Hidup Itu Indah"))
}
