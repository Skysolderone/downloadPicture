package main

import (
	"DownLoadPicture/controller"
	"DownLoadPicture/model"
	"DownLoadPicture/util"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/panjf2000/ants/v2"
)

func main() {
	// 创建、追加、读写，777，所有权限
	f, err := os.OpenFile("log.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
	if err != nil {
		return
	}
	defer func() {
		f.Close()
	}()
	// os.Stdout代表标准输出流
	multiWriter := io.MultiWriter(os.Stdout, f)
	log.SetOutput(multiWriter)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	pool, _ := ants.NewPool(100, ants.WithPreAlloc(true))
	defer pool.Release()
	//初始化数据库连接
	//构建下载链接以及创建对应文件夹
	init := make(chan string, 1)
	init <- "start init InitDb"
	defer close(init)
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
			pool.Submit(controller.Download)
			//gopool.Go(controller.Download)
			//go controller.Download()
		case dccompress := <-controller.DecompressChanSig:
			log.Println(dccompress)
			//for i := 0; i < 5; i++ {
			pool.Submit(controller.Decompress)
			//	//gopool.Go(controller.Decompress)
			//	go controller.Decompress()
			//}
		case <-controller.GetHexChanSig:
			//log.Println(hex)
			//for i := 0; i < 100; i++ {
			pool.Submit(controller.GetHex)
			//gopool.Go(controller.GetHex)
			//go controller.GetHex()
			//}
		case <-controller.CompareChan:
			//for i := 0; i < 100; i++ {
			pool.Submit(controller.Compare)
			//gopool.Go(controller.Compare)
			//go controller.Compare()
			//}
		case insertData := <-model.InsertDataChanSig:
			fmt.Println(insertData)
			//for i := 0; i < 10; i++ {
			pool.Submit(util.Insert)
			//go util.Insert()
			//}
		case quit := <-controller.QuitChan:
			log.Println(quit)
			pool.Submit(controller.Quit)
		case <-controller.ExitChan:
			log.Println(model.FatalPath)
			log.Println(model.FatalDecompress)
			log.Println(model.FatalGetHex)
			log.Println(model.FatalRemove)
			return
		}
	}

}
