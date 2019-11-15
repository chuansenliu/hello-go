package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"../command"
)

// CommandString is for string test
type CommandString struct {
	Desc command.CommandDesc
}

// GetDesc return the description
func (cmd CommandString) GetDesc() *command.CommandDesc {
	return &cmd.Desc
}

// Execute is command function
func (cmd CommandString) Execute(argv []string) int {

	b := "abc efg hijk lmn\n\t"

	c := strings.TrimSpace(b)

	fmt.Printf("len of c = %d %s\n", len(c), c)

	fmt.Println(strings.Contains(b, "hij"))

	fmt.Println(strings.Index(c, "efg"))
	fmt.Println(strings.LastIndex(c, " "))
	fmt.Println(strings.Count(c, " "))

	reader := strings.NewReader("This is reader test")
	buf := make([]byte, 4)

	reader.WriteTo(os.Stdout)

	reader.Seek(0, io.SeekStart)

	for {
		fmt.Println(reader.Len())
		count, err := reader.Read(buf)

		if err != nil {
			if err == io.EOF {
				fmt.Println("EOF: ", count)
				break
			}

			fmt.Println(err)
			break
		}

		fmt.Println(count, string(buf[:count]))
	}

	return 0
}

func init() {
	command.AddCommand("string", CommandString{Desc: command.CommandDesc{Info: "Strings test", Help: "Usage:\n string"}})
}
