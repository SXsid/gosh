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
		parser(input)

	}
}
