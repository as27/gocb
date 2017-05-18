package gocb

import (
	"os"
)

// HashFiler enables to change the hash algorithm
type HashFiler interface {
	HashFile(filepath string) string
}

type GOCBFile struct {
	ID    string `storm:"id"`
	Files []FileDef
}

type FileDef struct {
	FileInfo os.FileInfo
	Fpath    string
}
