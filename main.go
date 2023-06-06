package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var Version int = 1

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Missing arguments. See command `help` for more information.")
	}

	command := os.Args[1]
	options := os.Args[2:]

	if command == "help" || command == "-h" || command == "--help" {
		usage()
	} else if command == "version" || command == "--version" || command == "-V" {
		fmt.Printf("rbook v%v\n", Version)
	} else if command == "r" || command == "rename" {
		if len(options) < 2 {
			log.Fatal("Missing arguments. Must provide input_file, title and author.")
		}

		inputFile := options[0]
		title := options[1]
		author := options[2]
		tags := options[3:]
		extension := filepath.Ext(inputFile)

		fname := createFileName(title, author, tags) + extension
		err := os.Rename(inputFile, fname)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		log.Fatal("Invalid command", command, "See command `help` for more information.")
	}
}

func createFileName(title string, author string, tags []string) string {
	excludedPunctuationRegexp := regexp.MustCompile("[][{}!@#$%^&*()=+'\"?,.|;:~`‘’“”/]*")
	var result string

	lowercasedTitle := strings.Trim(strings.ToLower(title), " ")
	lowercasedAuthor := strings.Trim(strings.ToLower(author), " ")

	result = excludedPunctuationRegexp.ReplaceAllString(lowercasedTitle, "")
	result += "__" + excludedPunctuationRegexp.ReplaceAllString(lowercasedAuthor, "")

	for _, tag := range tags {
		lowercasedTag := strings.Trim(strings.ToLower(tag), " ")
		sanitizedTag := excludedPunctuationRegexp.ReplaceAllString(lowercasedTag, "")
		result += "_" + sanitizedTag
	}

	return strings.ReplaceAll(result, " ", "-")
}

func usage() {
	fmt.Println("USAGE: rbook <command>.")
	fmt.Println("")
	fmt.Println("COMMANDS")
	fmt.Println("--------")
	fmt.Println("[r]ename <input_file> [options] [tags]: rename input_file; no options = interactive mode")
	fmt.Println("help: print usage information (this)")
	fmt.Println("version: print version")
}
