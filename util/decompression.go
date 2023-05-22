package util

import (
	"archive/tar"
	"compress/gzip"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"v1/model"
)

func Decompression(tarName string) (string, error) {
	//filePath := "002235-202206.tgz"
	s := strings.Index(tarName, "-")
	localFilePath := "../image/" + tarName[:s] + tarName[6:]
	downloadPath := "./downloadFile/" + tarName[:s] + tarName[6:]
	srcFile, err := os.Open(downloadPath)
	if err != nil {
		return "", err
	}
	defer srcFile.Close()
	gr, err := gzip.NewReader(srcFile)
	if err != nil {
		return "", err
	}
	defer gr.Close()
	tr := tar.NewReader(gr)
	cout := 0
	for {
		cout++
		hdr, err := tr.Next()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return "", err
			}
		}
		filename := localFilePath + hdr.Name
		err = os.MkdirAll(string([]rune(filename)[0:strings.LastIndex(filename, "/")]), 0755)
		if err != nil {
			model.FatalDecompress = append(model.FatalDecompress, filename)
			return "", err
		}
		file, err := os.Create(filename)
		if err != nil {
			model.FatalDecompress = append(model.FatalDecompress, filename)
			return "", err
		}
		_, err = io.Copy(file, tr)
		if err != nil {
			model.FatalDecompress = append(model.FatalDecompress, filename)
			return "", err
		}
	}
	log.Println("decompressing image:", cout)
	//path := downloadPath[1:]
	err = os.RemoveAll(downloadPath)
	if err != nil {
		return "", errors.New(fmt.Sprintf("remove file failed %s", downloadPath))
	}
	log.Println("remove downloaded file:", tarName)
	model.SuccessDecompress = append(model.SuccessDecompress, downloadPath)
	return localFilePath, nil
}
