package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func count(i []byte, countLines bool, countBytes bool, countWords bool) int {
	r := strings.NewReader(string(i))
	scanner := bufio.NewScanner(r)

	switch {
	case countLines:
		scanner.Split(bufio.ScanLines)
	case countBytes:
		scanner.Split(bufio.ScanBytes)
	case countWords:
		scanner.Split(bufio.ScanWords)
	}

	wc := 0
	for scanner.Scan() {
		wc++
	}
	return wc
}

func main() {
	lines := flag.Bool("l", false, "Count lines")
	words := flag.Bool("w", false, "Count words")
	bytes := flag.Bool("b", false, "Count bytes")
	flag.Parse()
	i, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not parse input: %v\n", err)
	}

	// Default is to count all types
	if !*lines && !*bytes && !*words {
		l := count(i, true, *bytes, *words)
		b := count(i, *lines, true, *words)
		w := count(i, *lines, *bytes, true)
		fmt.Printf("\t%d\t%d\t%d\n", l, w, b)
		return
	}

	// Count specific type
	wc := count(i, *lines, *bytes, *words)
	fmt.Println(wc)
}
