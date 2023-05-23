package main

import (
	"DownLoadPicture/controller"
	"DownLoadPicture/model"
	"fmt"
	"log"
)

func main() {

	//初始化数据库连接
	//构建下载链接以及创建对应文件夹
	init := make(chan string, 1)
	init <- "start init InitDb"
	defer close(init)
	defer controller.CloseChannel()
	for {
		select {
		case initDb := <-init:
			fmt.Println(initDb)
			model.InitDb1()
			model.InitDb2()
			model.CreateUrl()
			//model.AddStringToChannel(model.DownloadChan, "exe start ")
			controller.DownloadChan <- "exe start "
		case download := <-controller.DownloadChan:
			log.Println(download)
			go controller.Download()
		case dccompress := <-controller.DecompressChanSig:
			log.Println(dccompress)
			for i := 0; i < 10; i++ {
				go controller.Decompress()
			}
		case hex := <-controller.GetHexChanSig:
			log.Println(hex)
			for i := 0; i < 10; i++ {
				go controller.GetHex()
			}
		case <-controller.CompareChan:
			for i := 0; i < 10; i++ {
				go controller.Compare()
			}
		case msg := <-controller.PublicChan:
			fmt.Println(msg)
			fmt.Println(model.FatalPath)
			fmt.Println(model.FatalDecompress)
			fmt.Println(model.FatalGetHex)
			fmt.Println(model.FatalRemove)
		}
	}

	////测试数据
	//urlArray := util.InitUrl()
	//urlArray = append(urlArray, "")
	//test := ""
	//data, err := controller.Start(test)
	//if err != nil {
	//	println(err)
	//}

}
