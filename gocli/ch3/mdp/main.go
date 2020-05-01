package main

import (
	"bytes"
	"flag"
	"fmt"
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
	header = `<!DOCTYPE html>
<html>
  <head>
    <meta http-equiv="content-type" content="text/html; charset=utf-8">
      <title>MarkdownPreviewTool</title>
  </head>
<body>
`

	footer = `
  </body>
</html>
`
)

func main() {
	filename := flag.String("file", "", "Markdown file to preview")
	skipPreview := flag.Bool("skip", false, "Skip auto-preview")
	flag.Parse()

	if *filename == "" {
		flag.Usage()
		os.Exit(1)
	}
	if err := run(*filename, os.Stdout, *skipPreview); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(filename string, out io.Writer, skipPreview bool) error {
	// Read all the data from the input file and check for errors
	input, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	htmlData := parseContent(input)

	temp, err := ioutil.TempFile("", "mdp")
	if err != nil {
		return err
	}
	if err := temp.Close(); err != nil {
		return err
	}
	outName := temp.Name() + ".html"
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

func parseContent(source []byte) []byte {
	md := goldmark.New(
		goldmark.WithExtensions(
			highlighting.Highlighting,
		),
		goldmark.WithRendererOptions(
			html.WithHardWraps(),
			html.WithXHTML(),
		))
	var buf bytes.Buffer
	buf.WriteString(header)
	md.Convert(source, &buf)
	return buf.Bytes()
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
