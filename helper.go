package main

import (
	"fmt"
	"os"
	"os/exec"
	"unicode"
)

type TokenType string

const (
	word  TokenType = "WORD"
	quote TokenType = "QUOTE"
	pipe  TokenType = "PIPE"
)

type Token struct {
	Type  TokenType
	value string
}

// parser the input string into command line and argument and refien those
func parser(input string) (string, []Token) {
	i := 0
	data := make([]Token, 0)
	args := make([]Token, 0)
	for i < len(input) {

		curr := rune(input[i])
		if unicode.IsSpace(curr) {
			i++
			continue

		}
		switch curr {
		case '"':
			j := i + 1
			for rune(input[j]) != '"' {
				if len(input) <= j {
					fmt.Fprint(os.Stderr, "inccorec sytax")
					os.Exit(1)
				}
				j++
			}
			data = append(data, Token{
				Type:  quote,
				value: input[i : j+1],
			})
			i = j
		case '|':
			data = append(data, Token{
				Type:  pipe,
				value: string(curr),
			})
		default:

			if unicode.IsLetter(curr) || unicode.IsDigit(curr) {
				j := i
				for j < len(input) && !unicode.IsSpace(rune(input[j])) {
					j++
				}
				data = append(data, Token{
					Type:  word,
					value: string(input[i:j]),
				})
				i = j
			} else {

				fmt.Fprint(os.Stderr, "invalid '%s' input ")
				os.Exit(1)
			}
		}
		i++

	}
	if len(data) >= 2 {
		args = data[1:]
	}
	return data[0].value, args
}

// decide whetehre it's buldign or extern or not found
func execute(command string, args []Token) {
	tokenValue := make([]string, len(args))
	for i, value := range args {
		tokenValue[i] = value.value
	}
	if isBuiltin(command) {
		execBuiltin(command, args)
	} else {
		cmd := exec.Command(command, tokenValue...)
		output, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Printf("%s ivalid command\n", command)
		} else {
			fmt.Println(string(output))
		}
	}
}
