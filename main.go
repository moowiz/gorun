package main

import (
	"flag"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"text/template"
)

var prog = `
package main

import (
	"fmt"
  )
var _ = fmt.Print

func main() {
  {{.Code}};
}
`

func main() {
	flag.Parse()
	flags := flag.Args()
	command := flags[0]

	tmpl, err := template.New("prog").Parse(prog)
	if err != nil {
		panic(err)
	}

	dir, err := ioutil.TempDir("", "tempProg")
	if err != nil {
		panic(err)
	}
	f, err := os.Create(filepath.Join(dir, "main.go"))
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(f, map[string]interface{}{
		"Code": command,
	})
	if err != nil {
		panic(err)
	}
	cmd := exec.Command("go", "run", f.Name())
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		panic(err)
	}
}
