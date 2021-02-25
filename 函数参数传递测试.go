package main

import "fmt"

type InnerData struct {
	a int
}

type Data struct {
	complex []int
	instance InnerData
	ptr *InnerData
}

func paramsTest() {
	in := Data{
		complex: []int{1,2,3},
		instance: InnerData{
			5,
		},
		ptr: &InnerData{
			1,
		},
	}
	fmt.Printf("in value: %+v\n",in)
	fmt.Printf("in ptr:%p\n",&in)

	out := passByValue(in)
	fmt.Printf("out value: %+v\n",out)
	fmt.Printf("out ptr: %p\n",&out)
}

func passByValue(inFunc Data) Data {
	// 输出参数的成员情况
	fmt.Printf("inFunc value:%+v\n",inFunc)

	// 打印inFunc的指针
	fmt.Printf("inFunc ptr:%p\n",&inFunc)

	return inFunc
}

