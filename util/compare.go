package util

import (
	"log"
	"strconv"
	"strings"
	"time"

	"gorm.io/hints"
)

func CompareData(hexmodel model.PictureInfo) {
	insertStruct := model.LocalImage{}
	store := ""
	//for _, imageId := range hexmodel {
	compareData := model.GetMAP()
	store = hexmodel.Store
	DbPath := "z_images_" + store
	data := compareData[DbPath]
	query := "/CartvuMedia/FTP/Cartvu/Images/" + store + "/" + hexmodel.FileName
	for i := range data {
		if query == data[i].ImagePath {
			switch data[i].Type {
			case 2:
				insertStruct.Manual_check_type = 0
				break
			case 3:
				insertStruct.Manual_check_type = 2
				break
			case 1:
				insertStruct.Manual_check_type = 1
				break
			default:
				break
			}
		} else {
			insertStruct.Manual_check_type = -1
		}
	}
	insertStruct.Create_time = time.Now()
	insertStruct.Update_time = time.Now()
	insertStruct.Image_name = hexmodel.FileName
	insertStruct.Store_no = hexmodel.Store
	insertStruct.Lane_no = hexmodel.Lane
	t := hexmodel.AlertTime
	time, _ := strconv.ParseInt(t[1:], 10, 64)
	insertStruct.Image_create_time = time
	insertStruct.Confirm = hexmodel.Confirmed
	insertStruct.Edge_forward_type = hexmodel.IsAi
	insertAyyay := model.GetInsertArray()
	if len(insertAyyay) == 1000 {
		model.ClearInsertArray()
		model.InsertChan <- insertAyyay
		model.CheckChan()
	} else {
		model.AddInsertArray(insertStruct)
	}
}

func findPicture(path string, storeId string, imageId string) model.RemoteImage {
	remoteImage := model.RemoteImage{}
	SelectType := model.SelectType{}
	query := "/CartvuMedia/FTP/Cartvu/Images/" + storeId + "/" + imageId
	trimSpace := "/CartvuMedia/FTP/Cartvu/Images/" + storeId + "/"

	model.Db.Table(path).
		Select("Type").
		Where("ImagePath=? and Note=?", query, "audited").
		Clauses(hints.UseIndex("ImagePath")).Find(&SelectType)
	//model.Db.Table(path).Select("Type").
	//	Where("ImagePath=? and Note=?", query, "audited").
	//	Find(&remoteImage)
	if SelectType.Type != "" {
		remoteImage.ImagePath = strings.Trim(remoteImage.ImagePath, trimSpace)
		remoteImage.Type, _ = strconv.Atoi(SelectType.Type)
	}
	return remoteImage

}
func insertData(remoteData model.RemoteImage, pictureInfo model.PictureInfo) {
	insertStruct := model.LocalImage{}

	if remoteData.ImagePath != "" {

	} else {
		insertStruct.Manual_check_type = -1
	}

	insertStruct.Create_time = time.Now()
	insertStruct.Update_time = time.Now()
	insertStruct.Image_name = pictureInfo.FileName
	insertStruct.Store_no = pictureInfo.Store
	insertStruct.Lane_no = pictureInfo.Lane
	t := pictureInfo.AlertTime
	time, _ := strconv.ParseInt(t[1:], 10, 64)
	insertStruct.Image_create_time = time
	insertStruct.Confirm = pictureInfo.Confirmed
	insertStruct.Edge_forward_type = pictureInfo.IsAi
	insertAyyay := model.GetInsertArray()
	if len(insertAyyay) == 500 {
		model.ClearInsertArray()
		model.InsertChan <- insertAyyay
		model.CheckChan()
	} else {
		model.AddInsertArray(insertStruct)
	}
}

func Insert() {
	insertArray := <-model.InsertChan
	err := model.Db2.Table("bob_images").Save(&insertArray).Error
	if err != nil {
		log.Println("Error inserting")
	}
}
