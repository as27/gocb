package gocb

import "testing"

func TestFolderInit(t *testing.T) {
	type args struct {
		fpath string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := FolderInit(tt.args.fpath); (err != nil) != tt.wantErr {
				t.Errorf("FolderInit() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFpath(t *testing.T) {
	Verbose = true
	FolderInit("../")

}
