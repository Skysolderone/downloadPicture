package util

import (
	"log"
	"os"
	"v1/model"
)

func InitUrl() []string {
	result := make([]string, 0)
	storeId := []model.StoreId{}
	model.Db.Table("storeinfo").Select("Id").Find(&storeId)
	ip := "http://images.checkpointcn.com"
	remoUrl := make([]string, 0)
	//remoUrl[1] = "/202201-202205/202201-202205.tgz"
	remoUrl = append(remoUrl, "202206")
	//remoUrl = append(remoUrl, "202207")
	//remoUrl = append(remoUrl, "202208")
	//remoUrl = append(remoUrl, "202209")
	//remoUrl = append(remoUrl, "202210")
	//remoUrl = append(remoUrl, "202211")
	downloadPath := ""
	//start
	for i := range remoUrl {
		for j := range storeId {
			filePath := "./downloadFile" + "/" + remoUrl[i] + "/" + storeId[j].Id
			Create(filePath)
			urlPath := ip + "/" + remoUrl[i] + "/" + storeId[j].Id
			downloadPath = urlPath + "-" + remoUrl[i] + ".tgz"
			result = append(result, downloadPath)
		}
	}
	return result
}
func Create(localfilepath string) {
	err := os.MkdirAll(localfilepath, 0755)
	if err != nil && !os.IsExist(err) {
		log.Println("mkdir error:", err)
	}
}
