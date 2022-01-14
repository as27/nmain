package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

var (
	flagFileName  = flag.String("file", "main.go", "change the name of the generated file")
	flagOverwrite = flag.Bool("force", false, "overwrite existing files")
)

func main() {
	flag.Parse()

	fileName := *flagFileName
	force := *flagOverwrite
	fileExists := false
	f, err := os.Open(fileName)
	if err == nil {
		fileExists = true
		f.Close()
	}
	if !force && fileExists {
		fmt.Println("file already exists!")
		fmt.Println("nothing is done here.")
		os.Exit(0)
	} else {
		os.RemoveAll(fileName)
	}
	f, err = os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		fmt.Println("error OpenFile: ", err)
		os.Exit(1)
	}
	r := strings.NewReader(mfile)
	io.Copy(f, r)
	f.Close()
}

const mfile = `package main

func main() {

}`
