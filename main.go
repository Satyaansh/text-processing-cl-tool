package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	
	//inputs for defCmd
	byteCount := flag.Bool("c", false, "Get number of bytes in file")
	lineCount := flag.Bool("l", false, "Get number of lines in file")
	wordCount := flag.Bool("w", false, "Get number of words in file")
	charCount := flag.Bool("m", false, "Get number of characters in file")
	byteLineWordCount := flag.Bool("", false, "Get number of bytes, lines, words in file")

	if len(os.Args) < 1 {
		fmt.Println("missing file name")
		os.Exit(1)
	}
} 