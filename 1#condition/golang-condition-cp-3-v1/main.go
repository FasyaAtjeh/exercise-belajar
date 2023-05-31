package main

import "fmt"

func GetPredicate(math, science, english, indonesia int) string {
	var predicate string
	var avg = (math + science + english + indonesia) / 4
	fmt.Println(avg)
	if avg == 100 {
		predicate = "Sempurna"
	} else if avg >= 90 {
		predicate = "Sangat Baik"
	} else if avg >= 80 {
        predicate = "Baik"
     } else if avg >= 70 {
        predicate = "Cukup"
     } else if avg >= 60 {
        predicate = "Kurang"
     } else {
        predicate = "Sangat kurang"
     }
 	return predicate
	}
// gunakan untuk melakukan debug
func main() {
	fmt.Println(GetPredicate(50, 80, 100, 60))
}
