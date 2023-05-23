package test

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

var ch2 = make(chan int, 10)
var ch1 = make(chan int, 10)

func test1() {
	fmt.Println("test1 go")
	ch2 <- 1
}
func test2() {
	for i := 0; i < 10; i++ {
		fmt.Println(1)

	}
	fmt.Println("test2 go")
	ch1 <- 2
}
func TestChannel(t *testing.T) {
	ch1 <- 3
	for {
		select {
		case c := <-ch1:
			fmt.Println(c)
			go test1()
		case c1 := <-ch2:
			fmt.Println(c1)
			go test2()
		default:
			time.Sleep(5 * time.Second)
			fmt.Println("wart")
		}
	}

}
func TestParse(t *testing.T) {
	s := "2022060106521"
	time, _ := strconv.ParseInt(s, 10, 64)
	t.Log(int64(time))
}

//var s = []string{"test1", "test2", "test3", "test4", "test5"}

var global = make(chan []string, 10)

func test3() {
	//for _, v := range s {
	//	fmt.Println(v)
	//}
	s := <-global
	fmt.Println(s)
	//wg.Done()
}
func TestAnyGoroutine(t *testing.T) {
	map1 := make(map[string][]string)
	map1["test1"] = []string{"2", "3", "4", "5"}
	map1["test2"] = []string{"6", "7", "8", "9"}
	map1["test3"] = []string{"10", "11", "12", "13"}
	//s := []string{"test1", "test2", "test3", "test4", "test5"}
	ch := make(chan string, 1)
	che := make(chan string, 1)
	ch <- "start"
	//wg := &sync.WaitGroup{}
	for {
		select {
		case msg := <-ch:
			fmt.Println(msg)
			for _, v := range map1 {
				global <- v
			}
			che <- "action"
		case msg1 := <-che:
			fmt.Println(msg1)
			for i := 0; i < 10; i++ {
				go test3()
			}

		default:
			che <- "end"
			time.Sleep(1 * time.Second)
		}
	}

}
func test4(s string) {
	fmt.Println(s)
	ch4 <- "ss"
	ch3 <- "ch3 ziji"
	time.Sleep(10 * time.Second)
}
func test5(s string) {
	fmt.Println(s)
	ch4 <- "ss"
	ch3 <- "ch3 ziji"

}

var ch3 = make(chan string, 1)
var ch4 = make(chan string, 1)

func TestGoroutine(t *testing.T) {

	ch3 <- "init"
	for {
		select {
		case s := <-ch3:
			go test4(s)
		case s1 := <-ch4:
			go test5(s1)
		}
	}
}
