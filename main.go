package main

import (
	"fmt"
	"log"
	"sort"
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

	if os.Args[1] == "help" {
		usage()
		return
	}

	options := os.Args[1:]

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
}

func createFileName(title string, author string, tags []string) string {
	sort.Strings(tags)
	excludedPunctuationRegexp := regexp.MustCompile("[][{}!@#$%^&*()=+'\"?,.|;:~`‘’“”/]*")

	sanitizedTitle := strings.Trim(strings.ToLower(title), " ")
	sanitizedAuthor := strings.Trim(strings.ToLower(author), " ")

	result := excludedPunctuationRegexp.ReplaceAllString(sanitizedTitle, "") +
		"__" + excludedPunctuationRegexp.ReplaceAllString(sanitizedAuthor, "")

	for _, tag := range tags {
		trimmedTag := strings.Trim(tag, " ")
		sanitizedTag := excludedPunctuationRegexp.ReplaceAllString(trimmedTag, "")
		result += "_" + strings.ToLower(sanitizedTag)
	}

	return strings.ReplaceAll(result, " ", "-")
}

func usage() {
	fmt.Println("USAGE: rb <input_file|help|>")
	fmt.Println("")
}
