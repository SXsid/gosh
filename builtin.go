package main

import (
	"fmt"
	"os"
	"os/exec"
)

type commandHandler func(args []string)

var BuilteinCommands = map[string]commandHandler{
	"pwd":  pwdHandler,
	"exit": exitHandler,
}

func isBuiltin(command string) bool {
	_, ok := BuilteinCommands[command]

	return ok
}

func execBuiltin(command string, args []string) {
	BuilteinCommands[command](args)
}

func pwdHandler(args []string) {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Fprint(os.Stderr, err)
	}
	fmt.Println(dir)
}

func exitHandler(args []string) {
	os.Exit(0)
}

func typeHandler(args []string) {
	for _, arg := range args {
		ok := isBuiltin(arg)
		if ok {
			fmt.Printf("%s is a shell builtin\n", arg)
		} else {
			value, err := exec.LookPath(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%s: not found\n", arg)
			} else {
				fmt.Printf("%s is %s\n", arg, value)
			}
		}
	}
}

func echoHandler(args []string) {
}
