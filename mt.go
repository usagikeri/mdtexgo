package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
	"text/template"

	_ "./statik"
	"github.com/rakyll/statik/fs"
)

var (
	tpl      *template.Template
	fileName string
	buff     *bytes.Buffer
	fw       io.Writer
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

	buff = new(bytes.Buffer)
	fw = io.Writer(buff)
}

func runPandoc(mdfile string) string {
	out, err := exec.Command("pandoc", mdfile, "-t", "latex").Output()

	if err != nil {
		log.Fatal(err)
	}
	return string(out)
}

func cleanup(texfile string) string {
	data, err := ioutil.ReadFile(texfile)
	if err != nil {
		panic(err)
	}
	str := string(data)
	temp := strings.Replace(str, "\\begin{varbatim}", "\\begin{lstlisting}", -1)
	temp = strings.Replace(str, "\\end{varbatim}", "\\end{lstlisting}", -1)
	return temp
}

func main() {
	texText := runPandoc(fileName)
	outfileName := strings.Split(fileName, ".")[0] + ".tex"

	tpl.Execute(fw, struct {
		Text string
	}{
		Text: texText,
	})
	ioutil.WriteFile(outfileName, buff.Bytes(), os.ModePerm)
	fmt.Println("Create " + outfileName)
}
