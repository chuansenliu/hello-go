package main

import (
	"os"

	"github.com/eehut/hello-go/command"
	"github.com/sirupsen/logrus"
)

// doLogrus is command function
func doLogrus(argv []string) int {

	//logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.DebugLevel)

	logrus.WithFields(logrus.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")

	logrus.WithFields(logrus.Fields{
		"omg":    true,
		"number": 122,
	}).Warn("The group's number increased tremendously!")

	logrus.WithFields(logrus.Fields{
		"omg":    true,
		"number": 100,
	}).Debug("The ice breaks!")

	logrus.Debug("This is debug message")

	return 0
}

func init() {
	command.AddCommand("logrus", "logrus package test", "Usage:\n logrus", doLogrus)
}
