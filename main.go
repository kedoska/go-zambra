package main

import (
	"log"
	"os"

	"github.com/kedoska/go-zabra/cmd/command"
	"github.com/kedoska/go-zabra/cmd/workflow"
)

// go run main.go run <cmd> <args>
func main() {
	switch os.Args[1] {
	case "run":
		name := os.Args[2]
		args := os.Args[3:]
		command.Run(name, args)
	case "execute":
		if len(os.Args) < 3 {
			log.Printf("expected filename to execute")
			os.Exit(1)
		}
		filename := os.Args[2]
		workflow.ExecuteFromFile(filename)
	default:
		panic("help")
	}
}
