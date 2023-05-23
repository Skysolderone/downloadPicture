package controller

import (
	"DownLoadPicture/model"
	"DownLoadPicture/util"
	"fmt"
	"log"
	"os"
)

var DownloadChan = make(chan string, 10)
var CompareChan = make(chan string, 10)
var DecompressChan = make(chan string, 10)
var GetHexChan = make(chan string, 10)
var DataChan = make(chan model.PictureInfo, 1000)
var PublicChan = make(chan string, 100)

var DecompressChanSig = make(chan string, 10)
var GetHexChanSig = make(chan string, 10)

//var picinfo [][]model.PictureInfo

//	func Start(filepath string) (model.PictureInfo, error) {
//		//data := []model.HexImage{}
//		picdata := model.PictureInfo{}
//		downloadPath, err := util.Download(filepath)
//		if err != nil {
//			log.Println(err)
//			return picdata, errors.New("download err")
//		}
//		//测试数据
//		//downloadPath := "202206/002235-202206.tgz"
//		imagePath, err := util.Decompression(downloadPath)
//		if err != nil {
//			log.Println(err)
//			return picdata, errors.New("Decompression err")
//		}
//
//		//imagePath := "../image/202206/002235/002235-202206.tgz"
//		// 要读取文件的目录
//		files, err := os.ReadDir(imagePath)
//		if err != nil {
//			log.Println("read dir error:", err)
//			return picdata, err
//		}
//
//		totalFiles := len(files)
//		storeId := imagePath[16:22]
//		log.Print("Get Hex")
//		for i, file := range files {
//			imageFilePath := imagePath + "/" + file.Name()
//			pictureInfo, err := util.GetHex(imageFilePath, storeId)
//			if err != nil {
//				log.Println(err)
//				return picdata, errors.New("GetHex err")
//			}
//			//data = append(data, hexStruct)
//			if len(model.DataChan) < 100 {
//				//model.AddDataToChannel(model.DataChan, pictureInfo)
//				DataChan <- pictureInfo
//			} else {
//				//model.AddStringToChannel(model.CompareChan, "start compare")
//				CompareChan <- "start compare"
//			}
//			// 如果是最后一个文件，则退出循环
//			if i == totalFiles-1 {
//				break
//			}
//		}
//		log.Println("Get Hex done")
//		model.SuccessGetHex = append(model.SuccessGetHex, imagePath)
//
//		return picdata, nil
//	}
func Download() {
	urlArray := model.GetUrlArray()
	downloadPath, err := util.Download(urlArray[0])
	if err != nil {
		log.Printf("%s, %s", urlArray[0], err)
		model.ForwardArray()
		//model.AddStringToChannel(model.DownloadChan, "download next url")
		DownloadChan <- "download next url"
	} else {
		model.PublicChan <- fmt.Sprintf("download successfully,%s", urlArray[0])
		model.ForwardArray()
		//model.AddStringToChannel(model.DecompressChan, downloadPath)
		//model.AddStringToChannel(model.DecompressChanSig, "start decompress")
		DecompressChanSig <- "start decompress"
		DecompressChan <- downloadPath
		//model.AddStringToChannel(model.DownloadChan, "download next url")
		DownloadChan <- "download next url"
	}
}
func Decompress() {
	downloadPath := <-DecompressChan
	//downloadPath := <-model.DecompressChan
	log.Println("Decompress", downloadPath)
	imagePath, err := util.Decompression(downloadPath)
	if err != nil {
		log.Println(err)
	}
	//model.AddStringToChannel(model.GetHexChan, imagePath)
	//model.AddStringToChannel(model.GetHexChanSig, "start get hex")
	GetHexChanSig <- "start get hex"
	GetHexChan <- imagePath
}
func GetHex() {
	imagePath := <-GetHexChan
	log.Println("GetHex", imagePath)
	files, err := os.ReadDir(imagePath)
	if err != nil {
		log.Println("read dir error:", err)
		model.FatalGetHex = append(model.FatalGetHex, imagePath)
	}

	totalFiles := len(files)
	storeId := imagePath[16:22]

	for i, file := range files {
		imageFilePath := imagePath + "/" + file.Name()
		pictureInfo, err := util.GetHex(imageFilePath, storeId)
		if err != nil {
			log.Println(err)
			model.FatalGetHex = append(model.FatalGetHex, imageFilePath)
		}
		//s := model.GetDataChan()
		//cout := len(s)
		if len(DataChan) < 500 {
			//model.AddDataToChannel(model.DataChan, pictureInfo)
			DataChan <- pictureInfo
		} else {
			//model.AddStringToChannel(model.CompareChan, "start compare")
			CompareChan <- "start compare"
		}
		// 如果是最后一个文件，则退出循环
		if i == totalFiles-1 {
			break
		}
	}
	log.Println("Get Hex done")
	model.SuccessGetHex = append(model.SuccessGetHex, imagePath)
}

//	func DownloadAndDecompress() {
//		urlArray := model.GetUrlArray()
//		//data, err := Start(UrlArray[0])
//		_, err := Start(urlArray[0])
//		if err != nil {
//			log.Printf("%s, %s", urlArray[0], err)
//			model.ForwardArray()
//			model.DownloadChan <- "download next url"
//		} else {
//			//if len(DataChan) != 100 {
//			//	DataChan <- data
//			//}
//			//picinfo = append(picinfo, data)
//			model.PublicChan <- fmt.Sprintf(" download and decompress successfully,%s", urlArray[0])
//			model.ForwardArray()
//			model.DownloadChan <- "download next url"
//		}
//
// }
func Compare() {
	//if len(picinfo) != 0 {
	data := <-DataChan
	util.CompareData(data)
	//}
}
func CloseChannel() {
	close(DownloadChan)
	close(CompareChan)
	close(DecompressChan)
	close(GetHexChan)
	close(DataChan)
	close(PublicChan)
	close(DecompressChanSig)
	close(GetHexChanSig)
}
