package main

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/eehut/hello-go/command"
)

// CommandUnicode is for unicode test
type CommandUnicode struct {
	Desc command.CommandDesc
}

// GetDesc return the description
func (cmd CommandUnicode) GetDesc() *command.CommandDesc {
	return &cmd.Desc
}

// Execute is command function
func (cmd CommandUnicode) Execute(argv []string) int {

	s := "Hello A B C!"

	for _, r := range s {
		if unicode.IsUpper(r) {
			fmt.Println(r, "is upper letter")
		}
	}

	ss := "Hello 你好"

	for _, r := range ss {
		fmt.Printf("%c", unicode.ToUpper(r))
	}

	for _, r := range ss {
		fmt.Printf("%c", unicode.To(unicode.LowerCase, r))
	}

	fmt.Println()

	fmt.Println(strings.ToUpper(ss))

	fmt.Printf(" => runs(char): %q\n", []rune(ss))
	fmt.Printf(" => runs(hex): %x\n", []rune(ss))
	fmt.Printf(" => bytes(hex): [% x]\n", []byte(ss))

	for i, c := range ss {
		fmt.Printf("%d: %q [% x] [%x]\n", i, c, []byte(string(c)), c)
	}

	return 0
}

func init() {
	command.AddCommand("unicode", CommandUnicode{Desc: command.CommandDesc{Info: "unicode package test", Help: "Usage:\n unicode"}})
}
