package main

import (
	"fmt"
	"os"
)

const (
	ExitCodeSuccess = 0
	ExitCodeError   = 1
)

type CommandHandler interface {
	Handle(args []string) error
	Help()
}

func main() {
	var handlers = map[string]CommandHandler{
		"init":    InitHandler{},
		"info":    InfoHandler{},
		"package": PackageHandler{},
	}

	// Strip the first rgument out since it is the binary name.
	args := os.Args[1:]

	if len(args) == 0 {
		os.Exit(ExitCodeError)
	}

	command := args[0]
	tail := args[1:]

	if len(tail) == 1 && tail[0] == "help" {
		handlers[command].Help()
		os.Exit(ExitCodeSuccess)
	}

	handler := handlers[command]
	err := handler.Handle(tail)
	if err != nil {
		handler.Help()
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(ExitCodeError)
	}

	os.Exit(ExitCodeSuccess)
}
