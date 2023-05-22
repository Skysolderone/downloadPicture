package util

import (
	"strings"
	"v1/model"
)

func CompareData(hexmodel []model.PictureInfo) {

	storeId := []model.StoreId{}
	model.Db.Table("storeinfo").Select("Id").Find(&storeId)
	for _, store := range storeId {
		DbPath := "z_images_" + store.Id
		for _, imageId := range hexmodel {
			remoteData := findPicture(DbPath, store.Id, imageId.FileName)
			insertData(remoteData, imageId)
		}

	}

}
func findPicture(path string, storeid string, imageId string) []model.RemoteImage {
	remoteImageArray := []model.RemoteImage{}
	query := "/CartvuMedia/FTP/Cartvu/Images/" + storeid + "/" + imageId
	trimSpace := "/CartvuMedia/FTP/Cartvu/Images/" + storeid + "/"
	model.Db.Table(path).Select("ImagePath,Note,Type").Where("ImagePath=?", query).Find(&remoteImageArray)
	for _, image := range remoteImageArray {
		image.ImagePath = strings.Trim(image.ImagePath, trimSpace)
	}
	return remoteImageArray
}
func insertData(remoteData []model.RemoteImage, pictureInfo model.PictureInfo) {
	for _, image := range remoteData {
		insertStruct := model.LocalImage{}
		if image.Node != "" {
			insertStruct.Server_forward_type = 1
		}
		insertStruct.Image_name = pictureInfo.FileName
		insertStruct.Store_no = pictureInfo.Store
		insertStruct.Lane_no = pictureInfo.Lane
		insertStruct.Image_create_time = pictureInfo.AlertTime
		insertStruct.Confirm = pictureInfo.Confirmed
		insertStruct.Edge_forward_type = pictureInfo.IsAi

		model.Db2.Table("bob_images").Save(&insertStruct)
	}
}
