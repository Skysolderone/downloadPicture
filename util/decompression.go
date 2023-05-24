package util

import (
	"DownLoadPicture/model"
	"archive/tar"
	"compress/gzip"
	"io"
	"log"
	"os"
	"strings"
)

func Decompression(tarName string) (string, error) {
	s := strings.Index(tarName, "-")
	localFilePath := "./image/" + tarName[:s] + tarName[6:]
	downloadPath := "./downloadFile/" + tarName[:s] + tarName[6:]
	srcFile, err := os.Open(downloadPath)
	if err != nil {
		return "", err
	}
	removePath := strings.Trim(downloadPath, tarName[s-6:])
	defer removeDownloadFile(removePath)
	defer srcFile.Close()

	gr, err := gzip.NewReader(srcFile)
	if err != nil {
		return "", err
	}
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
			model.AddData(model.FatalDecompress, filename)
			return "", err
		}
		file, err := os.Create(filename)
		if err != nil {
			model.AddData(model.FatalDecompress, filename)
			return "", err
		}
		_, err = io.Copy(file, tr)
		if err != nil {
			model.AddData(model.FatalDecompress, filename)
			return "", err
		}
	}
	log.Println("decompressing image:", cout)
	model.AddData(model.SuccessDecompress, downloadPath)

	return localFilePath, nil
}
func removeDownloadFile(file string) {
	localPath := "." + file[:len(file)-1] + "/"
	err := os.RemoveAll(localPath)
	if err != nil {
		log.Println("remove file failed err:", err)
		model.AddData(model.FatalRemove, file)
	} else {
		file = strings.Trim(file, "")
		model.AddData(model.SuccessDecompress, file)
		log.Println("remove downloaded file:", file)
	}
}
