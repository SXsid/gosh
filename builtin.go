package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type commandHandler func(args []Token)

var BuilteinCommands = map[string]commandHandler{
	"pwd":  pwdHandler,
	"type": typeHandler,
	"exit": exitHandler,
	"echo": exitHandler,
}

func isBuiltin(command string) bool {
	_, ok := BuilteinCommands[command]

	return ok
}

func execBuiltin(command string, args []Token) {
	BuilteinCommands[command](args)
}

func pwdHandler(args []Token) {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Fprint(os.Stderr, err)
	}
	fmt.Println(dir)
}

func exitHandler(args []Token) {
	os.Exit(0)
}

func typeHandler(args []Token) {
	for _, arg := range args {
		ok := isBuiltin(arg.value)
		if ok {
			fmt.Printf("%s is a shell builtin\n", arg)
		} else {
			value, err := exec.LookPath(arg.value)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%s: not found\n", arg)
			} else {
				fmt.Printf("%s is %s\n", arg, value)
			}
		}
	}
}

func echoHandler(args []string) {
	fmt.Print(strings.Join(args, " "))
}
