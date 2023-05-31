package main
import "fmt"

func ExchangeCoin(amount int) []int {
	if amount == 0 {
        return []int{}
    }
	coins := []int{1000, 500, 200, 100, 50, 20, 10, 5, 1}
    var result []int

    for _, currentCoin := range coins  {
        if amount >= currentCoin {
            check := amount / currentCoin
            amount = amount % currentCoin
            for i := 0; i < check; i++ {
                result = append(result, currentCoin)
			}
		}
	}
	return result
}

func main (){
	fmt.Println(ExchangeCoin(0))

}