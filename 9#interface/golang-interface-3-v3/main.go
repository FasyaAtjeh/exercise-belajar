package main

import (
	"fmt"
)

type Time struct {
	Hour   int
	Minute int
}

func ChangeToStandartTime(time interface{}) string {
	var hour, minute int

	switch v := time.(type) {
	case string:
		_, err := fmt.Sscanf(v, "%d:%d", &hour, &minute)
		if err != nil {
			return "Invalid input"
		}
	case []int:
		if len(v) != 2 {
			return "Invalid input"
		}
		hour, minute = v[0], v[1]
	case map[string]int:
		var ok bool
		hour, ok = v["hour"]
		if !ok {
			return "Invalid input"
		}
		minute, ok = v["minute"]
		if !ok {
			return "Invalid input"
		}
	case Time:
		hour, minute = v.Hour, v.Minute
	default:
		return "Invalid input"
	}

	if hour < 0 || hour > 23 || minute < 0 || minute > 59 {
		return "Invalid input"
	}

	var amPm string
	if hour == 0 {
		hour = 12
		amPm = "AM"
	} else if hour == 12 {
		amPm = "PM"
	} else if hour < 12 {
		amPm = "AM"
	} else {
		hour -= 12
		amPm = "PM"
	}

	return fmt.Sprintf("%02d:%02d %s", hour, minute, amPm)
}

func main() {
	fmt.Println(ChangeToStandartTime("16:00"))
	fmt.Println(ChangeToStandartTime([]int{16, 0}))
	fmt.Println(ChangeToStandartTime(map[string]int{"hour": 16, "minute": 0}))
	fmt.Println(ChangeToStandartTime(Time{16, 0}))
}
