package parser

import (
	"os"
	"fmt"
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"log"
	"os/exec"
	"bytes"
)

func loadFileContents() {
	if _, err := os.Stat("./Ytrfile.yaml"); os.IsNotExist(err) {
		os.Stderr.WriteString("Error: Ytrfile missing")
		os.Exit(1)
	}

	content, err := ioutil.ReadFile("./Ytrfile.yaml")

	if err != nil {
		os.Stderr.WriteString("Error: Unable to read Ytrfile")
		os.Exit(1)
	}

	fileContent = content
}

func parseFile() {
	loadFileContents()

	err := yaml.Unmarshal([]byte(fileContent), &yamlContent)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
}

func ExecuteCommand(cmd string) {
	parseFile()

	if !stringInMap(cmd, yamlContent.Commands) {
		os.Stderr.WriteString("Error: Command not found")
		os.Exit(1)
	}

	for key, value := range yamlContent.Commands {
		if key == cmd {
			for _, command := range value {
				instruction := exec.Command(command)
				cmdOutput := &bytes.Buffer{}
				instruction.Stdout = cmdOutput
				err := instruction.Run()
				if err != nil {
					os.Stderr.WriteString(err.Error())
				}
				fmt.Print(string(cmdOutput.Bytes()))
			}
		}
	}
}

func stringInMap(a string, list map[string][]string) bool {
	for b := range list {
		if b == a {
			return true
		}
	}
	return false
}
