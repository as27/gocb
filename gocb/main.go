package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"encoding/gob"

	"github.com/as27/gocb"
	"github.com/as27/govuegui"
)

var gui = govuegui.NewGui()
var gobfile = "_gocbfile"

var srcFiles gocb.GOCBFiles

func main() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	srcForm := gui.Form("Src Folder")
	srcBox := srcForm.Box("Src Folder")
	srcBoxFilelist := srcForm.Box("Filelist")
	srcBox.Button("Initalize Src").Action(
		func() {
			srcBox.Text("Status").Set("Initializing Files")
			gui.Update("Status")
			src := gui.Form("Settings").Box("Settings").Input("Src Folder").Get().(string)
			srcFiles, _ = gocb.FolderInit(src)
			filesTxt := ""
			for _, f := range srcFiles {
				filesTxt = filesTxt + f.Name + "<br>\n"
			}
			srcBox.Text("Status").Set("Filelist is ready")
			srcBoxFilelist.Text("Files").Set(filesTxt)
			gui.Update("Files", "Status")
		})
	srcBox.Text("Status").Set("")
	gui.Form("Dts Folder").Box("Dst Folder").Button("Initialize Dst")
	gui.Form("Settings").Box("Settings").Input("Src Folder").Set(wd)
	gui.Form("Settings").Box("Settings").Input("Dst Folder").Set(wd)
	log.Fatal(govuegui.Serve(gui))
	// files, err := gocb.FolderInit(src)
}

func writeFiles(src string, fs gocb.GOCBFiles) error {
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
func readFiles(src string) (gocb.GOCBFiles, error) {
	var gocbfiles gocb.GOCBFiles
	fpath := filepath.Join(src, gobfile)
	by, err := ioutil.ReadFile(fpath)
	if err != nil {
		return gocbfiles, err
	}
	b := bytes.NewBuffer(by)
	dec := gob.NewDecoder(b)
	err = dec.Decode(&gocbfiles)
	return gocbfiles, err
}
