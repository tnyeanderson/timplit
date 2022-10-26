package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"text/template"
)

const (
	templateName = "timplit"
)

type arguments struct {
	invert bool
	path   string
}

func getArgs() (a arguments) {
	if len(os.Args) == 1 {
		panic("Missing required parameter: filename")
	}
	a.path = os.Args[1]
	if a.path == "-j" {
		a.invert = true
		a.path = os.Args[2]
	}
	return
}

func readStdin() []byte {
	return readAll(os.Stdin)
}

func readFile(path string) []byte {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	return readAll(f)
}

func readAll(f *os.File) []byte {
	b, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}
	return b
}

// Arrays won't unmarshal properly. Wrap them in an object under the "items"
// property.
func wrapJsonArray(b []byte) []byte {
	if bytes.HasPrefix(b, []byte("[")) {
		return []byte(fmt.Sprintf(`{"items": %s}`, b))
	}
	return b
}

func parseJson(b []byte) (data map[string]interface{}) {
	b = wrapJsonArray(b)
	if err := json.Unmarshal(b, &data); err != nil {
		panic(err)
	}
	return
}

func makeTemplate(templateString string) *template.Template {
	return template.Must(template.New(templateName).Parse(templateString))
}

func executeTemplate(t *template.Template, d map[string]interface{}) {
	if err := t.ExecuteTemplate(os.Stdout, templateName, d); err != nil {
		panic(err)
	}
}

func timplit(templateString []byte, data []byte) {
	t := makeTemplate(fmt.Sprintf("%s", templateString))
	d := parseJson(data)
	executeTemplate(t, d)
}

func main() {
	args := getArgs()
	fromStdin := readStdin()
	fromFile := readFile(args.path)
	if args.invert {
		timplit(fromStdin, fromFile)
	} else {
		timplit(fromFile, fromStdin)
	}
}
