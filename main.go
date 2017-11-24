package main

import (
	"flag"
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
	if ln == 1 { // default action
		list()
		return
	}
	var doList, doEdit, doMount, showVersion bool
	var slotNumberString string
	flag.BoolVar(&doList, "l", false, "list favorites")
	flag.BoolVar(&doMount, "m", false, "mount all favorites")
	flag.StringVar(&slotNumberString, "d", "--unset--", "dismount slot # (0=all)")
	flag.BoolVar(&doEdit, "e", false, "edit favorites XML configuration")
	flag.BoolVar(&showVersion, "version", false, "show version")
	flag.Parse()

	if doList {
		list()
		return
	}
	if doMount {
		mountFavorites()
		list()
		return
	}
	if doEdit {
		edit()
		return
	}
	if slotNumberString != "--unset--" {
    log.Println(">>>", flag.Args(), flag.NArg())
		slotNumber, err := strconv.Atoi(slotNumberString)
		if err != nil {
			log.Fatal(err)
		}
		if slotNumber == 0 {
			dismountAll()
		} else {
			dismountSlot(slotNumber)
		}
    return
	}
  if showVersion {
    fmt.Println(version)
    return
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
