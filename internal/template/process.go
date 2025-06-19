package template

import (
	"log"
	"os"
	"text/template"

	"github.com/template-driven-development/tdd/internal/arguments"
)

func Process(args arguments.Arguments) {
	text := Read(args.Input)

	tmpl, err := template.New("template").Funcs(Functions).Parse(text)
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

	log.Printf("template written successfully: %s\n", args.Output)
}
