package main

import "fmt"

func CountProfit(data [][][2]int) []int {
	if len(data) == 0 {
        return []int{}
    }
    var totalProfits []int
    for i := 0; i < len(data[0]); i++ {
        totalProfit := 0
        for j := 0; j < len(data); j++ {
            sales := data[j][i][0]
            expenses := data[j][i][1]
            profit := sales - expenses
            totalProfit += profit
        }
        totalProfits = append(totalProfits, totalProfit)
    }
    return totalProfits 
}
func main() {
	fmt.Println(CountProfit([][][2]int{{{1000, 800},{700, 500}}, {{1000, 800},{900, 200}}}))

}