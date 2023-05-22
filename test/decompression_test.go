package test

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"os"
	"strings"
	"testing"
)

func TestDecompression(t *testing.T) {
	filePath := "002235-202206.tgz"
	srcFile, err := os.Open(filePath)
	if err != nil {
		t.Log(err)
	}
	defer srcFile.Close()
	gr, err := gzip.NewReader(srcFile)
	if err != nil {
		t.Log(err)
	}
	defer gr.Close()
	tr := tar.NewReader(gr)
	for {
		hdr, err := tr.Next()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				t.Log(err)
			}
		}
		filename := "../image" + hdr.Name
		err = os.MkdirAll(string([]rune(filename)[0:strings.LastIndex(filename, "/")]), 0755)
		if err != nil {
			t.Log(err)
		}
		file, err := os.Create(filename)
		if err != nil {
			t.Log(err)
		}
		io.Copy(file, tr)
	}

}
