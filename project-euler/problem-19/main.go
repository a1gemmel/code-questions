package main

import "fmt"

var daysByMonth = []int{
	31, // January
	28, // February
	31, // March
	30, // April
	31, // May
	30, // June
	31, // July
	31, // August
	30, // September
	31, // October
	30, // November
	31, // December
}

func daysByMonthInYear(year int) []int {

	leapYear := (year%4 == 0) && (year%100 != 0 || year%400 == 0)
	if leapYear {
		daysByMonth[1] += 1
	}
	return daysByMonth
}

func main() {

	// Considering Sunday to be 0
	dayOfWeek := 1
	sundaysOnFirst := 0

	for year := 1900; year <= 2000; year++ {
		daysByMonth := daysByMonthInYear(year)

		for _, daysInMonth := range daysByMonth {
			if dayOfWeek == 0 && year > 1900 {
				sundaysOnFirst++
			}
			dayOfWeek = (dayOfWeek + daysInMonth) % 7
		}
	}

	fmt.Println(sundaysOnFirst)

}
