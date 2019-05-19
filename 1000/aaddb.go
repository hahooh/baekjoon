package main

import (
	"fmt"
)

func main() {
	//	a, erra := strconv.Atoi(os.Args[1])
	//	b, errb := strconv.Atoi(os.Args[2])

	//	if erra == nil && errb == nil {
	//		fmt.Println(a + b)
	//	}

	var a, b int
	fmt.Scanf("%d %d", &a, &b)
	fmt.Println(a + b)
}
