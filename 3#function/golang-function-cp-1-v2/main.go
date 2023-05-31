package main

import (
	"fmt"
)

func DateFormat(day, month, year int) string {
	var format_month string
	if month == 1 {
		format_month = "January"
	} else if month == 2 {
		format_month = "February"
    } else if month == 3 {
		format_month = "March"
	} else if month == 4 {
		format_month = "April"
    } else if month == 5 {
		format_month = "May"
	} else if month == 6 {
		format_month = "June"
    } else if month == 7 {
		format_month = "July"
	} else if month == 8 {
        format_month = "August"
    } else if month == 9 {
		format_month = "September"
	} else if month == 10 {
		format_month = "October"
    } else if month == 11 {
		format_month = "November"
	} else if month == 12 {
        format_month = "December"
    } 

	var format_day string
	if day < 10 {
		format_day = "0" + fmt.Sprint(day)
	} else {
		format_day = fmt.Sprint(day)
	}
	result := format_day + "-" + format_month + "-" + fmt.Sprint(year)

	return result // TODO: replace this
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(DateFormat(1, 1, 2012))
}
