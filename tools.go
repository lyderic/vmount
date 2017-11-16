package main

import (
	"os"
	"path/filepath"
)

func exists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) { // file not found
		return false
	}
	return true
}

func getFavoritesPath() string {
	return filepath.Join(os.Getenv("HOME"), ".config", "VeraCrypt", "Favorite Volumes.xml")
}
