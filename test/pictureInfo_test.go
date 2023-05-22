package test

import (
	"bufio"
	"io"
	"os"
	"testing"
)

type PictureInfo struct {
	fileName  [25]byte /*图片文件名 ASCII码 格式:I-收银台号-yyyyMMddhhmmss.jpg*/
	alertTime [15]byte /*图片的报警时间 ASCII码 格式:yyyyMMddhhmmss*/
	lane      byte     /*收银台号*/
	cashier   [21]byte /*收银员编号*/
	uiType    byte     /*已经无用*/
	isAi      byte     /*在AI处理时,标识AI处理的结果*/
	reserve   [34]byte /*保留*/
	store     [20]byte /*收银台号相应的MAC地址*/
	barcode   [30]byte /*条形码,已经无用*/
	confirmed byte
}

//func byte_slice_to_struct(data []byte) unsafe.Pointer {
//	return unsafe.Pointer((*reflect.SliceHeader)(unsafe.Pointer(&data)).Data)
//}
func TestPictureInfo(t *testing.T) {
	file, err := os.Open("I-01-20220601034407-0.jpg")
	if err != nil {
		t.Log(err)
	}

	s1 := byte(255)
	s2 := byte(217)
	label := 0
	defer file.Close()
	r := bufio.NewReader(file)
	chunks := make([]byte, 0)
	buf := make([]byte, 100)
	for {
		n, err := r.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if label == 1 {
			chunks = append(chunks, buf[:n]...)
		}
		for i := 1; i < n; i++ {
			if buf[i] == s2 {
				if buf[i-1] == s1 {
					label = 1
					chunks = append(chunks, buf[i+1:n]...)
				}
			}
		}
		if 0 == n {
			break
		}
	}
	chunks2 := make([]byte, 0)
	for i := range chunks {
		if chunks[i] != 0 {
			chunks2 = append(chunks2, chunks[i])
		}

	}
	//fmt.Println(chunks2)
	//result := hex.EncodeToString(chunks2)
	//hex_data, _ := hex.DecodeString(result)
	//// 将 byte 转换 为字符串 输出结果
	//t.Log(string(hex_data))

	//se, err := os.ReadFile("I-01-20220601034407-0.jpg")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//test := util.GetHex(se)
	//log.Println(se)
	//me := make([]string, 0)
	//for _, i := range se {
	//	s := strconv.FormatInt(int64(i), 16)
	//	s = "0x" + s
	//	me = append(me, s)
	//}
	//sbyte := Hex2Dec(me)
	//log.Println(me)
	//var cmd_info *PictureInfo

	//cmd_info = (*PictureInfo)(byte_slice_to_struct1(chunks2))
	//fmt.Printf("%s", cmd_info.fileName)

}
