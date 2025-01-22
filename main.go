package main

import (
	"log"
	"os"

	"github.com/template-driven-development/tdd/internal/arguments"
	"github.com/template-driven-development/tdd/internal/template"
)

const usage = "Usage: tdd <path-to-config>"

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
