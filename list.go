package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)

func list() {
	var veraCrypt VeraCrypt
	file, err := os.Open(config.FavoritesPath)
	if err != nil {
		log.Println("Maybe you should run: 'vmount -e' to edit the favorite config file")
		log.Fatal(err)
	}
	defer file.Close()
	byteval, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	xml.Unmarshal(byteval, &veraCrypt)
	volumes := veraCrypt.Favorites.Volumes
	longestPath := 0
	replacer := strings.NewReplacer(os.Getenv("HOME"), "~")
	for i := 0; i < len(volumes); i++ {
		volumes[i].ShortPath = replacer.Replace(volumes[i].Path)
		volumes[i].ShortMountpoint = replacer.Replace(volumes[i].Mountpoint)
		if len(volumes[i].ShortPath) > longestPath {
			longestPath = len(volumes[i].ShortPath)
		}
	}
	format := setFormat(longestPath)
	for _, volume := range volumes {
		if !exists(volume.Path) {
			log.Fatalln("ERROR: volume", volume.Path, "not found!")
		}
		if !exists(volume.Mountpoint) {
			log.Fatalln("ERROR: mountpoint", volume.Mountpoint, "not found!")
		}
		err := exec.Command("veracrypt", "--text", "--list",
			fmt.Sprintf("--slot=%d", volume.Slotnumber)).Run()
		mounted := "[*]"
		if err != nil {
			mounted = "[ ]"
		}
		fmt.Printf(format, volume.Slotnumber, mounted,
			volume.ShortPath, volume.ShortMountpoint)
	}
}
