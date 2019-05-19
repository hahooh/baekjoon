package main

import "fmt"

func main() {
	var kg int
	fmt.Scanf("%d", &kg)

	var results []int

	var result int
	for i := 0; i <= kg; i++ {
		if i == 0 {
			result = 0
		} else if i < 3 {
			result = kg + 1
		} else if i == 3 {
			result = 1
		} else if i == 5 {
			result = 1
		} else {
			threePrev, fivePrev := i-3, i-5
			if threePrev < 0 {
				result = kg + 1
			} else if fivePrev < 0 {
				result = results[threePrev]
			} else {
				result = min(results[threePrev], results[fivePrev]) + 1
			}
		}
		results = append(results, result)
	}

	if results[kg] > kg {
		fmt.Println(-1)
	} else {
		fmt.Println(results[kg])
	}
}

func min(inta int, intb int) int {
	if inta < intb {
		return inta
	}

	return intb
}
