package main
import (
	"fmt"
	"sort"
)
func Sortheight(height []int) []int {
	sort.Ints(height)
    	return height
}
func main() {
	fmt.Println(Sortheight([]int{172, 170, 150, 165, 144, 155, 159} ))

}