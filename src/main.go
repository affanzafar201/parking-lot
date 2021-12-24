package main

import (
	"bufio"
	"os"
	"processor"
	"strings"
)

func main() {
	v := new(processor.Processor)
	for {
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimRight(input, "\n")
		v.Execute(input)
	}
}
