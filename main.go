package main

import (
	"com/github/tdd/tdd/internal/arguments"
	"com/github/tdd/tdd/internal/template"
	"log"
	"os"
)

const usage = "Usage: go run <path-to-config>"

func main() {
	if len(os.Args) != 2 {
		log.Fatal(usage)
	}

	argsPath := os.Args[1]
	if argsPath == "" {
		log.Fatal(usage)
	}

	args := arguments.FromFile(argsPath)
	for _, element := range args {
		template.Process(element)
	}
}
