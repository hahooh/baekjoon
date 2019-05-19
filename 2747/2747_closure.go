package main

import (
	"fmt"
)

func fibo() func(n int) int {
	prev1 := 0
	prev2 := 0
	return func(n int) int {
		if n == 0 {
			return 0
		} else if n == 1 {
			prev2 = 1
			return 1
		} else {
			result := prev1 + prev2
			prev1 = prev2
			prev2 = result
			return result
		}
	}
}

func main() {
	var n int
	fmt.Scanf("%d", &n)

	fib := fibo()
	for i := 0; i < n; i++ {
		fib(i)
	}
	fmt.Println(fib(n))
}
