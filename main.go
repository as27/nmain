package main

import (
	"flag"
	"fmt"
	"os"
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

}

const mfile = `package main

func main() {

}`
