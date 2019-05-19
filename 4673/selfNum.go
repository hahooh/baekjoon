package main

import (
	"fmt"
)

func main() {
	var mem [10000]bool

	for i := 1; i <= 10000; i++ {
		nextNum := getNextNum(i)
		if nextNum > 10000 {
			continue
		}
		mem[nextNum-1] = true
	}

	for index, value := range mem {
		if !value {
			fmt.Println(index + 1)
		}
	}
}

func getNextNum(number int) int {
	var total = 0
	currNum := number
	for currNum >= 1 {
		total += currNum % 10
		currNum /= 10
	}
	return number + total
}
