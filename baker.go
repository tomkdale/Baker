package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
)

func printHelp() {
	helpString := `bak helper file

 -h shows help page.
 -c file copies file to file.bak or vise-versa.
 -d file shows diff on file and file.bak
`
	fmt.Println(helpString)
	return

}

func getFlags() (bool, bool) {
	helpPtr := *flag.Bool("h", false, "help")
	copyPtr := *flag.Bool("c", false, "Copy files/directories rather than move them")
	diffPtr := *flag.Bool("d", false, "Runs diff on file and file.bak")
	flag.Parse()
	// If no variables passed in show help screen and exit 2

	if helpPtr {
		printHelp()
		os.Exit(0)
	}
	return copyPtr, diffPtr
}
func getToFrom() (string, string) {
	if len(os.Args) <= 1 {
		printHelp()
		os.Exit(2)
	}
	fileName := os.Args[1]
	var filePath string
	//check if file path is relative or absolute
	isRelative := os.IsPathSeparator(os.Args[1][0])
	if isRelative {
		filePath = fileName
	} else {
		pwd, err := os.Getwd()
		if err != nil {
			fmt.Print("Could not get current working directory")
			os.Exit(2)
		}
		filePath = pwd + "/" + fileName
	}
	isBaker := strings.HasSuffix(filePath, ".bak")
	var to string
	var from string
	if isBaker {
		to = strings.TrimSuffix(filePath, ".bak")
		from = filePath
	} else {
		to = filePath + ".bak"
		from = filePath
	}
	return to, from
}
func exists(file string) bool {
	if _, err := os.Stat(file); errors.Is(err, os.ErrNotExist) {
		return false
	} else {
		return true
	}
}
func main() {
	copyPtr, diffPtr := getFlags()

	to, from := getToFrom()

	if exists(from) != true {
		fmt.Printf("%s does not exist \n", from)
		os.Exit(1)
	}
	if exists(to) {
		tmpFile := to + ".tmp"
		os.Rename(from, tmpFile)
		os.Rename(to, from)
		os.Rename(tmpFile, to)
		fmt.Println("Swapped " + from + " with " + to)
	} else {
		os.Rename(from, to)
		fmt.Println("Copied" + from + " to " + to)
	}

	_ = copyPtr
	_ = diffPtr
}
