package main

import (
	"fmt"
	"strings"
	"unicode"
)

func ReverseWord(str string) string {
	var result = ""
	var word = ""

	for i := 0; i < len(str); i++ {
		if str[i] != ' ' {
			word += string(str[i])
		}
	
	if str[i] == ' ' || i == len(str)-1 {
		var reverse = ""

		for j := len(word) -1; j>=0; j-- {
			reverse += string(word[j])
        }
		if unicode.IsUpper(rune(word[0])) == true {
			reverse = strings.ToUpper(string(reverse[0])) + reverse[1:]
		}
		if unicode.IsLower(rune(word[len(word)-1])) == true {
			reverse = reverse[:len(reverse)-1] + strings.ToLower(string(reverse[len(reverse)-1])) 
        }
		result += reverse + " "
		word = ""
	}
}
	return result [:len(result)-1] // TODO: replace this

}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(ReverseWord("Aku Sayang Ibu"))
	fmt.Println(ReverseWord("A bird fly to the Sky"))
}
