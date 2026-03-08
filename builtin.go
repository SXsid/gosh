package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type commandHandler func(args []Token)

var BuilteinCommands map[string]commandHandler

func Init() {
	BuilteinCommands = map[string]commandHandler{
		"pwd":  pwdHandler,
		"type": typeHandler,
		"exit": exitHandler,
		"cd":   cdHandler,
		"echo": echoHandler,
	}
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
			fmt.Printf("%s is a shell builtin\n", arg.value)
		} else {
			value, err := exec.LookPath(arg.value)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%s: not found\n", arg.value)
			} else {
				fmt.Printf("%s is %s\n", arg.value, value)
			}
		}
	}
}

func echoHandler(args []Token) {
	valueArray := make([]string, len(args))
	if len(valueArray) == 0 {
		return
	}
	for i, data := range args {
		valueArray[i] = data.value
	}
	fmt.Println(strings.Join(valueArray, " "))
}

// TODO: implement cd command
func cdHandler(args []Token) {
	path := args[0].value
	if len(path) <= 0 {
		return
	}
	switch path[0] {
	case '~':
	case '/':
		if err := os.Chdir(path); err != nil {
			fmt.Println(err.Error())
		}
	default:

	}
}
