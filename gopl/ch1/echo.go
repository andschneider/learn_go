package main

import (
	"fmt"
	"os"
	"strings"
)

func echoBase() {
	fmt.Println("-- base --")
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

// exercise 1.1
// Modify the echo program to also print os.Args[0]
func echo11() {
	fmt.Println("-- exercise 1.1 --")
	f := os.Args[0]
	fmt.Println(f)
}

// exercise 1.2
// Modify the echo program to print the index and value of each of its arguments
// Print one per line
func echo12() {
	fmt.Println("-- exercise 1.2 --")
	for i, arg := range os.Args[1:] {
		fmt.Println(i, arg)
	}
}

// exercise 1.3
// Experiment to measure the difference in running time between our potentially 
// inefficient versions and the one that uses strings.Join.
func echoJoin () {
	fmt.Println(strings.Join(os.Args[1:], " "))
}

func main() {
	echoBase()
	echo11()
	echo12()
}
