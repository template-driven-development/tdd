package template

import (
	"encoding/base64"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func Read(path string) string {
	if !isRemote(path) {
		return readLocal(path)
	}

	content := readCache(path)
	if content != "" {
		return content
	}

	content = readRemote(path)
	writeCache(path, content)

	return content
}

func cachePath(path string) string {
	userCacheDir, err := os.UserCacheDir()
	if err != nil {
		log.Printf("error getting user cache directory: %v\n", err)
		return ""
	}

	encodedPath := base64.RawURLEncoding.EncodeToString([]byte(path))

	return filepath.Join(userCacheDir, "tdd", "templates", encodedPath)
}

func readCache(path string) string {
	cachePath := cachePath(path)
	if cachePath == "" {
		return ""
	}

	content, err := os.ReadFile(cachePath)
	if err != nil {
		return ""
	}

	log.Printf("cache read successfully: %s, %s\n", path, cachePath)
	return string(content)
}

func writeCache(path, content string) {
	cachePath := cachePath(path)
	if cachePath == "" {
		return
	}

	cacheDir := filepath.Dir(cachePath)

	err := os.MkdirAll(cacheDir, os.ModePerm)
	if err != nil {
		log.Printf("error creating cache directory %s: %v\n", cacheDir, err)
		return
	}

	err = os.WriteFile(cachePath, []byte(content), 0644)
	if err != nil {
		log.Printf("error writing to cache file %s: %v\n", cachePath, err)
		return
	}

	log.Printf("cache written successfully: %s, %s\n", path, cachePath)
}

func isRemote(path string) bool {
	return strings.HasPrefix(path, "http") || strings.HasPrefix(path, "https")
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

	log.Printf("remote read successfully: %s\n", url)
	return string(body)
}

func readLocal(path string) string {
	content, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("error reading local template file %s: %v", path, err)
	}

	log.Printf("local read successfully: %s\n", path)
	return string(content)
}
