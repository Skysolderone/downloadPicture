package test

import (
	"fmt"
	"testing"
)

type set_cmd_info struct {
	Cmd_type [2]uint8
	Cmd_len  [2]uint8
	Data     [5]uint8
}

//func byte_slice_to_struct(data []byte) unsafe.Pointer {
//	return unsafe.Pointer((*reflect.SliceHeader)(unsafe.Pointer(&data)).Data)
//}
func TestStrcutToSlice(t *testing.T) {
	var net_data = []byte{0xA5, 0xA5, 0x04, 0x00, 0x11, 0x22, 0x33, 0x44, 0x55}
	var cmd_info *set_cmd_info

	cmd_info = (*set_cmd_info)(byte_slice_to_struct(net_data))

	fmt.Printf("cmd type:0x%x\n", cmd_info.Cmd_type)
	fmt.Printf("cmd len:0x%x\n", cmd_info.Cmd_len)
	fmt.Printf("cmd data:0x%x\n", cmd_info.Data)

}
