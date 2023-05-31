package main

import (
	"fmt"
	"strconv"
)

func ChangeToCurrency(change int) string {
	//Don't worry about this implementation
	//Just use this function :)

	var res = ""

	moneyStr := strconv.Itoa(change)

	for i := len(moneyStr) - 1; i >= 0; i-- {
		if i == len(moneyStr)-1 {
			res = string(moneyStr[i]) + res
		} else if (len(moneyStr)-i)%3 == 0 {
			res = "." + string(moneyStr[i]) + res
		} else {
			res = string(moneyStr[i]) + res
		}
	}

	if res[0] == '.' {
		res = res[1:]
	}

	return "Rp. " + res
}

func MoneyChange(money int, listPrice ...int) string {
	for _, itemPrice := range listPrice {
		money -= itemPrice
		if money < 0 {
            return "Uang tidak cukup"
        }
	}
	
	var res = ""
	moneyString := strconv.Itoa(money)

	for i := len(moneyString) - 1; i >= 0; i-- {
		if i == len(moneyString)-1 {
            res = string(moneyString[i]) + res
        } else if (len(moneyString)-i)%3 == 0 {
            res = "." + string(moneyString[i]) + res
        } else {
            res = string(moneyString[i]) + res
        }
	}
	if res[0] == '.' {
		res = res[1:]
    }
	return "Rp. " +res // TODO: replace this
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(MoneyChange(100000, 50000, 10000, 10000, 5000, 5000))
}
