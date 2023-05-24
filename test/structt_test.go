package test

import (
	"fmt"
	"testing"

	"github.com/panjf2000/ants/v2"
)

type SuccessData struct {
	SuccessPath       []string
	SuccessDecompress []string
	SuccessGetHex     []string
}
type FatalData struct {
	FatalPath       []string
	FatalDecompress []string
	FatalGetHex     []string
	FatalRemove     []string
}

func CreateFatalData() FatalData {
	return FatalData{
		FatalPath:       []string{},
		FatalDecompress: []string{},
		FatalGetHex:     []string{},
		FatalRemove:     []string{},
	}
}
func CreateSuccessData() SuccessData {
	return SuccessData{
		SuccessPath:       []string{},
		SuccessDecompress: []string{},
		SuccessGetHex:     []string{},
	}
}

func TestTst(t *testing.T) {
	CreateSuccessData()
	CreateFatalData()
}

func TestPoolT(t *testing.T) {
	ch := make(chan string, 100)
	ch2 := make(chan string, 100)
	ch <- "start"
	p, _ := ants.NewPool(1000)
	defer p.Release()
	for {
		select {
		case <-ch:
			ch2 <- "end"
			//for i := 0; i < 100; i++ {
			p.Submit(func() {
				fmt.Println("test")
			})
			//}
		case <-ch2:
			ch <- "ch2 controller"
		}
	}
}
