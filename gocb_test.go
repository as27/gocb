package gocb

import (
	"reflect"
	"testing"
	"time"
)

func TestCheckNotCopiedFiles(t *testing.T) {
	file1 := GOCBFile{
		Path:    "a/b/c",
		Name:    "c",
		Size:    123,
		ModTime: time.Date(2006, time.January, 01, 01, 01, 10, 0, time.UTC),
	}
	file2 := GOCBFile{
		Path:    "a/b/d",
		Name:    "d",
		Size:    234,
		ModTime: time.Date(2006, time.January, 01, 01, 01, 10, 0, time.UTC),
	}
	file3 := GOCBFile{
		Path:    "a/b/e",
		Name:    "e",
		Size:    345,
		ModTime: time.Date(2006, time.January, 01, 01, 01, 10, 0, time.UTC),
	}
	file4 := GOCBFile{
		Path:    "a/b/f",
		Name:    "f",
		Size:    3445,
		ModTime: time.Date(2006, time.January, 01, 01, 01, 10, 0, time.UTC),
	}
	file1copied := GOCBFile{
		Path:    "x/y/c",
		Name:    "c",
		Size:    123,
		ModTime: time.Date(2006, time.January, 01, 01, 01, 10, 0, time.UTC),
	}
	file2copied := GOCBFile{
		Path:    "u/d",
		Name:    "d",
		Size:    234,
		ModTime: time.Date(2006, time.January, 01, 01, 01, 10, 0, time.UTC),
	}
	src := []GOCBFile{file2, file3, file4, file1}
	dst := []GOCBFile{file1copied, file2copied}
	got := CheckNotCopiedFiles(src, dst)
	expect := []GOCBFile{file3, file4}
	if !reflect.DeepEqual(got, expect) {
		t.Errorf("\nExp: %v\nGot: %v ", expect, got)
	}
}
