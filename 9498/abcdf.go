package main

import "fmt"

func main() {
	var percentage int
	fmt.Scanf("%d", &percentage)
	fmt.Println(getGrade(percentage))
}

func getGrade(percentage int) string {
	if percentage < 60 {
		return "F"
	} else if 60 <= percentage && percentage < 70 {
		return "D"
	} else if 70 <= percentage && percentage < 80 {
		return "C"
	} else if 80 <= percentage && percentage < 90 {
		return "B"
	} else {
		return "A"
	}
}
