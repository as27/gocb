package main

import (
	"flag"
	"log"
	"os"

	"github.com/as27/gocb"
)

var verbose = false
var src string

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
	err := gocb.FolderInit(src)
	if err != nil {
		log.Fatal("Got an error while initializing: ", err)
	}
}
