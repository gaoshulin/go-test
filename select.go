package main

import (
	"fmt"
	"time"
)

func test1(ch chan string) {
	time.Sleep(time.Second * 5)
	ch <- "test1"
}
func test2(ch chan string) {
	time.Sleep(time.Second * 2)
	ch <- "test2"
}

func main() {
	// 2个管道
	ch1 := make(chan string)
	ch2 := make(chan string)

	// 跑2个子协程，写数据
	go test1(ch1)
	go test2(ch2)

	// 用select监控
	select {
	case s1 := <-ch1:
		fmt.Println("s1 =", s1)
	case s2 := <-ch2:
		fmt.Println("s2 =", s2)
	}

	// 判断管道有没有存满
	// 创建管道
	output1 := make(chan string, 10)
	// 子协程写数据
	go write(output1)
	// 取数据
	for s := range output1 {
		fmt.Println("res:", s)
		time.Sleep(time.Second)
	}
}

func write(ch chan string) {
	for {
		select {
		// 写数据
		case ch <- "hello":
			fmt.Println("write hello")
		default:
			fmt.Println("channel full")
		}
		time.Sleep(time.Millisecond * 500)
	}
}
