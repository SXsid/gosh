package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	Init()
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("$ ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("no data found ")
			os.Exit(0)
		}

		input = strings.TrimSpace(input)
		if len(input) == 0 {
			continue
		}
		command, args := parser(input)
		execute(command, args)

	}
}
