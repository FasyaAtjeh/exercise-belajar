package main

import "fmt"

type Product struct {
	Name  string
	Price int
	Tax   int
}

func MoneyChanges(amount int, products []Product) []int {
	total := 0
	for _, product := range products {
		total += product.Price + product.Tax
	}

	// Hitung kembalian
	changes := make([]int, 0)
	remaining := amount - total
	for _, value := range []int{1000, 500, 200, 100, 50, 20, 10, 5, 1} {
		if remaining == 0 {
			break
		}
		count := remaining / value
		for i := 0; i < count; i++ {
			changes = append(changes, value)
			remaining -= value
		}
	}

	return changes
}
func main() {
	// products := []Product{
	// 	{Name: "Baju", Price: 5000, Tax: 500},
	// 	{Name: "Celana", Price: 3000, Tax: 300},
	// }
	// amount := 10000
	products := []Product{
		{Name: "baju 1", Price: 10000, Tax: 1000},
		{Name: "Sepatu", Price: 15550, Tax: 1555},
	}
	amount := 30000
	changes := MoneyChanges(amount, products)
	fmt.Println("Changes:", changes)
}
