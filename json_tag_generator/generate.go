package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/aman-bansal/go_json_tag_generator/json_tag_generator/internal"
	"go/format"
	"io/ioutil"
	"log"
)

var (
	sourceFile = flag.String("source", "", "(source mode) Input Go source file; enables source mode.")
)

func main() {
	flag.Parse()
	fmt.Println("starting generate")
	if *sourceFile == "" {
		log.Fatal("Source file is empty")
	}

	fileset, file, err := internal.ParseFile(*sourceFile)
	if err != nil {
		log.Fatalf("Loading input failed: %v", err)
	}

	internal.FormatCode(fileset, file)

	buf := new(bytes.Buffer)
	err = format.Node(buf, fileset, file)
	if err != nil {
		log.Fatalf("error: %v\n", err)
	}
	byt := buf.Bytes()
	err = ioutil.WriteFile(*sourceFile, byt, 0664)
	if err != nil {
		log.Fatalf("error: %v\n", err)
	}

	fmt.Println("json tag generation done")
}
