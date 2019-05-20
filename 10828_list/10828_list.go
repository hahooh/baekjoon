package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strconv"
)

func main() {
	stack := list.New()
	var numCmd int
	var cmd string
	var el int
	buff := bufio.NewWriter(os.Stdout)
	defer buff.Flush()

	fmt.Scanf("%d", &numCmd)
	for i := 0; i < numCmd; i++ {
		fmt.Scanf("%s", &cmd)
		switch cmd {
		case "push":
			fmt.Scanf("%d", &el)
			stack.PushBack(el)
		case "pop":
			lastEl := stack.Back()
			result := -1
			if lastEl != nil {
				result = lastEl.Value.(int)
				stack.Remove(lastEl)
			}
			buff.WriteString(strconv.Itoa(result) + "\n")
		case "size":
			buff.WriteString(strconv.Itoa(stack.Len()) + "\n")
		case "empty":
			switch stack.Len() {
			case 0:
				buff.WriteString(strconv.Itoa(1) + "\n")
			default:
				buff.WriteString(strconv.Itoa(0) + "\n")
			}
		case "top":
			lastEl := stack.Back()
			result := -1
			if lastEl != nil {
				result = lastEl.Value.(int)
			}
			buff.WriteString(strconv.Itoa(result) + "\n")
		}
	}
}
