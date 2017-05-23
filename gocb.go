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

	"github.com/asdine/storm"
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

func FolderInit(fpath string) error {
	dbpath := filepath.Join(fpath, DBName)
	db, err := storm.Open(dbpath)
	defer db.Close()
	if err != nil {
		return err
	}
	filepath.Walk(fpath, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if strings.HasPrefix(info.Name(), DBName) {
			return nil
		}
		if Verbose {
			log.Println("Analysing: ", path)
		}
		absPath, err := filepath.Abs(path)
		if err != nil {
			log.Println(err)
		}
		gfile := &GOCBFile{
			Path: path,
		}
		err = db.One("Path", path, gfile)
		if err == storm.ErrNotFound || InitializeAll {
			gfile.Name = absPath
			gfile.Size = info.Size()
			gfile.ModTime = info.ModTime()
		}

		err = db.Save(gfile)
		if err != nil {
			log.Println(err)
		}
		return nil
	})
	return nil
}
