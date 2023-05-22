package model

type GetPictureInfo struct {
	FileName  [25]byte /*图片文件名 ASCII码 格式:I-收银台号-yyyyMMddhhmmss.jpg*/
	AlertTime [14]byte /*图片的报警时间 ASCII码 格式:yyyyMMddhhmmss*/
	Lane      uint8    /*收银台号*/
	//Cashier   [21]byte /*收银员编号*/
	//UiType    uint8    /*已经无用*/
	IsAi uint8 /*在AI处理时,标识AI处理的结果*/
	//Reserve   [33]byte /*保留*/
	Store [18]byte /*收银台号相应的MAC地址*/
	//Barcode   [29]byte /*条形码,已经无用*/
	Confirmed uint8
}

type PictureInfo struct {
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
