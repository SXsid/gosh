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
	path  TokenType = "PATH"
)

type Token struct {
	Type  TokenType
	value string
}

func extract_word(input string, i int, char rune) (string, int, error) {
	j := i + 1
	for len(input) > j && rune(input[j]) != char {
		j++
	}
	if len(input) > j {
		j++
	} else if rune(input[j-1]) != char {
		return "", i, fmt.Errorf("incoorect syntax %c^", rune(input[j-1]))
	}

	return input[i:j], j, nil
}

func extractPath(input string, i int) (string, int) {
	j := i + 1
	for j < len(input) && !unicode.IsSpace(rune(input[j])) {
		j++
	}
	return input[i:j], j
}

func parser(input string) (string, []Token) {
	i := 0
	data := make([]Token, 0)
	args := make([]Token, 0)
loop:
	for i < len(input) {

		curr := rune(input[i])
		if unicode.IsSpace(curr) {
			i++
			continue

		}
		switch curr {
		case '"':
			value, j, err := extract_word(input, i, '"')
			if err != nil {
				fmt.Println(err.Error())

				break loop
			}
			data = append(data, Token{
				Type:  quote,
				value: value,
			})
			i = j
			continue

		case '\'':

			value, j, err := extract_word(input, i, '\'')
			if err != nil {
				fmt.Println(err.Error())

				break loop
			}
			data = append(data, Token{
				Type:  quote,
				value: value[1 : len(value)-1],
			})
			i = j

			continue
		case '|':
			data = append(data, Token{
				Type:  pipe,
				value: string(curr),
			})
		case '/':
			value, j := extractPath(input, i)
			i = j
			data = append(data, Token{
				Type:  path,
				value: value,
			})
			continue
		case '~':
			value, j := extractPath(input, i)
			i = j
			data = append(data, Token{
				Type:  path,
				value: value,
			})

			continue
		case '.':
			value, j := extractPath(input, i)
			i = j
			data = append(data, Token{
				Type:  path,
				value: value,
			})

			continue
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
				continue
			} else {

				fmt.Fprintf(os.Stderr, "invalid '%c' input ", curr)
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
			fmt.Printf("%s: command not found\n", command)
		} else {
			fmt.Print(string(output))
		}
	}
}
