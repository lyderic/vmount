package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func main() {
	ln := len(os.Args)
	switch ln {
	case 0:
		log.Fatal("ERROR")
	case 1:
		list()
	default:
		switch os.Args[1] {
		case "-d":
			switch ln {
			case 2:
				dismountAll()
			case 3:
				slot, err := strconv.Atoi(os.Args[2])
				if err != nil {
					log.Fatal(err)
				}
				dismountSlot(slot)
			default:
				log.Fatal("invalid argument to -d switch")
			}
		case "-m":
			mountFavorites()
			list()
		case "-l":
			list()
		default:
			log.Fatal("invalid switch!")
		}
	}
}

func dismountAll() {
	veracrypt("--verbose", "--dismount")
}

func dismountSlot(slot int) {
	veracrypt("--verbose", "--dismount", fmt.Sprintf("--slot=%d", slot))
}

func mountFavorites() {
	veracrypt("--auto-mount=favorites", "--pim=0", "--keyfiles=", "--protect-hidden=no", "--verbose")
}

func veracrypt(args ...string) {
	args = append([]string{"--text"}, args...)
	cmd := exec.Command("veracrypt", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Run()
}

type VeraCrypt struct {
	XMLName   xml.Name  `xml:"VeraCrypt"`
	Favorites Favorites `xml:"favorites"`
}

type Favorites struct {
	XMLName xml.Name `xml:"favorites"`
	Volumes []Volume `xml:"volume"`
}

type Volume struct {
	XMLName         xml.Name `xml:"volume"`
	Path            string   `xml:",chardata"`
	ShortPath       string
	Mountpoint      string `xml:"mountpoint,attr"`
	ShortMountpoint string
	Readonly        int `xml:"readonly,attr"`
	Slotnumber      int `xml:"slotnumber,attr"`
	System          int `xml:"system,attr"`
}

func list() {
	var veraCrypt VeraCrypt
	file, err := os.Open(os.Getenv("HOME") + "/.config/VeraCrypt/Favorite Volumes.xml")
	if err != nil {
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

func setFormat(ln int) string {
	return fmt.Sprintf("%%02d  %%s  %%-%d.%ds  %%s\n", ln, ln)

}
