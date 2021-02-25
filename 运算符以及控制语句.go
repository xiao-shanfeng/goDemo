package main

import "fmt"

func yunsuan() {
	// 运算符

	// 1。算术运算符
	/*
	 +	-	*	/	%
	++、--在go语言里面算是语句而不是运算符
	 */
	a := 1
	b := 2
	fmt.Println(a + b) //3
	fmt.Println(a - b) //-1
	fmt.Println(a * b) //2
	fmt.Println(a / b) //0
	fmt.Println(a % b) //1



	// 2。关系运算符
	/*
	==	!=	>=	<=	>	<
	 */
	fmt.Println(a == b) // false
	fmt.Println(a != b) // true
	fmt.Println(a >= b) // flase
	fmt.Println(a <= b) // true
	fmt.Println(a > b) // false
	fmt.Println(a < b) // true



	// 3。逻辑运算符
	/*
	&&	||	!
	 */
	c := true
	d := false
	fmt.Println(c && d) // false
	fmt.Println(c || d) // true
	fmt.Println(!c) // false




	// 4。位运算符
	/*
	&	|	^	<<	>>
	 */
	e := 1
	f := 3
	fmt.Println(e & f) // 1
	fmt.Println(e | f) // 3
	fmt.Println(e ^ f) // 2
	fmt.Println(e >> 2) // 0
	fmt.Println(e << 2) // 4



	// 5。赋值运算符
	/*
	=	+=	-=	*=	/=	%=	<<=	<<=	&=	|=	^=
	 */







	// 流程控制
	// if判断
	flag := false

	if flag {
		fmt.Println("这是正确的")
	}else {
		fmt.Println("这是错误的")
	}

	ifDemo1()
	ifDemo2()



	// for循环
	for i:=0;i<10;i++ {
		fmt.Println(i)
	}
	i:=1
	for ; i < 10; {
		fmt.Println(i)
		i++
	}



	// for range遍历数组，切片，字符串，map以及通道
	/*
	1.数组，切片，字符串，返回的是索引和值
	2.map返回的是键和值
	3.通道返回的只有值
	 */





	// switch case
	switchDemo()


	// goto语句
	// goto语句通过标签进行代码间的无条件跳转。goto语句可以在快速跳出循环、避免重复退出上有一定的帮助。Go语言中使用goto语句能简化一些代码的实现过程
	gotoDemo()
	gotoDemo2()


	// break语句
	//break语句可以结束for、switch和select的代码块
	//break语句还可以在语句后面添加标签，表示退出某个标签对应的代码块，标签要求必须定义在对应的for、switch和 select的代码块上。
	breakDemo:
		for i := 0; i < 10; i++ {
			for j := 0; j < 10; j++ {
				if j == 2 {
					break breakDemo
				}
				fmt.Printf("%v-%v\n",i,j)
			}
		}
		fmt.Println("........")



	//continue语句可以结束当前循环，开始下一次的循环迭代过程，仅限在for循环内使用。
	//在 continue语句后添加标签时，表示开始标签对应的循环
	forloop:
		for i := 0; i < 10; i++ {
			for j := 0; j < 10; j++ {
				if i == 2 && j == 2 {
					continue forloop
				}
				fmt.Printf("%v-%v\n",i,j)
			}
		}







	//编写代码打印9*9乘法表。
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %-2d    ",j,i,i*j)
		}
		fmt.Println()
	}






















}

func gotoDemo() {
	var breakFlag bool
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				breakFlag = true
				break
			}
			fmt.Printf("%v-%v\n",i,j)
		}
		if breakFlag {
			break
		}
	}
}

func gotoDemo2() {
	for i := 0; i < 10; i++ {
		for j:=0;j<10;j++ {
			if j == 2 {
				goto breakTag
			}
			fmt.Printf("%v-%v\n",i,j)
		}
	}
	return
	breakTag:
		fmt.Println("结束for循环")
}

func ifDemo1()  {
	// 全局变量
	score := 65
	if score > 90 {
		fmt.Println("优秀")
	}else if score > 60 {
		fmt.Println("良好")
	}else {
		fmt.Println("不及格")
	}
	fmt.Println(score) // 65
}

func ifDemo2() {
	// 局部变量
	if score := 65; score > 90 {
		fmt.Println("优秀")
	}else if score > 60 {
		fmt.Println("良好")
	}else {
		fmt.Println("不及格")
	}
	//fmt.Println(score) // 报错
}

func switchDemo() {
	n := 3
	switch n {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	case 3:
		fmt.Println("3")
	default:
		fmt.Println("重新输入")
	}
}