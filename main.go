package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	release := flag.Bool("r", false, "Print release number")
	word := flag.Bool("w", false, "Print release number as words. Implicates -r")
	full := flag.Bool("a", true, "Print full release name and number. Default.")

	dat, err := os.ReadFile("/etc/fedora-release")

	if err != nil {
		log.Fatal("Couldn't open release file. Is it Fedora?")
	}

	words := strings.Fields(string(dat))

	flag.Parse()
	if *release {
		*full = false
		fmt.Println(words[2])
		return
	} else if *word {
		*full = false
		fmt.Println(words[3] + " " + words[4])
	} else {
		fmt.Print(string(dat))
	}
}
