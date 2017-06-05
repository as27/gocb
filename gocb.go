package gocb

import (
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"
)

var Hasher HashFiler = shaHasher{}

// Equal returns true, when the GOCB files are equal. This function can be changed
// by the caller for a different comparison. The default comparison uses
// -> Name
// -> Size
var Equal = func(f1, f2 GOCBFile) bool {
	if f1.FileName == f2.FileName &&
		f1.Size == f2.Size {
		return true
	}
	return false
}

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
	Path     string `storm:"id"`
	Name     string
	FileName string
	Size     int64
	ModTime  time.Time
}

type GOCBFiles []GOCBFile

func FolderInit(fpath string) ([]GOCBFile, error) {
	var files []GOCBFile
	var counter = 0
	filepath.Walk(fpath, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		absPath, err := filepath.Abs(path)
		if err != nil {
			log.Println(err)
		}
		gfile := GOCBFile{
			Path:     path,
			Name:     absPath,
			FileName: info.Name(),
			Size:     info.Size(),
			ModTime:  info.ModTime(),
		}
		files = append(files, gfile)
		counter++
		return nil
	})
	return files, nil
}

func CheckNotCopiedFiles(src, dst []GOCBFile) []GOCBFile {
	var notCopied []GOCBFile
SrcLoop:
	for _, gf := range src {
		for i, dgf := range dst {
			if Equal(gf, dgf) {
				dst = append(dst[:i], dst[i+1:]...)
				continue SrcLoop
			}
		}
		notCopied = append(notCopied, gf)
	}
	return notCopied

}
