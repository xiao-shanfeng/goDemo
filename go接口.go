package main

import "fmt"

type Cat struct {

}

func (c Cat) Say() string {
	return "喵喵喵"
}

type Dog struct {

}

func (d Dog) Say() string {
	return "汪汪汪"
}

type Sayer interface {
	say()
}

func (c Cat) say() {
	fmt.Println("喵喵喵")
}
func (d Dog) say() {
	fmt.Println("汪汪汪")
}

type Mover interface {
	move()
}

type dog struct {
	
}

func (d *dog) move() {
	fmt.Println("狗会动")
}



// 类型与接口

type Dog2 struct {
	name string
}

func (d Dog2) say() {
	fmt.Printf("%s会汪汪汪～\n",d.name)
}

func (d Dog2) move() {
	fmt.Printf("%s会动\n",d.name)
}


func interfaceFun() {
	c := Cat{}
	fmt.Println("猫:",c.Say())
	d := Dog{}
	fmt.Println("狗:",d.Say())


	var x Sayer
	a := Cat{}
	b := Dog{}
	x = a
	x.say()
	x = b
	x.say()


	/*var y Mover
	e := dog{}
	y = e
	y.move()
	g := &dog{}
	y = g
	y.move()*/
	/*
	从上面的代码中我们可以发现，使用值接收者实现接口之后，不管是dog结构体还是结构体指针*dog类型的变量都可以赋值给该接口变量。
	因为Go语言中有对指针类型变量求值的语法糖，dog指针fugui内部会自动求值*fugui。
	*/

	/*var i Mover
	f := dog{}
	i = f
	i.move()
	j := &dog{}
	i = j
	i.move()*/
	/*
	此时实现Mover接口的是*dog类型，所以不能给x传入dog类型的wangcai，此时x只能存储*dog类型的值。
	 */



	var s Sayer
	var m Mover

	var d1 = Dog2{
		name: "旺财",
	}
	s = d1
	m = d1
	s.say()
	m.move()



	// 类型断言
	var l interface{}
	l = "沙河"
	fmt.Printf("%#v %T\n",l,l)
	v,ok := l.(string)
	if ok {
		fmt.Println(v)
	}else {
		fmt.Println("类型断言失败")
	}



}
