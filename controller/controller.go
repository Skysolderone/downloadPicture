package controller

import (
	"errors"
	"log"
	"os"
	"v1/model"
	"v1/util"
)

func Start(filepath string) ([]model.PictureInfo, error) {
	//data := []model.HexImage{}
	picdata := []model.PictureInfo{}
	//downloadPath, err := util.Download(filepath)
	//if err != nil {
	//	log.Println(err)
	//	return errors.New("download err")
	//}
	//测试数据
	downloadPath := "202206/002235-202206.tgz"
	imagePath, err := util.Decompression(downloadPath)
	if err != nil {
		log.Println(err)
		return picdata, errors.New("Decompression err")
	}
	//imagePath := "../image/202206/002235/002235-202206.tgz"
	// 要读取文件的目录
	files, err := os.ReadDir(imagePath)
	if err != nil {
		log.Println("read dir error:", err)
		return picdata, err
	}

	totalFiles := len(files)
	for i, file := range files {
		imageFilePath := imagePath + "/" + file.Name()
		pictureInfo, err := util.GetHex(imageFilePath)
		if err != nil {
			log.Println(err)
			return picdata, errors.New("GetHex err")
		}

		//data = append(data, hexStruct)
		picdata = append(picdata, pictureInfo)
		// 如果是最后一个文件，则退出循环
		if i == totalFiles-1 {
			break
		}
	}
	model.SuccessGetHex = append(model.SuccessGetHex, imagePath)
	return picdata, nil
}
