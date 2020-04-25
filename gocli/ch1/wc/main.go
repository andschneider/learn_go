package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func count(r io.Reader, countLines bool, countBytes bool) int {
	// A scanner is used to read text from the a Reader (such as files)
	scanner := bufio.NewScanner(r)

	// Determine which type of parameter should be counted
	switch {
	case countBytes:
		scanner.Split(bufio.ScanBytes)
	case countLines:
		scanner.Split(bufio.ScanLines) // default operation of scan
	default:
		scanner.Split(bufio.ScanWords)
	}

	wc := 0
	for scanner.Scan() {
		wc++
	}
	return wc
}

func main() {
	// Defining a boolean flag -l to count lines instead of words
	lines := flag.Bool("l", false, "Count lines")
	// Defining a boolean flag -b to count bytes instead of words
	bytes := flag.Bool("b", false, "Count bytes")
	// Parsing the flags provided by the user
	flag.Parse()

	fmt.Println(count(os.Stdin, *lines, *bytes))
}
