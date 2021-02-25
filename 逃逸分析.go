package main

import "fmt"

func runFenXi() {
	var a int

	void()

	fmt.Println(a,dummy(0))
}

func dummy(c int) int {
	var b int
	b = c
	return b
}

func void() {

}