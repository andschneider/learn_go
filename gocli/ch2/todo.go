package todo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"text/tabwriter"
	"time"
)

type item struct {
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

type List []item

// Add creates a new ToDo item and appends it to the list
func (l *List) Add(task string) {
	t := item{
		Task:        task,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}
	*l = append(*l, t)
}

// Complete method marks a ToDo item as completed by
// settings Done = true and CompletedAt to the current time
func (l *List) Complete(i int) error {
	ls := *l
	if i <= 0 || i > len(ls) {
		return fmt.Errorf("item %d does not exist", i)
	}

	// Adjusting index for 0 based index
	ls[i-1].Done = true
	ls[i-1].CompletedAt = time.Now()

	return nil
}

// Delete method deletes a ToDo item from the list
func (l *List) Delete(i int) error {
	ls := *l
	if i <= 0 || i > len(ls) {
		return fmt.Errorf("item %d does not exist", i)
	}

	// Adjusting index for 0 based index
	*l = append(ls[:i-1], ls[i:]...)

	return nil
}

// Save method encodes the List as JSON and saves it
// using the provided file name
func (l *List) Save(filename string) error {
	js, err := json.Marshal(l)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, js, 0644)
}

// Get method opens the provided file name, decodes
// the JSON data and parses it into a List
func (l *List) Get(filename string) error {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	if len(file) == 0 {
		return nil
	}
	return json.Unmarshal(file, l)
}

// String prints out a formatted list
// Implements the fmt.Stringer interface
func (l *List) String() string {
	var formatted strings.Builder
	for k, t := range *l {
		prefix := "  "
		if t.Done {
			prefix = "X "
		}
		// Adjust the item number k to print numbers starting from 1
		fmt.Fprintf(&formatted, "%s%d: %s\n", prefix, k+1, t.Task)
	}
	return formatted.String()
}

// General display function to print tasks
// v is verbose which will display the created date
// o is open which will only display tasks that aren't completed
// Uses tabwriter to display output like a table
func (l *List) Display(v bool, o bool) {
	format := "Jan 2 15:04"
	w := tabwriter.NewWriter(os.Stdout, 8, 8, 1, '\t', 0)
	for k, t := range *l {
		prefix := "  "
		if t.Done {
			if o {
				// hide complete
				continue
			}
			prefix = "X "
		}
		if v {
			date := t.CreatedAt.Format(format) // format time to string
			fmt.Fprintf(w, "%s%d: %s\t%s\n", prefix, k+1, t.Task, date)
		} else {
			fmt.Fprintf(w, "%s%d: %s\n", prefix, k+1, t.Task)
		}
	}
	w.Flush()
}
