package test

import (
	"fmt"
	"testing"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/hints"
)

type StoreId struct {
	Id string
}
type RemoteImage struct {
	Id        int
	ImagePath string
	StoreId   string
	LaneNo    int
	AlarmTime time.Time
	Confirmed int
	Cashier   string
	BarCode   string
	Accounted int
	Type      int
	Node      string
	Backup    int
}
type LocalImage struct {
	Id                  int       `gorm:"id"`
	Create_time         time.Time `gorm:"create_time"`
	Update_time         time.Time `gorm:"update_time"`
	Image_name          string    `gorm:"image_name"`
	Store_no            string    `gorm:"store_no"`
	Lane_no             int       `gorm:"lane_no"`
	Image_create_time   int64     `gorm:"iamge_create_time"` //数据表结构是bigint 疑似要把时间转换
	Confirm             int       `gorm:"confirm"`
	Edge_forward_type   int       `gorm:"edge_forward_type"`
	Server_forward_type int       `gorm:"server_forward_type"`
	Manual_check_type   int       `gorm:"manual_check_type"`
}

func TestDbSelect(t *testing.T) {
	dsn := "readonly:ReadOnly)!@9@tcp(103.44.243.111:13306)/cartvumedia?charset=utf8mb4&parseTime=True&loc=Local"
	Db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("db1 open error: ", err)
	}
	remotemodel := RemoteImage{}
	Db.Table("z_images_002158").Where("Id=?", 2977924).Take(&remotemodel)
	fmt.Println("remotemodel:", remotemodel)
	dsn2 := "wws:Wuwansi202#@tcp(gz-cynosdbmysql-grp-go36exdn.sql.tencentcdb.com:25043)/bobsystems?charset=utf8mb4&parseTime=True&loc=Local"
	Db2, err := gorm.Open(mysql.Open(dsn2), &gorm.Config{
		SkipDefaultTransaction: false,
		Logger:                 logger.Default.LogMode(logger.Info),
		TranslateError:         true,
	})
	if err != nil {
		fmt.Println("db2 open error: ", err)
	}
	s := "test-name"
	localtemodel := LocalImage{}
	localtemodel.Id = remotemodel.Id
	localtemodel.Create_time = time.Now()
	localtemodel.Update_time = time.Now()
	localtemodel.Image_name = s

	localtemodel.Store_no = remotemodel.StoreId
	localtemodel.Lane_no = remotemodel.LaneNo
	//localtemodel.image_create_time = remotemodel.AlarmTime
	localtemodel.Image_create_time = 255555
	if remotemodel.Node == "audited" {
		localtemodel.Confirm = 1
	}
	if remotemodel.Node == "no" {
		localtemodel.Confirm = 0
	} else {
		localtemodel.Confirm = -1
	}
	localtemodel.Edge_forward_type = 1
	localtemodel.Server_forward_type = 1
	localtemodel.Manual_check_type = remotemodel.Type
	Db2.Table("bob_images").Save(&localtemodel)
}

type RemoteImage2 struct {
	ImagePath string
	Type      int
	Node      string
}

func TestDbSelect3(t *testing.T) {
	dsn := "readonly:ReadOnly)!@9@tcp(103.44.243.111:13306)/cartvumedia?charset=utf8mb4&parseTime=True&loc=Local"
	db3, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("db1 open error: ", err)
	}
	storeId := []StoreId{}
	db3.Table("storeinfo").Select("Id").Find(&storeId)

	remoteImage := []RemoteImage2{}
	query := "/CartvuMedia/FTP/Cartvu/Images/002235/I-02-20221120125530-0.jpg"
	db3.Table("z_images_002235").Select("ImagePath,Note,Type").Where("ImagePath=?", query).Find(&remoteImage)
	fmt.Println(remoteImage)
}

type RemoteImage3 struct {
	ImagePath string
	Type      string
}

func TestGetStoreInfo(t *testing.T) {
	dsn := "readonly:ReadOnly)!@9@tcp(103.44.243.111:13306)/cartvumedia?charset=utf8mb3&parseTime=True&loc=Local"
	db3, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("db1 open error: ", err)
	}
	var testMap = make(map[string][]RemoteImage3)
	storeId := []StoreId{}
	remoteImage := []RemoteImage3{}
	db3.Table("storeinfo").Select("Id").Find(&storeId)
	for i := range storeId {

		table := "z_images_" + storeId[i].Id

		db3.Table(table).Select("ImagePath,Type").
			Where("Note=?", "audited").
			Find(&remoteImage)
		testMap[table] = remoteImage
	}
	sqlDB, _ := db3.DB()
	// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(24 * time.Hour)
	fmt.Println(testMap)
}
func TestConnectDb(t *testing.T) {
	dsn := "readonly:ReadOnly)!@9@tcp(103.44.243.111:13306)/cartvumedia?charset=utf8mb4&parseTime=True&loc=Local"
	Db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
		Logger:                 logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		fmt.Println("db1 open error: ", err)
	}
	remotemodel := RemoteImage3{}
	Db.Table("z_images_002158").Select("ImagePath,Type,Note").Where("ImagePath=? and Note=?",
		"/CartvuMedia/FTP/Cartvu/Images/002158/I-14-20221123064703-0.jpg",
		"audited").
		Clauses(hints.UseIndex("ImagePath")).Find(&remotemodel)
	//Db.Table("z_images_002158").Where("Id=?", 2977924).Take(&remotemodel)
	if remotemodel.Type == "" {
		t.Log("type is empty")
	}
	fmt.Println("remotemodel:", remotemodel)
}

var s = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}

func tp() {
	if len(s) != 1 {
		s = s[1:]
	}

}
func TestURL(t *testing.T) {

	for _ = range s {
		test := s[0]
		tp()
		fmt.Println(test)
	}
}
