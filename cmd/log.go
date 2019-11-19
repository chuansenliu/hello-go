package main

import (
	"fmt"
	"log"
	"os"

	"github.com/eehut/hello-go/command"
)

func fileOpen(path string) (*os.File, error) {

	_, err := os.Stat(path)

	if os.IsNotExist(err) {
		return os.Create(path)
	}
	return os.OpenFile(path, os.O_APPEND, 0666)
}

// doLog is command function
func doLog(argv []string) int {

	var logFile *os.File
	var err error

	if len(argv) > 0 && argv[0] == "file" {
		logFile, err = fileOpen("./mylog.txt")

		if err != nil {
			fmt.Println("Open log file error")
			fmt.Println(err)
			return 0
		}

		defer logFile.Close()
	} else {
		logFile = os.Stdout
	}

	loger := log.New(logFile, "hello ", log.Ldate|log.Ltime|log.Lshortfile)

	fmt.Printf("loger flags: %x", loger.Flags())

	loger.Printf("Hello, this a log message")
	loger.Output(1, "call depth 1")
	loger.Output(2, "call depth 2")

	if len(argv) > 0 && argv[0] == "-f" {
		loger.Fatal("I am fatal error!!!")
	}

	if len(argv) > 0 && argv[0] == "-p" {
		loger.Panic("I am panic error!!!")
	}

	return 0
}

func init() {
	command.AddCommand("log", "log package test", "Usage:\n log [-f|-p]", doLog)
}
