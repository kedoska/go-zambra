package command

import (
	"log"
	"os"
	"os/exec"
)

type Step struct {
	cmd *exec.Cmd
}

func Create(name string, args []string) (Step, error) {
	cmd := exec.Command(name, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return Step{
		cmd,
	}, nil
}

func Run(name string, args []string) {

	log.Printf("Running %v\n", name)

	s, err := Create(name, args)
	if err != nil {
		log.Printf("could not create the step: %s\n", err)
		os.Exit(1)
	}

	err = s.cmd.Run()
	if err != nil {
		log.Printf("could not run the command: %s\n", err)
		os.Exit(1)
	}
	log.Printf("Exiting %v\n", name)
}
