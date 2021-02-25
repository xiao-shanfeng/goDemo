package main

import (
	"fmt"
	"sort"
)

func slice() {
	// 切片是一个拥有相同类型元素的可变长度的序列，支持自动扩容
	// 切片是引用类型，包含地址、长度、容量，

	// 切片的基本语法	var 变量名 []T(类型)
	var s1 []int	//声明一个int切片
	var s2 = []int{}	//声明一个int切片并初始化
	var s3 = []bool{true,false}	//声明一个bool切片并初始化
	var s4 = []bool{true,false}

	fmt.Printf("%T------%v\n",s1,s1)	//[]int------[]
	fmt.Println(s2)	//[]
	fmt.Println(s3)	//[true false]
	fmt.Println(s4)	//[true false]
	fmt.Println(s1 == nil)	//true
	fmt.Println(s2 == nil)	//false
	fmt.Println(s3 == nil)	//false
	//fmt.Println(s4 == s3)	//切片是引用类型，不支持直接比较，只能和nil比较


	//切片拥有自己的长度和容量，
	//我们可以通过使用内置的len()函数求长度，
	//使用内置的cap()函数求切片的容量。
	fmt.Println(len(s1))
	fmt.Println(cap(s1))


	//切片表达式
	// 1。简单切片表达式:指定low和high两个索引界限值的简单形式
	/*
	切片的底层实际上是一个数组，我我们可以基于数组通过切片表达式来得到切片。切片表达式里面的low和high表示一个范围（左包含、右不包含）low<=s<high
	 */
	a := [5]int{1,2,3,4,5}
	s := a[1:4]   //把a从索引为1的截到索引为3组成的
	fmt.Println(a)
	fmt.Printf("s:%v len:%v cap:%v\n",s,len(s),cap(s))

	fmt.Println(a[:])//[1 2 3 4 5]
	fmt.Println(a[2:])//[3 4 5]
	fmt.Println(a[:3])//[1 2 3]

	//对于数组或字符串，如果0 <=low<=high<=len(a)，则索引合法，否则就会索引越界（out of range）。
	//对切片再执行切片表达式时（切片再切片），high的上限边界是切片的容量cap(a)，而不是长度
	fmt.Printf("s:%v len:%v cap:%v\n",s,len(s),cap(s))
	s5 := s[1:3]
	// 对切片再切片时，底层数组还是同一个，并没有变
	fmt.Printf("s2:%v len:%v cap:%v\n",s5,len(s5),cap(s5))



	// 2。完整切片表达式：除了low和high索引界限之外还指定容量的形式
	/*
	a[low:high:max]
	上式子将会构造与简单表达式a[low:high]相同类型、相同长度和元素的切片，但是切片的容量是max-low。
	完整切片表达式只有low可以省略，默认是0
	*/
	b := [9]int{1,2,3,4,5,6,7,8,9}
	s6 := b[1:4:6] // [2 3 4] 3 5
	fmt.Printf("s6:%v len:%v cap:%v\n",s6,len(s6),cap(s6))

	// 完整切片表达式满足的条件是0<=low<=high<=max<=cap(a)















	//使用make函数构造切片，格式make([]T,size,cap)	T指切片的数据类型	size指切片的元素数量	cap指切片的容量

	slice1 := make([]int,3,5)
	/*
	上面代码中slice1的内部存储空间已经分配了5个，但实际上只用了3个。
	容量并不会影响当前元素的个数，所以len(slice1)返回3，cap(slice1)则返回该切片的容量。
	 */
	fmt.Printf("%T,%v,%v,%v\n",slice1,slice1,len(slice1),cap(slice1))

	// 判断切片是否为空，用len(s) == 0来判断，不要用s == nil
	fmt.Println(len(s2)==0,s2==nil)
	// 切片之间是不能比较的，不能使用==来比较两个切片是否含有全部相等元素。切片只能和nil做比较

	var s7 []int		//len:0	cap:0	s7==nil
	s8 := []int{}		//len:0	cap:0	s8!=nil
	s9 := make([]int,0)	//len:0	cap:0	s9!=nil
	fmt.Printf("len:%v cap:%v %v\n",len(s7),cap(s7),s7==nil)
	fmt.Printf("len:%v cap:%v %v\n",len(s8),cap(s8),s8==nil)
	fmt.Printf("len:%v cap:%v %v\n",len(s9),cap(s9),s9==nil)


	copy2()
	forloop()



	// append()方法为切片添加元素

	slice2 := []int{1,2,3}
	slice2 = append(slice2,4)
	slice2 = append(slice2,5,6,7)
	slice3 := []int{8,9}
	slice2 = append(slice2,slice3...)
	fmt.Println(slice2)

	//通过var声明的零值切片可以在append()函数直接使用，无需初始化。
	var slice4 []int
	slice4 = append(slice4,1)
	fmt.Println(slice4)




	/*
	每个切片会指向一个底层数组，这个数组的容量够用就添加新增元素。
	当底层数组不能容纳新增的元素时，切片就会自动按照一定的策略进行“扩容”，此时该切片指向的底层数组就会更换。
	“扩容”操作往往发生在append()函数调用时，所以我们通常都需要用原变量接收append函数的返回值。
	 */

	var numSlice []int
	for i := 0; i < 10; i++ {
		numSlice = append(numSlice,i)
		fmt.Printf("%v	len:%v	cap:%v	ptr:%p\n",numSlice,len(numSlice),cap(numSlice),numSlice)
	}

	/*
	使用copy()函数复制切片
	 */
	copy1()



}

func copy1() {
	a := []int{1,2,3,4}
	b := a
	fmt.Println(a,b)
	b[0] = 100
	fmt.Println(a,b)
	/*
	由于切片是引用类型，a和b都指向一个同一个内存地址，修改b的值，a也会同时修改
	使用copy()函数可以将一个切片的数据复制到另一个切片空间中，格式：copy(destSlice,srcSlice []T)	srcSlice数据来源切片	destSlice目标切片
	 */
	c := make([]int,5,10)
	copy(c,a)
	fmt.Println(a,c)
	c[0] = 1
	fmt.Println(a,c)







	// 从切片删除元素

	numSlice := []int{1,2,3,4,5,6,7,8,9}
	// 要删除索引为2的元素
	numSlice = append(numSlice[:2],numSlice[3:]...)
	fmt.Println(numSlice)

	//总结
	//总结一下就是：要从切片a中删除索引为index的元素，操作方法是a = append(a[:index], a[index+1:]...)











	// 作业
	homework()
	//请使用内置的sort包对数组var a = [...]int{3, 7, 8, 9, 1}进行排序（附加题，自行查资料解答）。
	sortHomework()

}


func sortHomework() {
	var a = [...]int{3,7,8,9,1}
	sort.Ints(a[:])
	fmt.Println(a)
}

//切片赋值拷贝
func copy2() {
	s1 := make([]int,3)
	s2 := s1	// s1和s2共用一个底层数组
	s2[0] = 100
	fmt.Println(s1)
	fmt.Println(s2)
}


// 切片遍历，和数组遍历一样两种形式，一是索引遍历一是for range
func forloop() {
	s1 := []int{1,2,3,4,5,6,7}
	for i := 0; i < len(s1); i++ {
		fmt.Printf("%v-%v  ",i,s1[i])
	}
	fmt.Println()
	for i, v := range s1 {
		fmt.Printf("%v-%v  ",i,v)
	}
	fmt.Println()
}

func homework() {
	var a = make([]string, 5, 10)
	for i := 0; i < 10; i++ {
		a = append(a, fmt.Sprintf("%v", i))
	}
	fmt.Println(a)
}