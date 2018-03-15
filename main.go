package main

import (
	"./parser"
	"os"
)

func main() {
	parser.ExecuteCommand(os.Args[1])
}
