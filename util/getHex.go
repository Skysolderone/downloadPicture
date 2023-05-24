package util

import (
	"DownLoadPicture/model"
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"reflect"
	"strings"
	"unsafe"
)

func GetHexData(filepath string, storeId string) (model.PictureInfo, error) {
	file, err := os.Open(filepath)
	if err != nil {
		log.Println(err)
	}
	defer removeFile(filepath)
	defer file.Close()
	fileInfo, err := file.Stat()
	if err != nil {
		log.Println(filepath)
		return model.PictureInfo{}, errors.New(fmt.Sprintf("file not exist :", err))
	}
	size := fileInfo.Size()
	_, err = file.Seek(size-153, 1)
	if err != nil {
		return model.PictureInfo{}, errors.New(fmt.Sprintf("seek error", err))
	}

	r := bufio.NewReader(file)
	chunks := make([]byte, 0)
	buf := make([]byte, 153)

	_, err = r.Read(buf)
	if err != nil && err != io.EOF {
		panic(err)
	}
	if buf[0] == byte(255) && buf[1] == byte(217) {
		chunks = append(chunks, buf[2:]...)
	} else if buf[1] == byte(255) && buf[2] == byte(217) {
		chunks = append(chunks, buf[3:]...)
	}
	pictureInfo := model.PictureInfo{}
	var cmd_info *model.GetPictureInfo

	cmd_info = (*model.GetPictureInfo)(byte_slice_to_struct(chunks))

	pictureInfo.FileName = fmt.Sprintf("%s", cmd_info.FileName)
	pictureInfo.AlertTime = fmt.Sprintf("%s", cmd_info.AlertTime)
	pictureInfo.Lane = cmd_info.Lane
	pictureInfo.Store = storeId
	pictureInfo.IsAi = cmd_info.IsAi

	pictureInfo.Confirmed = cmd_info.Confirmed

	model.SuccessGetHex = append(model.SuccessGetHex, filepath)
	return pictureInfo, nil

}
func byte_slice_to_struct(data []byte) unsafe.Pointer {
	return unsafe.Pointer((*reflect.SliceHeader)(unsafe.Pointer(&data)).Data)
}
func removeFile(filepath string) {
	path := strings.Trim(filepath, ".")

	err := os.Remove("." + path)
	if err != nil {
		log.Println("remove image file fail :", err)
		model.FatalRemove = append(model.FatalRemove, path)
	}
}
