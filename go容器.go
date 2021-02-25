package main

import (
	"container/list"
	"fmt"
	"sync"
)

func listAndArray() {

	//1。数组
	var team [3]string
	team[0] = "Tom"
	team[1] = "Jack"
	team[2] = "Mary"
	fmt.Println(team)

	for k, v := range team {
		fmt.Println(k,v)
	}




	//2。切片
	var a = []int{1,2,3}
	fmt.Println(a,a[1:2])
	// 切片复制
	const elementCount = 100
	srcData := make([]int,elementCount)
	for i := 0; i < elementCount; i++ {
		srcData[i] = i
	}
	refData := srcData
	copyData := make([]int,elementCount)
	copy(copyData,srcData)
	srcData[0] = 999
	fmt.Println(refData[0])
	fmt.Println(copyData[0],copyData[elementCount-1])
	copy(copyData,srcData[4:6])
	for i := 0; i < 5; i++ {
		fmt.Println(copyData[i])
	}






	//3。映射map
	//map采用hash散列表实现的
	scene := make(map[string]int)
	scene["route"] = 66
	scene["brazil"] = 4
	scene["china"] = 960
	fmt.Println(scene)
	for k, v := range scene {
		fmt.Println(k,v)
	}

	//并发环境下使用map
	/*m := make(map[int]int)
	go func() {
		for {
			m[1] = 1
		}
	}()

	go func() {
		_ = m[1]
	}()

	for {

	}*/

	var score sync.Map
	score.Store("greece",70)
	score.Store("london",100)
	score.Store("egypt",200)
	fmt.Println(score)
	data,ok := score.Load("london")
	if !ok {
		fmt.Println(ok)
	}
	fmt.Println(data)
	score.Delete("london")
	score.Range(func(k, v interface{}) bool {
		fmt.Println(k,v)
		return true
	})






	//4。列表list
	//列表是一个非连续存储的容器，由多个节点组成，节点之间通过一些变量记录彼此之间的关系。
	//1.初始化列表，New和声明
	l1 := list.New()
	var l2 list.List
	fmt.Println(l1,l2)
	//2.列表中插入元素
	l1.PushBack("first")
	l1.PushFront(65)
	fmt.Println(l1)
	//3.列表中删除元素
	element := l1.PushBack("two")
	l1.InsertAfter("three",element)
	l1.InsertBefore("one",element)
	l1.Remove(element)
	fmt.Println(l1)

	for i := l1.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}



}
