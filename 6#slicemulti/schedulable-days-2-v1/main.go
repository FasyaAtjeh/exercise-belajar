package main

import "fmt"

func SchedulableDays(villager [][]int) []int {
    maxDate := -1
    for _, schedule := range villager  {
        for _, date := range schedule {
            if date > maxDate {
                maxDate = date
            }
        }
    }

    count := make([]int, maxDate+1)    
    for _, schedule := range villager {
        for _, date := range schedule {
            count[date]++
        }
    }
    newSchedule := []int{}
    for date := 1; date <= maxDate; date++ {
        if count[date] == len(villager) {
            newSchedule = append(newSchedule, date)
        }
    }
    
    return newSchedule
}
func main() {
	fmt.Println(SchedulableDays([][]int{{7, 12, 19, 22}, {12, 19, 21, 23}, {7, 12, 19}, {12, 19}}))

}