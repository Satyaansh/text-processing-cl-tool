package main

import (
	"fmt"
	"flag"
)

func main() {
	
	//decalring subcommand
	defCmd := flag.NewFlagSet("", flag.ExitOnError)


	//inputs for defCmd
	byteCount := defCmd.Bool("c", false, "Get number of bytes in file")
	lineCount := defCmd.Bool("l", false, "Get number of lines in file")
	wordCount := defCmd.Bool("w", false, "Get number of words in file")
	charCount := defCmd.Bool("m", false, "Get number of characters in file")
	byteLineWordCount := defCmd.Bool("", false, "Get number of bytes, lines, words in file")
}