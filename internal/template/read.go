package template

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func Read(path string) string {
	if strings.HasPrefix(path, "http") || strings.HasPrefix(path, "https") {
		return readRemote(path)
	}

	return readLocal(path)
}

func readRemote(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("error fetching template from URL: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("received non-200 status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("error reading template response: %v", err)
	}

	return string(body)
}

func readLocal(path string) string {
	content, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("error reading local template file %s: %v", path, err)
	}

	return string(content)
}
