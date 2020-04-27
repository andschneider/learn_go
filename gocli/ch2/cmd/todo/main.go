package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/andschneider/gocli/ch2/todo"
)

// Default file name
var todoFileName = ".todo.json"

// getTask function decides where to get the description for a new
// task from: arguments of STDIN
func getTask(r io.Reader, args ...string) (string, error) {
	// if arguments are present, create a task from them
	if len(args) > 0 {
		return strings.Join(args, " "), nil
	}
	// otherwise read from STDIN
	s := bufio.NewScanner(r)
	s.Scan()
	if err := s.Err(); err != nil {
		return "", err
	}
	if len(s.Text()) == 0 {
		return "", fmt.Errorf("Task cannot be blank")
	}
	return s.Text(), nil
}

func main() {
	// Display custom message for CLI
	flag.Usage = func() {
		fmt.Fprintln(flag.CommandLine.Output(), "Usage information:")
		flag.PrintDefaults()
	}
	// Parsing command line flags
	add := flag.Bool("add", false, "Add task to the ToDo list")
	list := flag.Bool("list", false, "List all tasks")
	finish := flag.Int("finish", 0, "Item to be completed")
	del := flag.Int("delete", 0, "Item to be deleted")
	verbose := flag.Bool("v", false, "Verbose output")
	open := flag.Bool("o", false, "Only show open tasks")
	flag.Parse()

	if os.Getenv("TODO_FILENAME") != "" {
		todoFileName = os.Getenv("TODO_FILENAME")
	}
	l := &todo.List{}
	if err := l.Get(todoFileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	switch {
	case *verbose || *list:
		l.Display(*verbose, *open)
	case *finish > 0:
		// Complete the given item
		if err := l.Complete(*finish); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		// Save the new list
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	case *add:
		// When any arguments (excluding flags) are provided, they will be
		// used as the new task
		t, err := getTask(os.Stdin, flag.Args()...)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		l.Add(t)

		// Save the new list
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	case *del > 0:
		// Delete the given item
		if err := l.Delete(*del); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		// Save the new list
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	default:
		fmt.Fprintln(os.Stderr, "Invalid option")
		os.Exit(1)
	}
}
