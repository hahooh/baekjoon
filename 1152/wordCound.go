package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	sentence, _ := reader.ReadString('\n')
	words := strings.Fields(sentence)
	fmt.Println(len(words))
}
