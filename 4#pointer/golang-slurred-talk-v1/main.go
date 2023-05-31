package main

import "fmt"

func SlurredTalk(words *string) {
	wordsRune := []rune(*words)
	for i := 0 ; i < len(wordsRune) ; i++ {
		if string(wordsRune[i]) == "S" || string(wordsRune[i]) == "R" || string(wordsRune[i]) == "Z" {
			wordsRune[i] = 'L'
		} else if string(wordsRune[i]) == "s" || string(wordsRune[i]) == "r" || string(wordsRune[i]) == "z" {
			wordsRune[i] = 'l'
		}
	}
	*words = string(wordsRune)
}

func main() {
	// bisa dicoba untuk pengujian test case
	var words string = "Saya Steven, saya suka menggoreng telur dan suka hewan zebra"
	SlurredTalk(&words)
	fmt.Println(words)
}
