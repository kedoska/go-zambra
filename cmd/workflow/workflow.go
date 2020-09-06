package workflow

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/kedoska/go-zabra/cmd/command"
	"gopkg.in/yaml.v2"
)

type Definition struct {
	Name  string `yaml:"name"`
	Steps []struct {
		Name string   `yaml:"name"`
		Args []string `yaml:"args"`
	} `yaml:"steps"`
}

func ExecuteFromFile(filename string) {
	yamlFile, err := ioutil.ReadFile(filename)

	if err != nil {
		log.Printf("%s\n", err)
		os.Exit(1)
	}

	var d = Definition{}

	err = yaml.Unmarshal(yamlFile, &d)
	if err != nil {
		panic(err)
	}

	for i := 0; i < len(d.Steps); i++ {
		s := d.Steps[i]
		command.Run(s.Name, s.Args)
	}
}
