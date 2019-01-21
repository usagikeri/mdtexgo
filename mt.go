package main

import (
	_ "./statik"
	"flag"
	"fmt"
	"github.com/rakyll/statik/fs"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
	"text/template"
)

var (
	tpl      *template.Template
	fileName string
)

func init() {
	statikFS, err := fs.New()
	if err != nil {
		panic(err)
	}

	f, err := statikFS.Open("/template.tex")
	if err != nil {
		panic(err)
	}

	b, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}

	s := string(b)

	tpl = template.Must(template.New("").Parse(s))
	flag.Parse()
	fileName = flag.Arg(0)
}

func runPandoc(mdfile string) string {
	out, err := exec.Command("pandoc", mdfile, "-t", "latex").Output()

	if err != nil {
		log.Fatal(err)
	}
	return string(out)
}

func main() {
	texText := runPandoc(fileName)
	outfileName := strings.Split(fileName, ".")[0] + ".tex"

	file, err := os.create(outfilename)
	if err != nil {
		log.fatal(err)
	}
	defer file.close()

	tpl.Execute(file, struct {
		Text string
	}{
		Text: texText,
	})

	fmt.Println("Create " + outfileName)
}
