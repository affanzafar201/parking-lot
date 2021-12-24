package main

import (
	"bufio"
	"os"
	"processor"
	"strings"
)

func main() {
	parkingServiceProcessor := new(processor.Processor)
	parkingServiceProcessor.WelcomeMessage()

	for {
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimRight(input, "\n")
		parkingServiceProcessor.Execute(input)
	}
}
