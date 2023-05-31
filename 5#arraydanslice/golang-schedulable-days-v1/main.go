package main
import "fmt"

func SchedulableDays(date1 []int, date2 []int) []int {
	newSchedule := []int{}
	for _, imamDate := range date1 { 
		for _, permanaDate := range date2  {
			if imamDate == permanaDate {
				newSchedule = append(newSchedule, imamDate)
				break
			}
		}
	}

	return newSchedule
}
func main() {
	fmt.Println(SchedulableDays([]int{1, 2, 3, 4}, []int{3, 4, 5} ))

}