package controller

import (
	"DownLoadPicture/model"
	"DownLoadPicture/util"
	"log"
	"os"
)

// 下载
var DownloadChan = make(chan string, 100)

// 对比数据
var CompareChan = make(chan string, 100)

// 解压
var DecompressChan = make(chan string, 100)

// 获取图片元数据
var GetHexChan = make(chan string, 100)

// 插入数据
var DataChan = make(chan model.PictureInfo, 5000)
var InsertChan = make(chan []model.LocalImage, 100)

// 公共信息
var PublicChan = make(chan string, 100)

// 退出
var QuitChan = make(chan string, 1)

// 协程执行信号
var DecompressChanSig = make(chan string, 100)
var GetHexChanSig = make(chan string, 1000)
var ExitChan = make(chan bool, 1)

// 图片文件
var DirFileData = make(chan string, 5000)

func Download() {
	urlArray := model.GetUrlArray()
	downloadPath, err := util.Download(urlArray[0])
	if err != nil {
		log.Printf("%s, %s", urlArray[0], err)
		model.ForwardArray()
		if s := <-model.FinalDownloadSig; s {
			DownloadChan <- "Final Download Url"
			QuitChan <- "Have Finished Goroutine"
		} else {
			//model.AddStringToChannel(model.DownloadChan, "download next url")
			DownloadChan <- "download next url"
		}
	} else {
		model.ForwardArray()
		if s := <-model.FinalDownloadSig; s {
			DownloadChan <- "Final Download Url"
			QuitChan <- "Have Finished Goroutine"
		} else {
			//model.AddStringToChannel(model.DownloadChan, "download next url")
			DecompressChanSig <- "start decompress"
			DecompressChan <- downloadPath
			DownloadChan <- "download next url"
		}

	}
}
func Decompress1() {
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
func Decompress() {
	downloadPath := <-DecompressChan
	//downloadPath := <-model.DecompressChan
	log.Println("Decompress", downloadPath)
	imagePath, err := util.Decompression(downloadPath)
	if err != nil {
		log.Println(err)
	}

	files, err := os.ReadDir(imagePath)
	if err != nil {
		log.Println("read dir error:", err)
		model.FatalGetHex = append(model.FatalGetHex, imagePath)
	}
	totalFiles := len(files)
	for i, file := range files {
		imageFilePath := imagePath + "/" + file.Name()
		DirFileData <- imageFilePath
		if len(DirFileData) > 3000 {
			GetHexChanSig <- "start get hex"
		}
		//如果是最后一个文件，则退出循环
		if i == totalFiles-1 {
			break
		}
	}
	GetHexChanSig <- "start get hex"
	//model.AddStringToChannel(model.GetHexChan, imagePath)
	//model.AddStringToChannel(model.GetHexChanSig, "start get hex")
}

func GetHex1() {
	imagePath := <-GetHexChan
	log.Println("GetHex", imagePath)
	files, err := os.ReadDir(imagePath)
	if err != nil {
		log.Println("read dir error:", err)
		model.FatalGetHex = append(model.FatalGetHex, imagePath)
	}

	totalFiles := len(files)
	storeId := imagePath[15:21]

	for i, file := range files {
		imageFilePath := imagePath + "/" + file.Name()
		pictureInfo, err := util.GetHexData(imageFilePath, storeId)
		if err != nil {
			log.Println(err)
			model.FatalGetHex = append(model.FatalGetHex, imageFilePath)
		}
		//s := model.GetDataChan()
		//cout := len(s)
		if len(DataChan) < 1000 {
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
func GetHex() {
	file := <-DirFileData
	storeId := file[15:21]
	pictureInfo, err := util.GetHexData(file, storeId)
	if err != nil {
		log.Println(err)
		model.FatalGetHex = append(model.FatalGetHex, file)
	}
	if len(DataChan) < 500 {
		//model.AddDataToChannel(model.DataChan, pictureInfo)
		DataChan <- pictureInfo
	} else {
		//model.AddStringToChannel(model.CompareChan, "start compare")
		CompareChan <- "start compare"
	}
	model.SuccessGetHex = append(model.SuccessGetHex, file)
}

func Compare() {
	data := <-DataChan
	util.CompareData(data)

}
func Quit() {
	switch {
	case len(DownloadChan) != 0:
		QuitChan <- "waiting for download finish"
		DownloadChan <- "Final download"
	case len(DecompressChan) != 0:
		QuitChan <- "waiting for Decompress finish"
		DecompressChanSig <- "Final Decompress"
	case len(GetHexChan) != 0:
		QuitChan <- "waiting for GetHex finish"
		DecompressChanSig <- "Final GetHex"
	case len(CompareChan) != 0:
		QuitChan <- "waiting for Compare finish"
		DecompressChanSig <- "Final Compare"
	case len(model.InsertChan) != 0:
		QuitChan <- "waiting for insert finish"
		DecompressChanSig <- "Final insertdata"
	default:
		close(DownloadChan)
		close(CompareChan)
		close(DecompressChan)
		close(GetHexChan)
		close(DataChan)
		close(PublicChan)
		close(DecompressChanSig)
		close(GetHexChanSig)
		close(model.InsertChan)
		close(QuitChan)
		ExitChan <- true
	}
}
