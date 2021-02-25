package main

import (
	"fmt"
	"math"
	"strings"
)

func leixing() {
	// go语言中分两大数据类型，基本数据类型：整型，浮点型，布尔型，字符串型，复合数据类型：指针，数组，切片，通道，结构体，map

	// 整型，分为int8，int16，int32，int64以及相对应的无符号的，uint8时byte型

	//注意： 在使用int和uint类型时，不能假定它是32位或64位的整型，而是考虑int和uint可能在不同平台上的差异。
	//注意事项 获取对象的长度的内建len()函数返回的长度可以根据不同平台的字节长度进行变化。实际使用中，切片 或 map 的元素数量等都可以用int来表示。在涉及到 二进制传输、读写文件 的结构描述时，为了保持文件的结构不会受到不同编译目标平台字节长度的影响，不要使用int和 uint。

	n := 10
	fmt.Printf("10的二进制%b\n",n)
	fmt.Printf("10的八进制%o\n",n)
	fmt.Printf("10的十进制%d\n",n)
	fmt.Printf("10的十六进制%x\n",n)



	//浮点型，有两种float32，float64
	m := 3.1415926
	fmt.Println(m)
	fmt.Printf("float32的最大范围%f，faloat64的最大范围%f\n",math.MaxFloat32,math.MaxFloat64)
	fmt.Printf("%.8f\n",math.Pi) //%f默认是小数位6位，

	//复数，有两种complex64	complex28
	//复数有实部和虚部，complex64的实部和虚部为32位，complex128的实部和虚部为64位。
	//var c1 complex64
	//c1 = 3 + 4i
	//fmt.Println(c1)


	//布尔型，有true和false
	/*
	注意：
		布尔类型变量的默认值为false。
		Go 语言中不允许将整型强制转换为布尔型.
		布尔型无法参与数值运算，也无法与其他类型进行转换。
	*/
	flag := true
	var p bool
	fmt.Println(flag,p)




	// 字符串，字符串内部实现使用utf-8编码，用""包括
	s1 := "这是字符串"
	fmt.Println(s1)

	// 字符串转义符 \r 回车符	\t 制表符	\n 换行符	\' 单引号	\" 双引号	\\ 反斜杠

	// 多行字符串，使用反引号（`），反引号里面转义字符将会原样输出
	s2 := `
	第一行
	第二行
	第三行
`
	fmt.Println(s2)


	// 字符串的常用函数
	// len() 求字符串的长度
	s3 := "tom"
	fmt.Println(len(s3))

	// +或fmt.Sprintf 拼接字符串
	s4 := "h,e,l,l,o "
	fmt.Println(s4 + " " + s3)
	str := fmt.Sprint(s4," ",s3)
	fmt.Println(str)

	//split分割字符串
	s5 := strings.Split(s4,",")		// 把s4字符串按照,进行分割
	fmt.Println(s5)
	s6 := strings.SplitN(s4,",",2)		// 把s4字符串按照,进行分割为两个
	fmt.Println(s6)
	s7 := strings.SplitAfter(s4,",")		// 把s4字符串按照,进行分割，它与split差别是它包含分割符，split不包含分割符
	fmt.Println(s7)
	s8 := strings.SplitAfterN(s4,",",2)		// 把s4字符串按照,进行分割为两个
	fmt.Println(s8)


	// 包含 contains
	s9 := strings.Contains(s4, "a") // 判断s4中是否包含a
	fmt.Println(s9)
	s10 := strings.ContainsAny(s4,"") // 判断s4中是否包含"a,"中任意一个字符，如果chars为空返回false
	fmt.Println(s10)

	// join操作 strings.Join(a []string, sep string) //用sep把a拼接起来
	s11 := strings.Join([]string{"a","b","c"},",")
	fmt.Println(s11)


	// 字符串出现的位置
	n1 := strings.Index(s4,"l") // 返回substr(e)在s4中首次出现的位置（从0开始数）,如果找不到，则返回 -1，如果 sep 为空，则返回 0。
	fmt.Println(n1) // 4
	n2 := strings.LastIndex(s4,"l") // 返回substr(e)在s4中最后一次出现的位置（从0开始数）
	fmt.Println(n2) // 6
	fmt.Println(strings.IndexAny(s4,"l,")) // 返回字符串 chars 中的任何一个字符在字符串s4中第一次出现的位置
	fmt.Println(strings.IndexByte(s4, 'o'))
	// 返回字符串s4中第一个满足匿名函数的字符的位置
	fmt.Println(strings.IndexFunc(s4, func(r rune) bool {
		return r == 'l' || r == ','
	}))
	fmt.Println(strings.IndexRune(s4, 'o'))
	fmt.Println(strings.HasPrefix(s4,"h,")) // 判断s4是否以prefix开头
	fmt.Println(strings.HasSuffix(s4,",o ")) // 判断s4是否以suffix结尾


	/*
	组成每个字符串的元素叫做“字符”，可以通过遍历或者单个获取字符串元素获得字符。 字符用单引号（’）包裹起来
	go语言中有两种，一个是uint8（byte）代表一个ascii码的一个字符，一个是rune（int32）代表一个utf-8字符
	 */
	c1 := '中'
	c2 := 'x'
	fmt.Println(c1,c2)

	traversalString()
	//字符串底层是一个byte数组，所以可以和[]byte类型相互转换。
	//字符串是不能修改的 字符串是由byte字节组成，所以字符串的长度是byte字节的长度。
	//rune类型用来表示utf8字符，一个rune字符由一个或多个byte组成。

	changString()







	// go 类型转换
	/*
	go语言里面类型转换只有强制类型转换，没有隐式类型转换，只能在两个类型支持转换的时候使用
	 */
	sqrtDemo()

	countStr()
}

func sqrtDemo() {
	a,b := 12,13
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

func traversalString() {
	s := "你好，Go!"
	for i := 0; i < len(s); i++ {
		fmt.Printf("%v(%c) ",s[i],s[i])
	}
	fmt.Println()
	for _, c := range s {
		fmt.Printf("%v(%c) ",c,c)
	}
	fmt.Println()
}


// 修改字符串
/*
要修改字符串，需要先将其转换成[]rune或[]byte，完成后再转换为string。无论哪种转换，都会重新分配内存，并复制字节数组。
 */
func changString() {
	s := "big"
	byteS1 := []byte(s)
	byteS1[0] = 'p'
	fmt.Println(byteS1)
	fmt.Println(string(byteS1))

	s2 := "白萝卜"
	runeS1 := []rune(s2)
	runeS1[0] = '红'
	fmt.Println(runeS1)
	fmt.Println(string(runeS1))

}



/*
编写代码统计出字符串"hello沙河小王子"中汉字的数量。
 */

func countStr() {
	s := "hello沙河小王子"
	var count int
	for _, c := range s {
		if c > 127 {
			count++
		}
	}
	fmt.Println(count)
}