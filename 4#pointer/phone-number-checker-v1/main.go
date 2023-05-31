package main

import (
	"fmt"
	"strings"
)

func PhoneNumberChecker(number string, result *string){
	if number [0:2] == "08" {
		number = "62" + number[1:]
	}
	if number[0:3] != "628"{
		*result = "invalid"
		return
	}
	if len(number) < 11 {
		*result = "invalid"
		return
	}
	if strings.HasPrefix(number, "62811") || strings.HasPrefix(number, "62812") || strings.HasPrefix(number, "62813") || strings.HasPrefix(number, "62814") || strings.HasPrefix(number, "62815") {
		*result = "Telkomsel"
		return
	} 
	if strings.HasPrefix(number, "62816") || strings.HasPrefix(number, "62817") || strings.HasPrefix(number, "62818") || strings.HasPrefix(number, "62819") {
		*result = "Indosat"
		return
	} 
	if strings.HasPrefix(number, "62821") || strings.HasPrefix(number, "62822") || strings.HasPrefix(number, "62823") {
		*result = "XL"
		return
	} 
	if strings.HasPrefix(number, "62827") || strings.HasPrefix(number, "62828") || strings.HasPrefix(number, "62829") {
		*result = "Tri"
		return
	} 
	if strings.HasPrefix(number, "62852") || strings.HasPrefix(number, "62853") {
		*result = "AS"
		return
	} 
	if strings.HasPrefix(number, "62881") || strings.HasPrefix(number, "62882") || strings.HasPrefix(number, "62883") ||strings.HasPrefix(number, "62884") || strings.HasPrefix(number, "62885") || strings.HasPrefix(number, "62886") || strings.HasPrefix(number, "62887") || strings.HasPrefix(number, "62888") {
		*result = "Smartfren"
		return 
	}
	*result = "invalid"
}	


func main() {
	// bisa digunakan untuk pengujian test case
	// {"081111111111", "Telkomsel"},
	// {"081211111111", "Telkomsel"},
	// {"081311111111", "Telkomsel"},
	// {"081411111111", "Telkomsel"},
	// {"081511111111", "Telkomsel"},
	// {"6281111111111", "Telkomsel"},
	// {"6281211111111", "Telkomsel"},
	// {"6281311111111", "Telkomsel"},
	var number = "081211111111"
	var result string

	PhoneNumberChecker(number, &result)
	fmt.Println(result)
}
