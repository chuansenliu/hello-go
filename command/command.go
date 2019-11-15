package command

import (
	"fmt"
	"strings"
)

// Prompt is command prompt string after 'enter'
const Prompt string = "> "

// CommandDesc is a struct for command description
type CommandDesc struct {
	Info string
	Help string
}

// Command is an interface
type Command interface {
	Execute([]string) int
	GetDesc() *CommandDesc
}

// AllCommands saves all commands in map
var AllCommands = map[string]Command{}

/* help */

// CommandHelp is internal command 'help'
type CommandHelp struct {
	Desc CommandDesc
}

// GetDesc return the description of this command
func (cmd CommandHelp) GetDesc() *CommandDesc {
	return &cmd.Desc
}

// Execute is called when command matched
func (cmd CommandHelp) Execute(argv []string) int {

	num := len(argv)
	var arg string

	if num < 1 {
		arg = "help"
	} else {
		arg = argv[0]
	}

	c, exist := AllCommands[arg]

	if exist {
		d := c.GetDesc()
		fmt.Printf("%s\n", d.Help)
	} else {
		fmt.Printf("*** Command not found: %s\n", arg)
	}

	return 0
}

/* list command */

// CommandList is internal command 'list'
type CommandList struct {
	Desc CommandDesc
}

// GetDesc return the description of this command
func (cmd CommandList) GetDesc() *CommandDesc {
	return &cmd.Desc
}

// Execute is called when command matched
func (cmd CommandList) Execute(argv []string) int {

	for k, v := range AllCommands {
		desc := v.GetDesc()
		fmt.Printf(" %-16s : %s\n", k, desc.Info)
	}

	return 0
}

// StringSplit splits string into pices
// use strings.Fields is better
/*
func StringSplit(cmd string) []string {

	args := strings.Split(cmd, " ")
	var ret []string

	for i := 0; i < len(args); i++ {
		if len(args[i]) > 0 {
			ret = append(ret, args[i])
		}
	}

	return ret
}
*/

// DoCommand is command entry for input string
func DoCommand(str string) {

	/* empty, return */
	if len(str) <= 0 {
		return
	}

	//args := StringSplit(str)
	args := strings.Fields(str)

	/*
		fmt.Printf("Argc = %d\n", len(args))

		for i := 0; i < len(args); i ++ {
			fmt.Printf("Argv[%d] = %s\n", i, args[i])
		}
	*/
	cmd, exist := AllCommands[args[0]]

	if exist {
		ret := cmd.Execute(args[1:])
		fmt.Printf("\n==> return %d\n", ret)
	} else {
		fmt.Printf("*** Command not found: %s\n", args[0])
	}
}

// AddCommand adds command to AllCommands
func AddCommand(name string, cmd Command) {
	AllCommands[name] = cmd
}

func init() {

	AddCommand("list", CommandList{Desc: CommandDesc{"List built-in commands", "Usage:\n list"}})
	AddCommand("help", CommandHelp{Desc: CommandDesc{"Show command help", "Usage:\n help CMD"}})
}
