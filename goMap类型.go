package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
	"time"
)

func Map() {
	//map是一种无序的基于key-value的数据结构，map是引用类型，必须初始化才能使用

	// map的定义格式	map[keyType]valueType	keyType是键的类型valueType是键对应值的类型
	// map类型的变量默认初始值是nil，需要使用make来分配内存	make(map[keyType]valueType,[cap])	其中cap表示map的容量，该参数虽然不是必须的，但是我们应该在初始化map的时候就为其指定一个合适的容量。
	scoreMap := make(map[string]int,8)
	scoreMap["小明"] = 80
	scoreMap["张三"] = 100
	fmt.Println(scoreMap)
	fmt.Println(scoreMap["张三"])
	fmt.Printf("%T\n",scoreMap)

	userInfo := map[string]string{
		"username": "张三",
		"password": "123456",
	}
	fmt.Println(userInfo)





	// map的常见函数
	// 1。判断某个键是否存在，格式：value, ok := map[key]
	v, ok := scoreMap["张三"]
	// 如果存在这个键ok就是true，v是对应的值，否则ok为false，v是对应值类型的零值
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("查无此人")
	}

	// 2。map遍历	forRange
	for key, value := range scoreMap {
		fmt.Println(key,"------",value)
	}
	for key := range scoreMap{
		fmt.Println(key)
	}
	// map遍历的时候，元素顺序和添加时候键值对的顺序无关

	// 3。按照指定顺序遍历map
	forLoopMap()


	// 4。delete删除键值对，格式：delete(map,key)	map指要删除的map，key指删除的键
	scoreMap1 := make(map[string]int)
	scoreMap1["莉丝"] = 100
	scoreMap1["王武"] = 50
	scoreMap1["赵六"] = 3
	delete(scoreMap1,"莉丝")
	for key, val := range scoreMap1 {
		fmt.Println(key,"-----",val)
	}


	// 5。map类型的切片
	var mapSlice = make([]map[string]string,3)
	for index, value := range mapSlice {
		fmt.Printf("index:%d,value:%v",index,value)
	}
	fmt.Println("after init")
	mapSlice[0] = make(map[string]string,10)
	mapSlice[0]["name"] = "小明"
	mapSlice[0]["age"] = "18"
	mapSlice[0]["address"] = "沙河"
	for index, value := range mapSlice {
		fmt.Printf("index:%d,value:%v",index,value)
	}

	// 6。切片类型的map
	var sliceMap = make(map[string][]string,3)
	fmt.Println(sliceMap)
	fmt.Println("after init")
	key := "中国"
	val,ok := sliceMap[key]
	if !ok {
		val = make([]string,0,2)
	}
	val = append(val,"上海","北京")
	sliceMap[key] = val
	fmt.Println(sliceMap)

	// 作业
	//写一个程序，统计一个字符串中每个单词出现的次数。比如：”how do you do”中how=1 do=2 you=1。
	mapHomework()

	result()

}

/*
首先初始化随机数种子，然后创建一个map，获取map的所有key组成一个切片，然后给切片排序，在通过遍历切片，分别获取map的key和value
 */
func forLoopMap() {
	// 初始化随机数种子
	rand.Seed(time.Now().UnixNano())

	var scoreMap = make(map[string]int,200)
	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("stu%2d",i)
		val := rand.Intn(100)
		scoreMap[key] = val
	}

	var keys = make([]string,0,200)
	for key := range scoreMap {
		keys = append(keys,key)
	}

	sort.Strings(keys)

	for _, key := range keys {
		fmt.Println(key,scoreMap[key])
	}
}

func mapHomework() {
	str := "how do you do"
	m := make(map[string]int)
	wordSlice := strings.Split(str, " ")
	for _, value := range wordSlice {
		_, ok := m[value]
		fmt.Println(value,ok)
		if ok {
			m[value] = m[value] + 1
		} else {
			m[value] = 1
		}
	}
	fmt.Println(m)
}

func result() {
	type Map map[string][]int
	m := make(Map)
	s := []int{1, 2}
	s = append(s, 3)
	fmt.Printf("%+v\n", s)	//[1 2 3]
	m["q1mi"] = s
	s = append(s[:1], s[2:]...)
	fmt.Printf("%+v\n", s)	//[1 3]
	fmt.Printf("%+v\n", m["q1mi"])	//["q1mi"][1 2 3]
}