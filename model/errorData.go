/*
Create time at 2023年5月21日0021上午 10:11:31
Create User at Administrator
*/

package model

var SuccessPath []string
var SuccessDecompress []string
var SuccessGetHex []string

var FatalPath []string
var FatalDecompress []string
var FatalGetHex []string
var FatalRemove []string

func AddData(dataArray []string, data string) {
	dataArray = append(dataArray, data)
}
