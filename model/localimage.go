package model

import "time"

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
	Image_create_time   string    `gorm:"iamge_create_time"` //数据表结构是bigint 疑似要把时间转换
	Confirm             uint8     `gorm:"confirm"`
	Edge_forward_type   uint8     `gorm:"edge_forward_type"`
	Server_forward_type uint8     `gorm:"server_forward_type"`
	Manual_check_type   uint8     `gorm:"manual_check_type"`
}
