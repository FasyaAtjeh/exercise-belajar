package main

import "fmt"

func GetTicketPrice(VIP, regular, student, day int) float32 {
	var totalTicket = VIP + regular + student
	var totalPrice = (VIP * 30) + (regular * 20) + (student * 10)

	if totalPrice >= 100 {
		if day%2 != 0 {
			if totalTicket < 5 {
				return float32(totalPrice) - (float32(totalPrice) * 15 / 100)
			} else {
				return float32(totalPrice) - (float32(totalPrice) * 25 / 100)
			}
	}	else {
		if totalTicket < 5 {
                return float32(totalPrice) - (float32(totalPrice) * 10 / 100)
            } else {
                return float32(totalPrice) - (float32(totalPrice) * 20 / 100)
            }
	}
	} else {
		return float32(totalPrice)
	}
}
// gunakan untuk melakukan debug
func main() {
	fmt.Println(GetTicketPrice(1, 1, 1, 20))
}
