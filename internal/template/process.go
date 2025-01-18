package template

import (
	"com/github/template-driven-development/tdd/internal/arguments"
	"log"
	"os"
	"text/template"
)

func Process(args arguments.Arguments) {
	tmpl, err := template.ParseFiles(args.Input)
	if err != nil {
		log.Fatalf("failed to parse template: %v", err)
	}

	outputFile, err := os.Create(args.Output)
	if err != nil {
		log.Fatalf("failed to create output file: %v", err)
	}
	defer outputFile.Close()

	if err := tmpl.Execute(outputFile, args.Data); err != nil {
		log.Fatalf("failed to execute template: %v", err)
	}
}
