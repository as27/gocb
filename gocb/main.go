package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"

	"encoding/gob"

	"github.com/as27/gocb"
	"github.com/as27/govuegui"
)

var gui = govuegui.NewGui()
var gobfile = "_gocbfile"

var srcFiles gocb.GOCBFiles

func init() {
	govuegui.PathPrefix = ""
	gui.Title = "gocb - Go Check Backup"
}

func main() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	dblForm := gui.Form("Copy check")
	dblBox := dblForm.Box("Check files")

	statField := dblBox.Text("DblStatus")
	statField.SetLabel("Status")
	dblBox.Button("Check folders").Action(
		func(gui *govuegui.Gui) {
			srcPath := gui.Form("Settings").Box("Settings").Input("Src Folder").Get().(string)
			dstPath := gui.Form("Settings").Box("Settings").Input("Dst Folder").Get().(string)
			status := fmt.Sprintf("Analysing folder:<br>%s<br>", srcPath)
			statField.Set(status)
			gui.Update()
			srcFiles, _ := gocb.FolderInit(srcPath)
			status = fmt.Sprintf("%sAnalysing folder:<br>%s<br>", status, dstPath)
			statField.Set(status)
			gui.Update()
			dstFiles, _ := gocb.FolderInit(dstPath)
			status = status + "Comparing folders...<br> "
			statField.Set(status)
			gui.Update()
			doubles := gocb.CheckNotCopiedFiles(srcFiles, dstFiles)
			filesList := []string{}
			for _, f := range doubles {
				filesList = append(filesList, f.Name)
			}
			dblForm.Box("Dbl files").List("Double files").Set(filesList)
			status = status + "Ready"
			statField.Set(status)
			gui.Update()
		})
	statField.SetLabel("Status")
	gui.Form("Settings").Box("Settings").Input("Src Folder").Set(wd)
	gui.Form("Settings").Box("Settings").Input("Dst Folder").Set(wd)
	srcPath := gui.Form("Settings").Box("Settings").Input("Src Folder").Get().(string)
	dstPath := gui.Form("Settings").Box("Settings").Input("Dst Folder").Get().(string)
	dblBox.Text("Folders").Set(fmt.Sprintf("Src: %s<br>Dst: %s", srcPath, dstPath))
	log.Fatal(govuegui.Serve(gui))
	// files, err := gocb.FolderInit(src)
}

func initFolder(fpath string, box, boxFileList *govuegui.Box, boxStatus *govuegui.Element) {
	boxStatus.Set("Initializing Files")
	gui.Update("Status")
	timeStart := time.Now()
	srcFiles, _ = gocb.FolderInit(fpath)
	timeStop := time.Now()
	filesList := []string{}
	for _, f := range srcFiles {
		filesList = append(filesList, f.Name)
	}
	statusText := fmt.Sprintf(`Filelist is ready<br>
			Found %d files.<br>
			In %v`,
		len(srcFiles),
		timeStop.Sub(timeStart))
	boxStatus.Set(statusText)
	boxFileList.List("Files").Set(filesList)
	gui.Update("Files", "Status")
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
