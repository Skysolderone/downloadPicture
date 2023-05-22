package main

import (
	"v1/controller"
	"v1/model"
	"v1/util"
)

func main() {

	//初始化数据库连接
	model.InitDb1()
	model.InitDb2()
	//构建下载链接以及创建对应文件夹
	//urlArray := util.InitUrl()
	//
	////for i := range urlArray {
	////	data, err  := controller.Start(urlArray[i])
	////	if err != nil {
	////		log.Printf("%s, %s",urlArray[i],err)
	////		continue
	////	}
	////}
	//测试数据
	urlArray := util.InitUrl()
	urlArray = append(urlArray, "")
	test := ""
	data, err := controller.Start(test)
	if err != nil {
		println(err)
	}
	util.CompareData(data)
}
