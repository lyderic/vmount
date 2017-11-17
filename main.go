package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
)

func init() {
	log.SetFlags(log.Lshortfile)
	initConfig()
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
		case "-e":
			edit()
		case "-version":
			fmt.Println(version)
			return
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
	cmd := exec.Command(config.VeracryptBinaryPath, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Run()
}

func edit() {
	if !exists(config.FavoritesPath) { // file not found
		f, err := os.Create(config.FavoritesPath)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		f.WriteString(template)
	}
	cmd := exec.Command(config.EditorPath, config.FavoritesPath)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

func setFormat(ln int) string {
	return fmt.Sprintf("%%02d  %%s  %%-%d.%ds  %%s\n", ln, ln)
}
