package main

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func initConfig() {
	// check and set veracrypt executable
	veracryptBinaryPath, err := exec.LookPath("veracrypt")
	if err != nil {
		log.Fatalln("Veracrypt executable not found!")
	}
	config.VeracryptBinaryPath = veracryptBinaryPath
	// check that there is an editor we can use for edit action (-e)
	editor := os.Getenv("EDITOR")
	if len(editor) == 0 {
		editor = "vi"
	}
	editorPath, err := exec.LookPath(editor)
	if err != nil {
		log.Fatalln("Editor not found!")
	}
	config.EditorPath = editorPath
	// Path to VeraCrypt's 'Favorite Volumes.xml' file
	config.FavoritesPath = getFavoritesPath()
}

func exists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) { // file not found
		return false
	}
	return true
}

func getFavoritesPath() string {
	return filepath.Join(os.Getenv("HOME"), ".config", "VeraCrypt", "Favorite Volumes.xml")
}
