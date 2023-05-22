package model

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/*
查询对应z_images_店的图片信息
* 当adjustment不为空时,说明已经经过人工调整过,即服务端和前端识别的结果不一样
* 保存图片和数据库识别的结果
dbname
*/
var Db *gorm.DB
var Db2 *gorm.DB
var err error

func InitDb1() {
	dsn := "readonly:ReadOnly)!@9@tcp(103.44.243.111:13306)/cartvumedia?charset=utf8mb4&parseTime=True&loc=Local"
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("db1 open error: ", err)
	}

}

func InitDb2() {
	dsn := "wws:Wuwansi202#@tcp(gz-cynosdbmysql-grp-go36exdn.sql.tencentcdb.com:25043)/bobsystems?charset=utf8mb4&parseTime=True&loc=Local"
	Db2, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("db2 open error:")
	}

}
