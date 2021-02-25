package main

import (
	"encoding/json"
	"fmt"
)

//结构体中字段大写开头表示可公开访问，小写表示私有（仅在定义当前结构体的包中可访问）。

type person struct {
	name,city string
	age int8
}

type cat struct {
	string
	int
}

type Address struct {
	Province string
	City string
}

type Address2 struct {
	Province string
	City string
	CreateTime string
}

type Email struct {
	Account string
	CreateTime string
}

type User3 struct {
	Name string
	Gender string
	Address2
	Email
}

type User struct {
	Name string
	Gender string
	Address Address
}

type User2 struct {
	Name,Gender string
	Address
}

type Animal struct {
	name string
}

func (a *Animal) move() {
	fmt.Printf("%s 会动！\n",a.name)
}

type Dog1 struct {
	Feet int
	*Animal// 通过嵌套匿名结构体实现继承
}

func (d *Dog1) wang() {
	fmt.Printf("%s 会汪汪汪～\n",d.name)
}

type Student struct {
	ID int
	Gender string
	Name string
}

type Class struct {
	Title string
	Students []*Student
}

type Student1 struct {
	ID int `json:"id"`//通过指定tag实现json序列化该字段时的key
	Gender string//json序列化是默认使用字段名作为key
	name string//私有不能被json包访问
}

type Person1 struct {
	name string
	age int8
	dream []string
}

func (p *Person1) SetDream(dreams []string) {
	//p.dream = dreams
	p.dream = make([]string,len(dreams))
	// 因为切片是引用类型
	copy(p.dream,dreams)
}

type MyInt int
func StructFun() {

	var p1 person
	p1.age = 18
	p1.name = "沙河"
	p1.city = "北京"

	fmt.Printf("p1=%v\n",p1)
	fmt.Printf("p1=%#v\n",p1)


	// 匿名结构体
	var user struct{
		Name string
		Age int
	}
	user.Name = "王子"
	user.Age = 20
	fmt.Printf("%#v\n",user)



	// 指针类型结构体
	//var p2 = new(person)
	var p2 = &person{}
	p2.city = "上海"
	p2.name = "王子"
	p2.age = 20
	fmt.Printf("%T\n",p2)
	fmt.Printf("%#v\n",p2)



	// 结构体初始化
	p3 := person{
		age: 13,
		city: "北京",
	}
	fmt.Printf("%#v\n",p3)
	p4 := person{
		"王武",
		"上海",
		12,
	}
	fmt.Printf("p4=%#v\n",p4)
	/*
	使用这种格式初始化时，需要注意：
		必须初始化结构体的所有字段。
		初始值的填充顺序必须与字段在结构体中的声明顺序一致。
		该方式不能和键值初始化方式混用。
	 */

	type test struct {
		a int8
		b int8
		c int8
		d int8
	}

	n := test{1,2,3,4}
	fmt.Printf("n.a %p \n",&n.a)
	fmt.Printf("n.b %p \n",&n.b)
	fmt.Printf("n.c %p \n",&n.c)
	fmt.Printf("n.d %p \n",&n.d)

	/*type student struct {
		name string
		age int
	}
	m := make(map[string]*student)

	stus := []student{
		{name: "小王子",age: 18},
		{name: "娜扎",age: 23},
		{name: "大王八",age: 9000},
	}
	for _, stu := range stus {
		m[stu.name] = &stu
	}

	for k, v := range m {
		fmt.Println(k,"=>",v.name)
	}
	小王子 => 大王八
	娜扎 => 大王八
	大王八 => 大王八
	*/

	p5 := newPerson("李明","杭州",45)
	fmt.Printf("p5=%#v\n",p5)

	p5.Dream()

	p5.SetAge(12)
	fmt.Printf("p5=%#v\n",p5)
	p5.SetAge2(13)
	fmt.Printf("p5=%#v\n",p5)


	var m1 MyInt
	m1.SayHello()
	m1 = 100
	fmt.Printf("%#v %T\n",m1,m1)


	// 结构体的匿名字段
	c := cat{
		"小黄",
		12,
	}
	fmt.Printf("%#v\n",c)


	// 结构体嵌套
	user1 := User{
		Name: "小王子",
		Gender: "男",
		Address: Address{
			Province: "山东",
			City: "威海",
		},
	}
	fmt.Printf("%#v\n",user1)


	var user2 User2
	user2.Name = "小王子"
	user2.Gender = "男"
	user2.Address.Province = "山东"// 匿名字段默认使用类型名作为字段名
	user2.City = "威海"// 匿名字段可以省略
	fmt.Printf("%#v\n",user2)

	//嵌套结构体字段名冲突
	var user3 User3
	user3.Name = "娜扎"
	user3.Gender = "女"
	//user3.CreateTime = "2019"//Ambiguous reference 'CreateTime'
	// 分别制定每个结构体里面字段的值
	user3.Address2.CreateTime = "2019"
	user3.Email.CreateTime = "2020"
	fmt.Printf("%#v\n",user3)



	// 结构体继承
	d := &Dog1{
		Feet: 4,
		Animal: &Animal{
			name: "乐乐",
		},
	}
	fmt.Printf("%#v\n",d)
	d.move()
	d.wang()



	//结构体与json
	c1 := &Class{
		Title: "101",
		Students: make([]*Student,0,200),
	}
	for i := 0; i < 10; i++ {
		stu := &Student{
			Name: fmt.Sprintf("stu%02d",i),
			Gender: "男",
			ID: i,
		}
		c1.Students = append(c1.Students,stu)
	}

	data,err := json.Marshal(c1)
	if err != nil {
		fmt.Println("json marshal failed")
		return
	}
	fmt.Printf("json:%s\n",data)

	// json反序列化
	stu := `{"Title":"101","Students":[{"ID":0,"Gender":"男","Name":"stu00"},{"ID":1,"Gender":"男","Name":"stu01"},{"ID":2,"Gender":"男","Name":"stu02"},{"ID":3,"Gender":"男","Name":"stu03"},{"ID":4,"Gender":"男","Name":"stu04"},{"ID":5,"Gender":"男","Name":"stu05"},{"ID":6,"Gender":"男","Name":"stu06"},{"ID":7,"Gender":"男","Name":"stu07"},{"ID":8,"Gender":"男","Name":"stu08"},{"ID":9,"Gender":"男","Name":"stu09"}]}`
	c2 := &Class{}
	err = json.Unmarshal([]byte(stu),c2)
	if err != nil {
		fmt.Println("json unmarshal failed")
		return
	}
	fmt.Printf("%#v\n",c2)


	//结构体标签Tag
	s1 := Student1{
		ID: 1,
		Gender: "男",
		name: "娜扎",
	}
	fmt.Printf("%#v\n",s1)

	data1,err1 := json.Marshal(s1)
	if err1 != nil {
		fmt.Println("json marshal failed")
		return
	}
	fmt.Printf("json:%s\n",data1)



	//补充
	p6 := Person1{
		name: "小王子",
		age: 18,
	}
	data2 := []string{"吃饭","睡觉","打豆豆"}
	p6.SetDream(data2)
	data2[1] = "不睡觉"// 只想修改data2的值，并不想修改p6的值
	fmt.Printf("%#v",p6)
}

// 构造函数
//Go语言的结构体没有构造函数，我们可以自己实现。 例如，下方的代码就实现了一个person的构造函数。
//因为struct是值类型，如果结构体比较复杂的话，值拷贝性能开销会比较大，所以该构造函数返回的是结构体指针类型。
func newPerson(name string,city string,age int8) *person {
	return &person{
		name: name,
		city: city,
		age: age,
	}
}

//方法
//方法与函数的区别是，函数不属于任何类型，方法属于特定的类型。
func (p person) Dream() {
	fmt.Printf("%s的梦想是学好Go语言！\n",p.name)
}

// 指针接收者、值接收者
func (p *person) SetAge(newAge int8) {
	p.age = newAge
}
func (p person) SetAge2(newAge int8) {
	p.age = newAge
}
/*
什么时候应该使用指针类型接收者:
	需要修改接收者中的值
	接收者是拷贝代价比较大的大对象
	保证一致性，如果有某个方法使用了指针接收者，那么其他的方法也应该使用指针接收者。
 */


//注意事项： 非本地类型不能定义方法，也就是说我们不能给别的包的类型定义方法。
func (m MyInt) SayHello() {
	fmt.Println("Hello,我是int")
}