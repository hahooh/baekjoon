package main

import (
	"fmt"
)

func main() {
	var month, day int
	fmt.Scanf("%d %d", &month, &day)

	dNames := []string{"SUN", "MON", "TUE", "WED", "THU", "FRI", "SAT"}

	totalDays := 0
	for m := 1; m <= month; m++ {
		if m == month {
			totalDays += day
		} else {
			totalDays += getDaysOfMonth(m)
		}
	}

	fmt.Println(dNames[totalDays%len(dNames)])
}

func getDaysOfMonth(month int) int {
	switch month {
	case 2:
		return 28
	case 4, 6, 9, 11:
		return 30
	}

	return 31
}
