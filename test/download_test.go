package test

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"testing"
)

// download path :http:images.checkpointcn.com
// http://images.checkpointcn.com/202206/002001-202206.tgz
func TestDownLoad(t *testing.T) {
	filepath := "http://images.checkpointcn.com/202206/002235-202206.tgz"
	resp, err := http.Get(filepath)
	if err != nil {
		t.Log("Download error:", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		fmt.Println("Download error")
	}
	localFilePath := "002235-202206.tgz"
	file, err := os.Create(localFilePath)
	if err != nil {
		t.Log("create file error", err)
	}
	defer file.Close()
	writesize, err := io.Copy(file, resp.Body)
	if err != nil {
		t.Log("copy errr", err)
	}
	t.Log("write size:", writesize)

}
