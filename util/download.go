package util

import (
	"DownLoadPicture/model"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func Download(filepath string) (string, error) {
	log.Println("starting download")
	//filepath = "http://images.checkpointcn.com/202206/002235-202206.tgz"
	resp, err := http.Get(filepath)
	defer resp.Body.Close()

	if err != nil {
		model.AddData(model.FatalPath, filepath)
		return "", err
	}

	if resp.StatusCode != 200 {
		remove(filepath)
		return "", errors.New(fmt.Sprintf("file not exist,%s", filepath))
	}
	//stringTrim(string(resp.Body))
	//localFilePath := "002235-202206.tgz"
	trimstring := strings.Trim(filepath, "http://images.checkpointcn.com/")
	s := strings.Index(trimstring, "-")
	localFilePath := "./downloadFile/" + trimstring[:s] + trimstring[6:]

	file, err := os.Create(localFilePath)
	if err != nil {
		model.AddData(model.FatalPath, filepath)
		log.Println("create file error", err)
		return "", err
	}
	defer file.Close()

	writesize, err := io.Copy(file, resp.Body)
	if err != nil {
		model.AddData(model.FatalPath, filepath)
		log.Println("copy errr", err)
		return "", err
	}
	model.AddData(model.SuccessPath, trimstring)
	log.Printf("url download success:%s,download size:%dMb", trimstring[6:], writesize/1024/1000)

	return trimstring, nil

}
func remove(path string) {
	path = strings.Trim(path, "http://images.checkpointcn.com")
	localPath := "./downloadFile/" + path[0:13] + "/"
	_, err := os.Stat(localPath)
	if err != nil {
		log.Println("file not exist", err)
	}
	err = os.RemoveAll(localPath)
	if err != nil {
		log.Println("remove err", err)
		model.AddData(model.FatalRemove, path)

	}
}
