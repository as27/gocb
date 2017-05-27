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
	src := GOCBFiles{file1, file2, file3}
	dst := GOCBFiles{file1copied, file2copied}
	got := CheckNotCopiedFiles(src, dst)
	expect := GOCBFiles{file3}
	if !reflect.DeepEqual(got, expect) {
		t.Error("Expect: ", expect)
	}
}
