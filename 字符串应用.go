package main


import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
	"time"
	"unicode/utf8"
)

type ChipType int
const (
	None ChipType = iota
	CPU
	GPU
)

func stringFun() {
	fmt.Println("字符串应用")
	// 1。字符串长度
	str := "hello,Go!"
	str1 := "忍者"
	fmt.Println(len(str))	//9
	fmt.Println(len(str1))	//6
	/*
	由于go语言的字符串是以utf-8格式保存的，每个中文占3个字节，所以len()获取两个中文是6个字节
	如果想要按照字符个数计算，需要使用utf-8包提供的RuneCountInString()函数来统计Uncode字符数量
	 */
	fmt.Println(utf8.RuneCountInString(str1))	//2
	str2 := "你好，Go"
	fmt.Println(utf8.RuneCountInString(str2))	//5
	/*
	总结 :
		ASCII字符串长度使用len()函数。
		Unicode 字符串长度使用 utf8.RuneCountlnString()函数。
	 */




	//2。遍历字符串
	temp := "狙击 start"
	for i := 0; i < len(temp); i++ {
		fmt.Printf("ascii: %c	%d\n",temp[i],temp[i])//中文显示乱码
	}
	for _, v := range temp {
		fmt.Printf("unicode: %c	%d\n",v,v)
	}
	/*
	总结:
		ASCII字符串遍历直接使用下标。
		Unicode字符串遍历用 forrange,
	 */





	//3。获取字符串的某一段字符
	tracer := "死神来了，死神bye bye"
	comma := strings.Index(tracer,"，")// 查找"，"在字符串tracer里面的ASCII码位置
	fmt.Println(comma,tracer[comma:])
	pos := strings.Index(tracer[comma:],"死神")
	fmt.Println(pos)
	fmt.Println(tracer[comma+pos:])



	//4。修改字符串
	// 由于go语言中字符串无法直接修改每一个字符元素，只能通过重新构造新的字符串并赋值给原来的字符串变量
	angel := "Heros never die"
	angelBytes := []byte(angel)
	for i := 5; i < 10; i++ {
		angelBytes[i] = ' '
	}
	fmt.Println(string(angelBytes))
	/*
	总结:
		Go 语言的字符串是不可变的。
		修改字符串时，可以将字符串转换为 []byte进行修改。
		[]byte和白ing可以通过强制类型转换互转。
	 */





	//5。连接字符串
	a := "你好"
	b := "go"
	fmt.Println(a + b)
	hammer := "吃我一锤"
	sickle := "死吧"

	// 定义一个字节缓冲
	var stringBuilder bytes.Buffer

	// 将字符串写入缓冲
	stringBuilder.WriteString(hammer)
	stringBuilder.WriteString(sickle)

	// 将缓冲以字符串形式输出
	fmt.Println(stringBuilder.String())






	// 6。格式化
	var progress = 2
	var target = 8
	title := fmt.Sprintf("已采集%d个药草，还需要%d个完成任务",progress,target)
	fmt.Println(title)

	pi := 3.14159
	variant := fmt.Sprintf("%v %v %v","月球基地",pi,true)//按值原来的值输出
	fmt.Println(variant)

	profile := &struct {
		Name string
		HP int
	}{
		Name: "rat",
		HP: 150,
	}
	fmt.Printf("使用'%%+v' %+v\n",profile)//在%v的基础上，对结构体字段名称和值进行展开
	fmt.Printf("使用'%%#v' %#v\n",profile)//输出Go语言语法格式的值
	fmt.Printf("使用'%%T' %T\n",profile)//输出Go语言语法格式的类型和值
	/*
	%b二进制	%o八进制	%d十进制	%x十六进制
	%X十六进制字母大写	%UUnicode字符	%f浮点数	%p指针，以十六进制显示
	 */

	encoding()

	findIniFoldValue()






	//7。常量
	//7.1。枚举
	type Weapon int
	const (
		Arrow Weapon = iota
		Shuriken
		SniperRifle
		Rifle
		Blower
	)
	fmt.Println(Arrow,Shuriken,SniperRifle,Rifle,Blower)
	var weapon Weapon = Blower
	fmt.Println(weapon)

	const (
		FlagNone = 1 << iota
		FlagRed
		FlagGreen
		Flagblue
	)
	fmt.Printf("%d %d %d %d\n",FlagNone,FlagRed,FlagGreen,Flagblue)
	fmt.Printf("%b %b %b %b\n",FlagNone,FlagRed,FlagGreen,Flagblue)


	fmt.Printf("%s %d\n",CPU,CPU)





	//8。类型别名
	// 类型别名是Go1.9版本添加的功能，主要用于代码升级以及迁移中类型的兼容问题

	// go1.9之前内建类型定义代码
	/*
	type byte uint8
	type rune int32
	 */
	// go1.9之后
	/*
	type byte = uint8
	type rune = int32
	 */

	//类型别名与类型定义
	// 类型别名格式：type TypeAlias = Type
	// 类型定义：   type TypeAlias Type

	type MyTime = time.Time
	var t MyTime = time.Now()
	fmt.Println(t)

	type D = int
	type I int
	v := 100
	var d D = v
	//var i I = v	error 因为I和int是两个类型
	fmt.Println(d)





	var cat Vehicle
	cat.FakeBrand.show()
	ta := reflect.TypeOf(cat)
	for i := 0; i < ta.NumField(); i++ {
		f := ta.Field(i)
		fmt.Printf("FiledName:%v,FiledType:%v\n",f.Name,f.Type.Name())
	}
}

// 结构体中使用类型别名

// 定义商标结构
type Brand struct {

}

func (b Brand)show()  {

}

type FakeBrand = Brand

type Vehicle struct {
	FakeBrand
	Brand
}

//7.2。将枚举转换成字符串
func (c ChipType)String() string {
	switch c {
	case None:
		return "None"
	case CPU:
		return "CPU"
	case GPU:
		return "GPU"
	}
	return "N/A"
}