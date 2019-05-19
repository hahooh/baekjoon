package main

import (
	"container/list"
	"fmt"
)

func main() {
	stack := list.New()
	var numCmd int
	var cmd string
	var el int
	var results []int
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
			results = append(results, result)
		case "size":
			results = append(results, stack.Len())
		case "empty":
			if stack.Len() == 0 {
				results = append(results, 1)
			} else {
				results = append(results, 0)
			}
		case "top":
			lastEl := stack.Back()
			result := -1
			if lastEl != nil {
				result = lastEl.Value.(int)
			}
			results = append(results, result)
		}
	}

	for _, result := range results {
		fmt.Println(result)
	}
}
