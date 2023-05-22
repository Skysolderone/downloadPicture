package util

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"reflect"
	"strings"
	"unsafe"
	"v1/model"
)

func GetHex(filepath string) (model.PictureInfo, error) {
	//hexImage := model.HexImage{}
	//file, err := os.Open("../image/I-01-20220601034407-0.jpg")
	file, err := os.Open(filepath)
	if err != nil {
		log.Println(err)
	}
	fileInfo, err := file.Stat()
	size := fileInfo.Size()
	_, err = file.Seek(size-153, 1)
	if err != nil {
		return model.PictureInfo{}, errors.New(fmt.Sprintf("seek error", err))
	}

	defer file.Close()
	r := bufio.NewReader(file)
	chunks := make([]byte, 0)
	buf := make([]byte, 153)

	_, err = r.Read(buf)
	if err != nil && err != io.EOF {
		panic(err)
	}
	if buf[0] == byte(255) && buf[1] == byte(217) {
		chunks = append(chunks, buf[2:]...)
	}
	pictureInfo := model.PictureInfo{}
	var cmd_info *model.GetPictureInfo

	cmd_info = (*model.GetPictureInfo)(byte_slice_to_struct(chunks))

	pictureInfo.FileName = fmt.Sprintf("%s", cmd_info.FileName)
	pictureInfo.AlertTime = fmt.Sprintf("%s", cmd_info.AlertTime)
	pictureInfo.Lane = cmd_info.Lane

	pictureInfo.IsAi = cmd_info.IsAi

	pictureInfo.Store = fmt.Sprintf("%s", cmd_info.Store)

	pictureInfo.Confirmed = cmd_info.Confirmed
	//hex_data := make([]byte, len(chunks), cap(chunks))
	//result := hex.EncodeToString(chunks2)
	//hex_data, _ = hex.DecodeString(result)
	////把数据写进结构体 alarmtime imagename
	//hexImage.ImageName = string(hex_data[:25])
	path := strings.Trim(filepath, ".")
	err = os.Remove(path)
	if err != nil {
		return model.PictureInfo{}, errors.New(fmt.Sprintf("remove image file %s,error", filepath))

	}

	return pictureInfo, nil

}
func byte_slice_to_struct(data []byte) unsafe.Pointer {
	return unsafe.Pointer((*reflect.SliceHeader)(unsafe.Pointer(&data)).Data)
}
