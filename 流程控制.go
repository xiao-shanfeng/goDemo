package main

import "fmt"

func forLoop() {
	for i := 1; i < 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d ",j,i,i*j)
		}
		fmt.Println()
	}
	
	//遍历数组
	for k, v := range []int{1, 2, 3, 4, 5} {
		fmt.Println(k,v)
	}
	//遍历字符串
	for k, v := range "hello，你好" {
		fmt.Printf("key:%d,value:0x%x\n",k,v)
	}
	//遍历map
	m := map[string]int{
		"hello":100,
		"go":200,
	}
	for k, v := range m {
		fmt.Printf("key:%s,value:%d\n",k,v)
	}
	//遍历通道
	c := make(chan int)
	go func() {
		c <- 1
		c <- 2
		c <- 3
		close(c)
	}()
	for i := range c {
		fmt.Println(i)
	}
}
