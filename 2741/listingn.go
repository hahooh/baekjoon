package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// input, err := strconv.Atoi(os.Args[1])
	var input int
	fmt.Scanf("%d", &input)

	/*
		if err != nil {
			fmt.Println("hello world")
		}
	*/

	i := 0
	var buf []byte
	for i < input {
		i++
		buff := []byte(strconv.Itoa(i) + "\n")
		for j := range buff {
			buf = append(buf, buff[j])
		}
	}
	os.Stdout.Write(buf)
}
