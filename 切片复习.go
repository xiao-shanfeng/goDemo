package main

import "fmt"

func sliceFuXi() {
	// go指针
	// go语言中指针不能偏移和运算，两个符号：&取地址*根据地址取值

	// 基本语法：ptr := &v	v代表被取地址的变量（类型是T），ptr表示接收地址的变量（类型是*T）
	a := 10
	b := &a
	fmt.Printf("a:%d  ptr:%p\n",a,&a)
	fmt.Printf("b:%p  b:%T\n",b,b)
	fmt.Println(&b)


	// 指针取值
	c := *b
	fmt.Printf("c type:%T,c value:%d\n",c,c)
	/*变量、指针地址、指针变量、取地址、取值的相互关系和特性如下：
	对变量进行取地址（&）操作，可以获得这个变量的指针变量。
	指针变量的值是指针地址。
	对指针变量进行取值（*）操作，可以获得指针变量指向的原变量的值。*/
	modify1(a)
	fmt.Println(a)
	modify2(&a)
	fmt.Println(a)












	//new和make的区别
	// new和make主要是用来分配内存的
	//demo()

	// 1。new() :	func new(Type) *Type	type表示类型，*type表示类型指针
	newDemo()

	// 2。make() :	func make(t Type,size ...IntegerType) Type
	// make只用于slice，map，chan的内存创建，而且返回的是类型本身，而不是指针类型，因为这三种类型本身就是引用类型，没有必要返回他们的指针了
	makeDemo()






	// 总结:
	/*
	new与make的区别:
		1.二者都是用来做内存分配的。
		2.make只用于slice、map以及channel的初始化，返回的还是这三个引用类型本身；
		3.而new用于类型的内存分配，并且内存对应的值为类型零值，返回的是指向类型的指针。
	*/

	fmt.Println("正弦图像")
	runSin()


}

func modify1(x int) {
	x = 100
}
func modify2(x *int) {
	*x = 100
}

func demo() {
	var a *int
	fmt.Println(a)
	*a = 100
	fmt.Println(*a)

	var b map[string]int
	b["沙河"] = 100
	fmt.Println(b)
}

func newDemo() {
	a := new(int)
	b := new(bool)
	fmt.Printf("%T %v\n",a,a)
	fmt.Printf("%T %v\n",b,b)


	var c *int
	c = new(int)
	*c = 100
	fmt.Println(*c)

}

func makeDemo() {
	var b map[string]int
	b = make(map[string]int,10)
	b["沙河"] = 100
	fmt.Println(b)
}
