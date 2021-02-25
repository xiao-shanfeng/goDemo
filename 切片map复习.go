package main

import "fmt"

func review() {
	// 切片
	s := make([]string,0,5)
	s = append(s,"北京","上海")
	fmt.Println(s)
	// map类型的切片
	mSlice := make([]map[string]int,0,5)
	m1 := map[string]int{
		"北京":34,
		"上海":45,
	}
	mSlice = append(mSlice,m1)
	fmt.Println(mSlice)



	// map
	// 切片类型的map
	m := make(map[string][]int)
	m["北京"] = make([]int,0,5)
	m["北京"] = append(m["北京"],1,2,3,4,5)
	fmt.Println(m)

}
