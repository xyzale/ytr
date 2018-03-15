package parser

import (
	"os"
)

var fileContent []byte

var yamlContent = Yaml{}

type Yaml struct {
	Commands map[string][]string   `yaml:"commands"`
}

func (m Yaml) findCommand(cmd string) []string {
	if !m.hasCommand(cmd) {
		os.Stderr.WriteString("Error: Command not found")
		os.Exit(1)
	}

	return yamlContent.Commands[cmd]
}

func (m Yaml) hasCommand(cmd string) bool {
	for command := range m.Commands {
		if command == cmd {
			return true
		}
	}
	return false
}
