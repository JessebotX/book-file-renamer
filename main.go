package main

import (
	"fmt"
	"log"
	"os"
)

var Version int = 1

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Missing arguments. See command `help` for more information.")
	}

	command := os.Args[1]

	if command == "help" || command == "-h" || command == "--help" {
		usage()
	} else if command == "version" || command == "--version" || command == "-V" {
		fmt.Printf("rbook v%v\n", Version)
	} else {
		log.Fatal("Invalid command", command, "See command `help` for more information.")
	}
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
