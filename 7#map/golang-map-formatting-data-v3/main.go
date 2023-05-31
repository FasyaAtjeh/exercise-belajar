package main

import (
	"fmt"
	"strconv"
	"strings"
)

func ChangeOutput(data []string) map[string][]string {
	result := make(map[string][]string)
	for _, d := range data {
		fmt.Println("d", d)
		tokens := strings.Split(d, "-")
		fmt.Printf("tokens: %#v\n", tokens)

		label := tokens[0]
		index, _ := strconv.Atoi(tokens[1])
		firstOrLast := tokens[2]
		value := tokens[3]
		fmt.Println("label", label)
		fmt.Println("index", index)
		fmt.Println("firstOrLast", firstOrLast)
		fmt.Println("value", value)

		if _, ok := result[label]; !ok {
			result[label] = make([]string, 0)
		}
		if firstOrLast == "first" {
			if index >= len(result[label]) {
				result[label] = append(result[label], value)
			} else {
				result[label][index] = value + result[label][index]
			}
		} else {
			if index >= len(result[label]) {
				result[label] = append(result[label], value)
			} else {
				result[label][index] = result[label][index] + " " + value
			}
		}
		fmt.Println()

	}
	fmt.Println("result", result)
	return result
}

// bisa digunakan untuk melakukan debug
func main() {
	data := []string{"account-0-first-John", "account-0-last-Doe", "account-1-first-Jane", "account-1-last-Doe", "address-0-first-Jaksel", "address-0-last-Jakarta", "address-1-first-Bandung", "address-1-last-Jabar", "phone-0-first-081234567890", "phone-1-first-081234567891"}
	res := ChangeOutput(data)

	fmt.Println(res)
}
