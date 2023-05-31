package main

import "fmt"

func TicketPlayground(height, age int) int {
	if age > 12 {
		return 100_000
	}
	if age == 12 || height > 160 {
	    return 60_000
	}
	if age >= 10 && age <= 11 || height > 150 {
		return 40_000
	}
	if age >= 8 && age <= 9 || height > 135   {
        return 25_000
    }
	if age >= 5 && age <= 7 || height > 120 {
		return 15_000
	}
	if age < 5 {
		return -1
	}

	return -1 // TODO: replace this
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(TicketPlayground(160, 11))
}
