package main
import "fmt"

func ReverseData(arr [5]int) [5]int {
	for i := 0; i < len(arr)/2; i++ {
		j := len(arr) - i - 1
		arr[i], arr[j] = arr[j], arr[i]
	}
	for i := 0; i < len(arr); i++ {
		num := arr[i]
		reversedNum := 0

		for num > 0 {
			reversedNum = reversedNum*10 + num%10
			num /= 10
		}

		arr[i] = reversedNum
	}
	return arr
}
func main () {
	fmt.Println(ReverseData([5]int{123, 456, 11, 1, 2}))
}
