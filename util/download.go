package util

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"v1/model"
)

func Download(filepath string) (string, error) {
	//只下载该数据进行测试
	filepath = "http://images.checkpointcn.com/202206/002235-202206.tgz"
	resp, err := http.Get(filepath)
	if err != nil {
		log.Println("Download url error:", err)
		model.FatalPath = append(model.FatalPath, filepath)
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		err := os.Remove(filepath)
		if err != nil {
			model.FatalRemove = append(model.FatalRemove, filepath)
			return "", err
		}
		return "", errors.New(fmt.Sprintf("file not exist,%s", filepath))
	}
	//stringTrim(string(resp.Body))
	//localFilePath := "002235-202206.tgz"
	trimstring := strings.Trim(filepath, "http://images.checkpointcn.com/")
	s := strings.Index(trimstring, "-")
	localFilePath := "./downloadFile/" + trimstring[:s] + trimstring[6:]

	file, err := os.Create(localFilePath)
	if err != nil {
		model.FatalPath = append(model.FatalPath, filepath)
		log.Println("create file error", err)
		return "", err
	}
	defer file.Close()
	writesize, err := io.Copy(file, resp.Body)
	if err != nil {
		model.FatalPath = append(model.FatalPath, filepath)
		log.Println("copy errr", err)
		return "", err
	}
	model.SuccessPath = append(model.SuccessPath, trimstring)
	log.Printf("url download success:%s,download size:%d", trimstring[6:], writesize)
	return trimstring, nil

}
