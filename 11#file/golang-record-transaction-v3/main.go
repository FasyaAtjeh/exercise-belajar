package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

type Transaction struct {
	Date   string
	Type   string
	Amount int
}

func RecordTransactions(path string, transactions []Transaction) error {
	if len(transactions) == 0 {
		return nil
	}
	sort.Slice(transactions, func(i int, j int) bool {
		if transactions[i].Date < transactions[j].Date {
			return true
		} else {
			return false
		}
	})
	lastDate := transactions[0].Date
	money := 0
	output := make([]string, 0)

	for _, transaction := range transactions {
		if transaction.Date == lastDate {
			if transaction.Type == "income" {
				money += transaction.Amount
			} else {
				money -= transaction.Amount
			}
		} else {
			if money < 0 {
				output = append(output, fmt.Sprintf("%s;expense;%d", lastDate, -money))
			} else {
				output = append(output, fmt.Sprintf("%s;income;%d", lastDate, money))
			}
			money = 0
			if transaction.Type == "income" {
				money += transaction.Amount
			} else {
				money -= transaction.Amount
			}
			lastDate = transaction.Date
		}
	}
	if money < 0 {
		output = append(output, fmt.Sprintf("%s;expense;%d", lastDate, -money))
	} else {
		output = append(output, fmt.Sprintf("%s;income;%d", lastDate, money))
	}

	f, err := os.Create(path)

	if err != nil {
		return err
	}
	defer f.Close()
	fmt.Println(output)
	_, err2 := f.WriteString(strings.Join(output, "\n"))

	if err2 != nil {
		return err2
	}
	return nil
}
func main() {
	// bisa digunakan untuk pengujian test case
	var transactions = []Transaction{
		{"01/01/2021", "income", 100000},
		{"01/01/2021", "expense", 50000},
		{"01/01/2021", "expense", 30000},
		{"01/01/2021", "income", 20000},
	}

	err := RecordTransactions("transactions.txt", transactions)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success")
}
