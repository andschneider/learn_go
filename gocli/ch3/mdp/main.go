package main

import (
	"bytes"
	"flag"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"time"

	"github.com/yuin/goldmark"
	highlighting "github.com/yuin/goldmark-highlighting"
	"github.com/yuin/goldmark/renderer/html"
)

const (
	defaultTemplate = `<!DOCTYPE html>
<html>
  <head>
    <meta http-equiv="content-type" content="text/html; charset=utf-8">
      <title>{{ .Title }}</title>
  </head>
  <body>
  {{ .Body }}
  </body>
</html>
`
)

// content type represents the HTML content to add into the template
type content struct {
	Title string
	Body  template.HTML
}

func main() {
	filename := flag.String("file", "", "Markdown file to preview")
	skipPreview := flag.Bool("skip", false, "Skip auto-preview")
	tFname := flag.String("t", "", "Alternate template name")
	flag.Parse()

	if *filename == "" {
		flag.Usage()
		os.Exit(1)
	}
	if err := run(*filename, *tFname, os.Stdout, *skipPreview); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(filename string, tFname string, out io.Writer, skipPreview bool) error {
	// Read all the data from the input file and check for errors
	input, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	htmlData, err := parseContent(input, tFname)
	if err != nil {
		return err
	}

	// Create temporary file
	temp, err := ioutil.TempFile("", "mdp.*.html")
	if err != nil {
		return err
	}
	if err := temp.Close(); err != nil {
		return err
	}
	outName := temp.Name()
	fmt.Fprintln(out, outName)

	if err := saveHTML(outName, htmlData); err != nil {
		return err
	}

	if skipPreview {
		return nil
	}

	defer os.Remove(outName)
	return preview(outName)
}

func parseContent(source []byte, tFname string) ([]byte, error) {
	// Convert markdown to HTML
	md := goldmark.New(
		goldmark.WithExtensions(
			highlighting.Highlighting,
		),
		goldmark.WithRendererOptions(
			html.WithHardWraps(),
			html.WithXHTML(),
		))
	var con bytes.Buffer
	if err := md.Convert(source, &con); err != nil {
		return nil, err
	}

	// Parse the contents of the defaultTemplate const into a new Template
	t, err := template.New("mdp").Parse(defaultTemplate)
	if err != nil {
		return nil, err
	}

	// If user provided alternate template file, replace template
	if tFname != "" {
		t, err = template.ParseFiles(tFname)
		if err != nil {
			return nil, err
		}
	}

	// Add markdown to template
	c := content{
		Title: "Markdown Preview Tool",
		Body:  template.HTML(con.String()),
	}
	var buf bytes.Buffer
	if err := t.Execute(&buf, c); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func preview(fname string) error {
	// Locate Chrome in the PATH
	// browserPath, err := exec.LookPath("firefox")
	browserPath, err := exec.LookPath("google-chrome")
	if err != nil {
		return err
	}

	// Open the file in the browser
	if err := exec.Command(browserPath, fname).Start(); err != nil {
		return err
	}

	// Give the browser some time to open the file before deleting it
	time.Sleep(2 * time.Second)
	return nil
}

func saveHTML(outFname string, data []byte) error {
	return ioutil.WriteFile(outFname, data, 0644)
}
