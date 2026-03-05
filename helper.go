package main

import (
	"fmt"
	"unicode"
)

// parser the input string into command line and argument and refien those
func parser(input string) (string, []string) {
	data := make([]string, 0)
	args := make([]string, 0)
	currentWord := ""
	doubleQuoted := ""
	for _, rune := range input {
		char := string(rune)
		fmt.Println(char)
		if unicode.IsSpace(rune) {
			if len(currentWord) > 0 {
				data = append(data, currentWord)
				currentWord = ""
			} else if len(doubleQuoted) > 0 {
				doubleQuoted += char
			}
		} else if rune == '"' {
			if doubleQuoted == "" {
				doubleQuoted += char
			} else {
				data = append(data, doubleQuoted)
				doubleQuoted = ""
			}
		} else {
			if doubleQuoted == "" {
				currentWord += char
			} else {
				doubleQuoted += char
			}
		}
	}
	if len(data) >= 2 {
		args = data[1:]
	}
	return data[0], args
}

// decide whetehre it's buldign or extern or not found
func execute(command string, args []string) {
}
