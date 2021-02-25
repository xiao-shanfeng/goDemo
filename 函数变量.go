package main

import "fmt"

func fire() {
	fmt.Println("fire")
}

func funcVar() {
	var f func()
	f = fire
	f()
}
