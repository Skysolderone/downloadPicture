package model

import (
	"time"
)

type StoreId struct {
	Id string
}
type LocalImage struct {
	Id                  uint8     `gorm:"id"`
	Create_time         time.Time `gorm:"create_time"`
	Update_time         time.Time `gorm:"update_time"`
	Image_name          string    `gorm:"image_name"`
	Store_no            string    `gorm:"store_no"`
	Lane_no             uint8     `gorm:"lane_no"`
	Image_create_time   int64     `gorm:"iamge_create_time"` //数据表结构是bigint 疑似要把时间转换
	Confirm             uint8     `gorm:"confirm"`
	Edge_forward_type   uint8     `gorm:"edge_forward_type"`
	Server_forward_type uint8     `gorm:"server_forward_type"`
	Manual_check_type   int8      `gorm:"manual_check_type"`
}

var FinalDownloadSig = make(chan bool, 1)
var InsertArray = make([]LocalImage, 0)
var UrlArray = make([]string, 0)
var InsertChan = make(chan []LocalImage, 100)
var InsertDataChanSig = make(chan string, 10)

func GetInsertArray() []LocalImage {
	return InsertArray
}
func ClearInsertArray() {
	InsertArray = make([]LocalImage, 0)
}
func AddInsertArray(data LocalImage) {
	InsertArray = append(InsertArray, data)
}
func CreateUrl() {
	UrlArray = InitUrl()
}
func GetUrlArray() []string {
	return UrlArray
}
func ForwardArray() {
	if len(UrlArray) != 1 {
		UrlArray = UrlArray[1:]
		FinalDownloadSig <- false
	} else {
		FinalDownloadSig <- true
	}

}
func CheckChan() {
	if len(InsertChan) != 0 {
		InsertDataChanSig <- "StartInsert"
	}
}
