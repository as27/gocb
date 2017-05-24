package main

import (
	"bytes"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"encoding/gob"

	"github.com/as27/gocb"
)

var verbose = false
var src string
var gobfile = "_gocbfile"

func init() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal("Cannot get the working directory\n", err)
	}
	flag.BoolVar(&verbose, "v", false, "Verbose output")
	flag.StringVar(&src, "src", wd, "Set a source folder")
}
func main() {
	flag.Parse()
	log.Println("Initialize\n", src)
	gocb.Verbose = verbose
	files, err := gocb.FolderInit(src)
	log.Println("Finished")
	if err != nil {
		log.Fatal("Got an error while initializing: ", err)
	}
	err = writeFiles(files)
	if err != nil {
		log.Fatal("Got an error writing file", err)
	}
}

func writeFiles(fs gocb.GOCBFiles) error {
	b := &bytes.Buffer{}
	fpath := filepath.Join(src, gobfile)
	enc := gob.NewEncoder(b)
	err := enc.Encode(fs)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(fpath, b.Bytes(), 0777)
	return err
}
