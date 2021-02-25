package main

import (
	"fmt"
	"time"
)

func main() {
	go running()
	var input string
	fmt.Scan(&input)
}

func running() {
	var timer int
	for  {
		timer++
		fmt.Println("tick",timer)
		time.Sleep(time.Second)
	}
}