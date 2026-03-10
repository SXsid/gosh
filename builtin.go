package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
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

func cdHandler(args []Token) {
	if len(args) <= 0 {
		return
	}
	path := args[0].value
	switch path[0] {
	// TODO: impletmt the realvative path
	case '~':
		homepath, err := os.UserHomeDir()
		if err != nil {
			return
		}
		restPath := ""
		if len(path) > 1 {
			restPath = path[1:]
		}
		fielpath := filepath.Join(homepath, restPath)
		if err := os.Chdir(fielpath); err != nil {
			fmt.Printf("cd: %s: No such file or directory\n", fielpath)
		}
	case '/':
		if err := os.Chdir(path); err != nil {
			fmt.Printf("cd: %s: No such file or directory\n", path)
		}
	default:
		return

	}
}
