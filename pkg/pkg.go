package pkg

import "fmt"

var a = 10

const Mode = 1

type person struct {
	name string
}

type Students struct {
	Name string
	class string
}

type Payer interface {
	init()
	pay()
}

func Add(x,y int) int {
	return x + y
}

func age() {
	var Age = 18
	fmt.Println(Age)
}

func init() {
	fmt.Println("pkg init")
}