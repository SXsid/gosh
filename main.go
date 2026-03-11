//go:build !js

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
	fileinfo, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}
	isInteractive := fileinfo.Mode()&os.ModeCharDevice != 0
	for {
		if isInteractive {
			fmt.Print("$ ")
			os.Stdout.Sync()
		}
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
		// fmt.Println(command)
		// fmt.Println(args)
		execute(command, args)

	}
}
