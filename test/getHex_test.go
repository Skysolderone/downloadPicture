package test

import (
	"DownLoadPicture/model"
	"bufio"
	"fmt"
	"io"
	"os"
	"reflect"
	"strconv"
	"testing"
	"unsafe"
)

type PictureInfo1 struct {
	FileName  string /*图片文件名 ASCII码 格式:I-收银台号-yyyyMMddhhmmss.jpg*/
	AlertTime string /*图片的报警时间 ASCII码 格式:yyyyMMddhhmmss*/
	Lane      uint8  /*收银台号*/
	//Cashier   string /*收银员编号*/
	//UiType    uint8  /*已经无用*/
	IsAi uint8 /*在AI处理时,标识AI处理的结果*/
	//Reserve   string /*保留*/
	Store string /*收银台号相应的MAC地址*/
	//Barcode   string /*条形码,已经无用*/
	Confirmed uint8
}

func TestGetHex(t *testing.T) {
	file, err := os.Open("I-01-20220601034054-0.jpg")
	if err != nil {
		t.Log(err)
	}
	fileInfo, err := file.Stat()
	size := fileInfo.Size()
	_, err = file.Seek(size-153, 1)
	if err != nil {
		t.Log(err)
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
	} else if buf[1] == byte(255) && buf[2] == byte(217) {
		chunks = append(chunks, buf[3:]...)
	}

	//fmt.Println(chunks)
	//result := hex.EncodeToString(chunks)
	//hex_data, _ := hex.DecodeString(result)
	//// 将 byte 转换 为字符串 输出结果
	//t.Log(string(hex_data))
	chunks2 := make([]byte, 0)
	for i := range chunks {
		if chunks[i] == 0 {
			continue
		}
		chunks2 = append(chunks2, chunks[i])
	}
	pictureInfo := PictureInfo1{}
	var cmd_info *model.GetPictureInfo
	cmd_info = (*model.GetPictureInfo)(byte_slice_to_struct(chunks2))
	//t.Log(cmd_info)
	pictureInfo.FileName = fmt.Sprintf("%s", cmd_info.FileName)
	pictureInfo.AlertTime = fmt.Sprintf("%s", cmd_info.AlertTime)
	pictureInfo.Lane = cmd_info.Lane
	testtime := fmt.Sprintf("%s", cmd_info.AlertTime)
	time, _ := strconv.Atoi(testtime)
	pictureInfo.Store = "test"
	//time, _ := strconv.ParseInt(testtime, 10, 64)
	t.Log(time)
	pictureInfo.IsAi = cmd_info.IsAi

	//pictureInfo.Store = fmt.Sprintf("%s", cmd_info.Store)

	pictureInfo.Confirmed = cmd_info.Confirmed

	t.Log(pictureInfo)

}
func byte_slice_to_struct(data []byte) unsafe.Pointer {
	return unsafe.Pointer((*reflect.SliceHeader)(unsafe.Pointer(&data)).Data)
}
