package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type Extensions struct {
	Extensions []string `json:"extensions"`
}

func install(path string) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	defer file.Close()

	var extensions Extensions
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&extensions)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}

	installedExtensions, err := get()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}

	for _, ext := range extensions.Extensions {
		if isInstalled(ext, installedExtensions) {
			fmt.Printf("Extension %s is already installed, skipping.\n", ext)
			continue
		}

		fmt.Printf("Installing %s...\n", ext)
		cmd := exec.Command("code", "--install-extension", ext)
		err := cmd.Run()
		if err != nil {
			fmt.Printf("Failed to install %s: %s\n", ext, err)
		} else {
			fmt.Printf("Successfully installed %s\n", ext)
		}
	}
}

func get() ([]string, error) {
	cmd := exec.Command("code", "--list-extensions")
	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("failed to get installed extensions: %s", err)
	}

	extensions := strings.Split(string(output), "\n")
	var installedExtensions []string
	for _, ext := range extensions {
		if ext != "" {
			installedExtensions = append(installedExtensions, ext)
		}
	}
	return installedExtensions, nil
}

func isInstalled(ext string, installedExtensions []string) bool {
	for _, installed := range installedExtensions {
		if installed == ext {
			return true
		}
	}
	return false
}

func main() {
	install("extensions.json")
}
