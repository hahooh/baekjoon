package main

import (
	"fmt"
	"os"
	"strconv"
)

// MyStack
// This is my stack implementation
type MyStack struct {
	s []int
}

func (stack *MyStack) push(element int) {
	stack.s = append(stack.s, element)
}

func (stack *MyStack) pop() int {
	if stack.size() == 0 {
		return -1
	}
	result := stack.s[stack.size()-1]
	stack.s = stack.s[:stack.size()-1]
	return result
}

func (stack *MyStack) size() int {
	return len(stack.s)
}

func (stack *MyStack) empty() int {
	if stack.size() > 0 {
		return 0
	}
	return 1
}

func (stack *MyStack) top() int {
	if stack.size() == 0 {
		return -1
	}
	return stack.s[stack.size()-1]
}

func insertBuffer(buff []byte, results *[]byte) {
	for _, value := range buff {
		*results = append(*results, value)
	}
}

func main() {
	stack := MyStack{}
	var numCmd int
	fmt.Scanf("%d", &numCmd)
	var inputElement int
	var cmd string
	var results []byte
	for i := 0; i < numCmd; i++ {
		fmt.Scanf("%s", &cmd)

		switch cmd {
		case "push":
			fmt.Scanf("%d", &inputElement)
			stack.push(inputElement)
		case "pop":
			insertBuffer([]byte(strconv.Itoa(stack.pop())+"\n"), &results)
		case "size":
			insertBuffer([]byte(strconv.Itoa(stack.size())+"\n"), &results)
		case "empty":
			insertBuffer([]byte(strconv.Itoa(stack.empty())+"\n"), &results)
		case "top":
			insertBuffer([]byte(strconv.Itoa(stack.top())+"\n"), &results)
		}
	}
	os.Stdout.Write(results)
}
