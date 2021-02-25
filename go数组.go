package main

import "fmt"

func shuzu() {
	// go语言中数组，数组是同一种数据类型元素的集合。 在Go语言中，数组从声明时就确定，使用时可以修改数组成员，但是数组大小不可变化。
	var arr [3]int	// 定义一个长度为3元素类型为int的数组arr
	fmt.Println(arr)

	// 数组定义	var 数组变量名 [元素数量]T(类型)	var arr [3]int
	// 数组的长度是数组类型的一部分，一旦长度确定了，就不能修改了
	/*
	var a [5]int
	var b [10]int
	a = b	//Cannot use 'b' (type [10]int) as type [5]int

	数组可以通过下标进行访问，下标是从0开始，最后一个元素下标是：len-1，
	访问越界（下标在合法范围之外），则触发访问越界，会panic。
	*/

	// 数组初始化
	// 1。初始化列表
	var testArr [3]int
	var numArr = [3]int{1,2}
	var cityArr = [3]string{"北京","上海","杭州"}
	fmt.Println(testArr)
	fmt.Println(numArr)
	fmt.Println(cityArr)

	// 2。让编辑器通过初始值来推断数组长度
	var numArr1 = [...]int{1,2}
	fmt.Println(numArr1)
	fmt.Printf("%T\n",numArr1)

	// 3。通过索引值的方式来初始化
	a := [...]int{1:3, 4:2, 7:9} //[0 3 0 0 2 0 0 9]
	fmt.Println(a)
	fmt.Printf("%T\n",a)


	// 数组遍历
	arr1 := []string{"上海","杭州","长沙","无锡"}
	for i := 0; i < len(arr1); i++ {
		fmt.Printf("%v	",arr1[i])
	}
	fmt.Println()
	for index, str := range arr1 {
		fmt.Printf("索引是%d,值是%v\n",index,str)
	}



	// 多维数组
	// 定义一个三行两列的数组
	b := [3][2]string{
		{"上海","杭州"},
		{"北京","大连"},
		{"西安","河北"},
	}
	fmt.Println(b)
	fmt.Println(len(b))
	fmt.Println(b[2][1]) // 根据索引值取值
	// 遍历
	for i := 0; i < len(b); i++ {
		for j := 0; j < len(b[0]); j++ {
			fmt.Printf("%v  ",b[i][j])
		}
		fmt.Println()
	}
	for _, stringArr := range b {
		for _, strArr := range stringArr {
			fmt.Printf("%v	",strArr)
		}
		fmt.Println()
	}

	// 多维数组只有第一层通过...来推导数组长度
	c := [...][2]int{
		{1,3},
		{2,4},
	}
	/*c := [2][...]int{
		{1,3},
		{2,4},
	}// use of [...] array outside of array literal*/
	fmt.Println(c)
	fmt.Printf("%T\n",c)





	//数组是值类型

	d := [...]int{1,2,3}
	fmt.Println("变化之前的，",d)
	modifyArr(d)	// 传过去的是复制d之后的一份
	modifyArr2(&d)	// 传过去的是d的地址，是同一份变量
	fmt.Println("变化之后的，",d)


	//[n]*T表示指针数组，*[n]T表示数组的指针 。
	f := 3
	e := [...]*int{&f}
	g := &[...]int{1,2,3}
	fmt.Println(e,g)





	//作业
	/*1。求数组[1, 3, 5, 7, 8]所有元素的和
	2。找出数组中和为指定值的两个元素的下标，比如从数组[1, 3, 5, 7, 8]中找出和为8的两个元素的下标分别为(0,3)和(1,2)。*/
	numArr2 := [...]int{1,3,5,7,8}
	sum := 0
	for _, arr := range numArr2 {
		sum += arr
	}
	fmt.Println(sum)

	findIndex(numArr2, 8)

	
}

func findIndex(arr [5]int,c int) {
	for i := 0; i < len(arr); i++ {
		for j := i+1; j < len(arr); j++ {
			if arr[i] + arr[j] == c {
				fmt.Printf("(%v,%v)",i,j)
			}
		}
	}
}


func modifyArr(arr [3]int) {
	// arr和d不是同一个变量，是两个变量
	arr[1] = 100
}

func modifyArr2(arr *[3]int)  {
	// arr是d的地址，通过*arr取到变量d
	(*arr)[1] = 100
}

func modifyArr3(arr [2][3]int) {
	arr[0][0] = 3
}