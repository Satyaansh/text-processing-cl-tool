package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	
	defCmd := flag.NewFlagSet("ccwc", flag.ExitOnError)
	byteCount := defCmd.Bool("c", false, "Get number of bytes in file")
	lineCount := defCmd.Bool("l", false, "Get number of lines in file")
	wordCount := defCmd.Bool("w", false, "Get number of words in file")
	charCount := defCmd.Bool("m", false, "Get number of characters in file")
	byteLineWordCount := defCmd.Bool("", false, "Get number of bytes, lines, words in file")

	if len(os.Args) < 2 {
		fmt.Println("expected 'ccwc' command")
		os.Exit(1)
	}

	switch os.Args[1] {
		case "ccwc": // if its the 'get' command
			HandleCcwc()
		default: // if we don't understand the input

	}

} 