package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	for {
		fmt.Print("$ ")
		input, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			panic(err)
		}

		input = strings.TrimSpace(input)
		if len(input) == 0 {
			continue
		}
		command, args := parser(input)
		execute(command, args)

	}
}
