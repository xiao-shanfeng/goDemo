package main

import (
	"errors"
	"fmt"
	"strings"
)

var globalNum int = 34



var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))
)

func func1() {
	fmt.Println(initSum(3,5))
	initSum(3,4)
	sayHello()
	fmt.Println(initSum2(3))
	fmt.Println(initSum2(3,4,5))
	fmt.Println(calc(4,3))
	fmt.Println(calc2(4,3))
	fmt.Println(someFunc(""))
	fmt.Println(someFunc("a"))
	testGlobalVar()
	testLocalVar()
	//fmt.Println(x)//undefined: x
	testNum()

	var c calculation
	c = initSum
	fmt.Printf("%T\n",c)
	fmt.Println(c(2,3))

	fmt.Println(calc3(3,4,add))

	a,err := do("+")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(a)
	fmt.Println(a(3,4))
	fmt.Println(do("*"))




	//匿名函数
	add := func(a,b int) int {
		return a + b
	}
	fmt.Println(add(3,4))

	func () {
		fmt.Println("自执行函数")
	}()



	f := adder()
	fmt.Println(f(1))
	fmt.Println(f(2))
	fmt.Println(f(3))




	//defer语句

	fmt.Println("start")
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
	fmt.Println("end")

	fmt.Println(f1())//	5
	fmt.Println(f2())//	6
	fmt.Println(f3())//	5
	fmt.Println(f4())//	5
	f5()
	f6()


	// 注意：defer注册要延迟执行的函数时该函数所有的参数都需要确定其值
	x := 1
	y := 2
	defer calc4("AA",x,calc4("A",x,y))
	x = 10
	defer calc4("BB",x,calc4("B",x,y))
	y = 20


	// 内置函数
	// close，len，new，make，append，panic和recover
	fmt.Println("内置函数")
	funA()
	funB()
	funC()



	left := dispatchCoin()
	fmt.Println("剩下：", left)
	fmt.Println("每个人：",distribution)


}

//函数定义
/*
func 函数名(参数) (返回值){
	函数体
}
 */

func initSum(i, j int) int {
	return i + j
}

func sayHello() {
	fmt.Println("Hello")
}

// 可变参数
func initSum2(i ...int) int {
	fmt.Printf("%T %v",i,i)
	sum := 0
	for _, v := range i {
		sum += v
	}
	return sum
}

// 返回值
func calc(x, y int) (int, int) {
	sum := x + y
	sub := x - y
	return sum,sub
}
// 返回值命名
func calc2(x, y int) (sum, sub int) {
	sum = x + y
	sub = x - y
	return
}

func someFunc(x string) []int {
	if x == "" {
		return nil
	}
	return []int{1,2,3}
}

func testGlobalVar() {
	fmt.Printf("num=%d\n",globalNum)
}
func testNum() {
	globalNum := 300
	fmt.Println(globalNum)
}
func testLocalVar() {
	var x int64 = 100
	fmt.Println(x)
}

//函数类型
type calculation func(int,int) int

// 高阶函数
// 1。函数作为参数
func add(x,y int) int {
	return x + y
}

func calc3(x, y int, op func(int, int) int) int {
	return op(x,y)
}

// 2。函数作为返回值
func do(s string) (func(int, int) int,error) {
	switch s {
	case "+":
		return add,nil
	case "-":
		return sub,nil
	default:
		return nil,errors.New("无法识别操作符")
	}
}

func sub(x, y int) int {
	return x - y
}


// 闭包函数
//闭包=函数+引用环境

func adder() func(int) int {
	var x int
	return func(y int) int {
		x += y
		return x
	}
}

func f1() int {
	x := 5
	defer func() {
		x++
	}()
	return x
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5
}

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x
}

func f4() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5
}

func f5() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}

func f6() {
	for i := 0; i < 5; i++ {
		defer func() {
			fmt.Println(i)
		}()
	}
}

func calc4(index string,a,b int) int {
	ret := a + b
	fmt.Println(index,a,b,ret)
	return ret
}

func funA() {
	fmt.Println("func A")
}
func funB() {
	/*
	注意：
		1。recover()必须搭配defer使用。
		2。defer一定要在可能引发panic的语句之前定义。
	 */
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("recover in B")
		}
	}()
	panic("func B")
}
func funC() {
	fmt.Println("func C")
}


//作业
/*
你有50枚金币，需要分配给以下几个人：Matthew,Sarah,Augustus,Heidi,Emilie,Peter,Giana,Adriano,Aaron,Elizabeth。
分配规则如下：
a. 名字中每包含1个'e'或'E'分1枚金币
b. 名字中每包含1个'i'或'I'分2枚金币
c. 名字中每包含1个'o'或'O'分3枚金币
d: 名字中每包含1个'u'或'U'分4枚金币
写一个程序，计算每个用户分到多少金币，以及最后剩余多少金币？
程序结构如下，请实现 ‘dispatchCoin’ 函数
*/

func dispatchCoin() int {
	fmt.Println("==========================")
	for _, user := range users {
		sum := 0
		for _, s := range strings.Split(user, "") {
			switch {
			case s == "e" || s == "E":
				sum ++
			case s == "i" || s == "I":
				sum += 2
			case s == "o" || s == "O":
				sum += 3
			case s == "u" || s == "U":
				sum += 4
			}
		}
		distribution[user] = sum
		coins -= sum
	}
	return coins
}
