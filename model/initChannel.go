package model

var DownloadChan = make(chan string, 10)
var CompareChan = make(chan string, 10)
var CompareChanSig = make(chan string, 10)
var DecompressChan = make(chan string, 10)
var DecompressChanSig = make(chan string, 10)
var GetHexChan = make(chan string, 10)
var GetHexChanSig = make(chan string, 10)
var DataChan = make(chan PictureInfo, 1000)
var PublicChan = make(chan string, 100)

func AddStringToChannel(ch chan string, data string) {
	ch <- data
}
func AddDataToChannel(ch chan PictureInfo, data PictureInfo) {
	ch <- data
}
func GetDataChan() chan PictureInfo {
	return DataChan
}
func CloseChannel() {
	close(DownloadChan)
	close(CompareChan)
	close(DecompressChan)
	close(GetHexChan)
	close(DataChan)
	close(PublicChan)
}
