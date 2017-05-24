package gocb

import (
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var DBName = "_gocb.db"

// InitializeAll defines if all entries are generated completely new
var InitializeAll = false

var Verbose = false

var Hasher HashFiler = shaHasher{}

type shaHasher struct {
}

func (sh shaHasher) HashFile(fpath string) string {
	f, err := os.Open(fpath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}

	return fmt.Sprintf("%x", h.Sum(nil))
}

// HashFiler enables to change the hash algorithm
type HashFiler interface {
	HashFile(filepath string) string
}

type GOCBFile struct {
	Path    string `storm:"id"`
	Name    string
	Size    int64
	ModTime time.Time
}

type GOCBFiles []GOCBFile

func FolderInit(fpath string) ([]GOCBFile, error) {
	var files []GOCBFile
	var counter = 0
	filepath.Walk(fpath, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if strings.HasPrefix(info.Name(), DBName) {
			return nil
		}
		if Verbose {
			if counter%100 == 0 {
				log.Println("Analysed: ", counter, " files")
			}
			// log.Println("Analysing: ", path)
		}
		absPath, err := filepath.Abs(path)
		if err != nil {
			log.Println(err)
		}
		gfile := GOCBFile{
			Path:    path,
			Name:    absPath,
			Size:    info.Size(),
			ModTime: info.ModTime(),
		}
		files = append(files, gfile)
		counter++
		return nil
	})
	return files, nil
}
