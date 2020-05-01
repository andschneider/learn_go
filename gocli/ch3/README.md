# mdp

A simple markdown previewer that converts markdown to html and displays in a
browser.

## Usage

To run the CLI, use the `-file` flag followed by a path to a markdown file. If
you don't want to automatically display the html in a browser add the `-skip`
flag.

For example, to display this README run the following from inside the mdp
directory.

```bash
./mdp -file ../README.md
```

## Building

To build:

```bash
make build
```

To run the tests:

```bash
make test
```

