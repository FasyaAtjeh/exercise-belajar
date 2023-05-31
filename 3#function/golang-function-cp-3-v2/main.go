package main

import (
	"fmt"
)

func nameShort(name1 string, name2 string) string {
	if len(name1) < len(name2) {
		return name1
	} else if len(name1) > len(name2) {
		return name2
	} else {
		if name1 < name2 {
			return name1
		} else{
			return name2
        }
	}
}
func FindShortestName(name string) string {
	curName := ""
	first := true
	minName := ""
	for _, c := range name {
		if string(c) == ";" || string(c) == " " || string(c) =="," {
			if first == true {
				minName = curName
				first = false
			} else {
				minName = nameShort(curName, minName)
			}
			curName = ""
		} else {
			curName += string(c)
		}
	}
	if curName != "" {
		if first == true {
			minName = curName
			first = false
		} else {
			minName = nameShort(curName, minName)
		}

	}
	return minName
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(FindShortestName("Hanif Joko Tio Andi Budi Caca Hamdan")) // "Tio"
	fmt.Println(FindShortestName("Budi;Tia;Tio"))                         // "Tia"
}
