package main

import "fmt"

type Invoker interface {
	Call(interface{})
}

type Struct struct {
	
}


func (s *Struct) Call(p interface{})  {
	fmt.Println("from struct",p)
}

type FuncCaller func(interface{})

func (f FuncCaller)Call(p interface{}) {
	f(p)
}

func RunStruct() {
	// 结构体实现接口
	var invoker Invoker
	s := new(Struct)
	invoker = s
	invoker.Call("hello")


	// 函数实现接口
	invoker = FuncCaller(func(v interface{}) {
		fmt.Println("from function",v)
	})

	invoker.Call("hello")
}
