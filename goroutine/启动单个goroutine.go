package main

import (
	"fmt"
	"time"
)

func hello1() {
	fmt.Println("Hello Goroutine!")
}

func a() {
	go hello1()// 启动另外一个goroutine去执行hello函数
	fmt.Println("main Goroutine done!")
	time.Sleep(time.Second)
}
