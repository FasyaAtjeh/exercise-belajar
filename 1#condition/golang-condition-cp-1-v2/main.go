package main

import "fmt"

func GraduateStudent(score int, absent int) string {
	if score >= 70 && absent < 5 {// TODO: replace this
	    return "lulus"
	} else {
	return "tidak lulus"
	}
}
// gunakan untuk melakukan debug
func main() {
	fmt.Println(GraduateStudent(70, 4))
}
