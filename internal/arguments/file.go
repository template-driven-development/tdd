package arguments

import (
	"encoding/json"
	"log"
	"os"
)

func FromFile(path string) []Arguments {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("failed to open config file: %v", err)
	}
	defer file.Close()

	var configs []Arguments
	if err := json.NewDecoder(file).Decode(&configs); err != nil {
		log.Fatalf("failed to parse config file: %v", err)
	}

	return configs
}
