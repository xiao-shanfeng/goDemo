package main

import "fmt"


func test() {
	// 声明变量
	//var 变量名 类型
	var name string
	// var 变量名 类型 = XXX     初始化变量
	var tom,xiaoming = "Tom","小明"
	var m = 100// 类型可以根据初始化的值进行推到
	name = "Tom"

	// 短变量声明
	// 在函数内部可以使用:=的方式声明并初始化一个变量，不能在函数外部使用
	n := "100"
	fmt.Println("n:",n)

	a,_ := foo()
	fmt.Println(a)

	fmt.Println(name)
	fmt.Println(m)
	fmt.Println(tom,xiaoming)
	// 批量声明变量
	var (
		age int
		hobby string
	)
	fmt.Println(age,hobby)





	// 常量    定义必须赋值
	const PI = 3.1415926
	fmt.Println(PI)

	// 常量也可以批量初始化
	const (
		b = 40
		c = "timi"
	)
	fmt.Println(b,c)

	// 批量初始化时，后面变量省略值至和上面一行的值一样
	const (
		n1 = 100
		n2
		n3
		n4 = 500
		n5
		n6
	)
	fmt.Println(n1,n2,n3,n4,n5,n6)







	// iota 常量计数器，常用在常量表达式
	// iota在const关键字出现时将被重置为0。const中每新增一行常量声明将使iota计数一次(iota可理解为const语句块中的行索引)
	const (
		m1 = iota
		m2
		m3 = 100
		m4 = iota
	)
	fmt.Println(m1,m2,m3,m4) // 0  1  100  3
	/*const (
		p1 = iota
		_
		p2
		p3
	)
	fmt.Println(p1,p2,p3)*/

	const (
		p1 = iota
		p2 = 100
		p3 = iota
		p4
	)
	const p5 = iota
	fmt.Println(p1,p2,p3,p4,p5)

	// 定义一个数量级
	const (
		_ = iota
		KB = 1 << (10 * iota)
		MB
		GB
		TB
		PB
	)
	fmt.Println(KB,MB,GB,TB,PB)

	//多个iota定义在一行
	/*const (
		d,e = iota + 1,iota + 1
		f,g
	)
	fmt.Println(d,e,f,g) // 1 1 2 2*/
	const (
		d,e = iota + 1,iota + 2
		f,g
	)
	num,err := fmt.Println(d,e,f,g)

	fmt.Printf("%#v%#v",num,err)



}

func foo() (int,string){
	return 100,"tom"
}
