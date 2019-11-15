package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/eehut/hello-go/command"
)

/* for console message */
var chConsole = make(chan string, 1)

func consoleInit() {

	/* reading stdin routine */
	go func() {

		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			//fmt.Println(scanner.Text())
			chConsole <- scanner.Text()
		}
	}()

	fmt.Printf("Press enter to active this console!")
}

func consoleInput(str string) {

	/* do command */
	command.DoCommand(str)

	/* print prompt */
	fmt.Printf(command.Prompt)
}

func main() {

	/* dump arguments */
	fmt.Printf("Argc = %d\n", len(os.Args))

	for i := 0; i < len(os.Args); i++ {
		fmt.Printf("Argv[%d] = %s\n", i, os.Args[i])
	}

	print("This built-in function: print\n")

	/* init commands */
	//InitCommand()

	//AddCommand("string", cmdString)
	//AddCommand("printf", cmdPrintf)

	/* init console */
	consoleInit()

	/* make channels */
	chExit := make(chan bool, 1)
	chSignal := make(chan os.Signal, 1)
	signal.Notify(chSignal, syscall.SIGINT, syscall.SIGTERM /*, syscall.SIGUSR1, syscall.SIGUSR2*/)

	/* say hello world */
	//fmt.Println("Hello world")

	/* man loop */
	for {
		select {
		case s := <-chSignal:
			fmt.Printf("Receive Signal:%d\n", s)

			if s == syscall.SIGINT || s == syscall.SIGTERM {
				chExit <- true
			}
			// TODO: change to switch ?

		case cmd := <-chConsole:
			consoleInput(cmd)

		case <-chExit:
			fmt.Println("Program Exit")
			goto MainExit
		}

		//fmt.Println("Select again...")
	}

	/* main xit */
MainExit:
	fmt.Println("Byebye!!!")

}
