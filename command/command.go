package command

import (
	"fmt"
	"strings"
)

// Prompt is command prompt string after 'enter'
const Prompt string = "> "

// Command struct
type Command struct {
	Execute func(argv []string) int
	Name    string
	Desc    string
	Usage   string
}

// AllCommands saves all commands in map
var AllCommands = map[string]*Command{}

// NewCommand returns a new, empty command with the specified infos
func AddCommand(name, desc, usage string, execute func(argv []string) int) *Command {
	f := &Command{
		Name:    name,
		Desc:    desc,
		Usage:   usage,
		Execute: execute}

	AllCommands[name] = f
	return f
}

/* help */
// doHelp is called when command matched
func doHelp(argv []string) int {

	num := len(argv)
	var arg string

	if num < 1 {
		arg = "help"
	} else {
		arg = argv[0]
	}

	c, exist := AllCommands[arg]

	if exist {
		fmt.Printf("%s\n", c.Usage)
	} else {
		fmt.Printf("*** Command not found: %s\n", arg)
	}

	return 0
}

/* list */

// Execute is called when command matched
func doList(argv []string) int {

	for k, v := range AllCommands {
		fmt.Printf(" %-16s : %s\n", k, v.Desc)
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

func init() {

	AddCommand("help", "Show command help", "Usage:\n help CMD", doHelp)
	AddCommand("list", "List built-in commands", "Usage:\n list", doList)
}
