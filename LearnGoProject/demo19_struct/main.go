package main

import (
	"encoding/json"
	"fmt"
	"go-Demos/LearnGoProject/demo19_struct/model"
	"time"
)

//golang 面对对象  它没有其它语言的class, 而是使用struct来实现.
//去掉了传统语言的继承.重载.构造函数.析构函数.this等. 它只是用其他的方式来实现类似的功能


//结构体  值类型
/*
1. 结构体的属性声明不需要var 需指定类型.
2. 取指针类型的字段时, 需注意自己带上*.  注意区分加*的位置, 如"*cat4.size" 和 "*(cat4.size)"
3. 结构体的内存布局是连续性的, 有时候我们可以直接推算出地址, 使用地址直接修改会更快
4. 结构体是用户单独定义的类型, 和其他类型转换时, 需要有完全相同的字段(字段名, 字段个数, 字段类型)
5. 结构体的每个字段可以设置一个tag 该tag可以通过反射机制获取. 常用场景:(序列化/反序列化).
	1. 使用tag时要注意:
		1. 字段名如果是小写, json包无法访问字段, 解析不出数据, 也不会产生error
		2. 如果是tag标记为json的格式错误, 也无法解析出正确的字段. 会默认struct自己的字段名
6. 结构题可增加方法, 为结构体增加一些扩展
	1. 方法的调用机制和原理:
		1. 方法的调用和函数基本一致, 不一样的地方. 方法在调用时, 会传入自己(类似oc的实例方法, 默认带self).
		2. 方的调用带入的自己, 要区分是带入的引用类型还是值类型.  引用类型带入的是地址, 修改影响到其他地方. 值类型是直接拷贝数据, 可能影响性能
		3. go中的方法是指定数据类型的, 所以任何类型都可以添加方法. 类似oc的category扩展(系统类型无法直接扩展, 需要重定义一个type, 再进行扩展)
		4. go中可对任何类型实现String()方法, 返回该值得描述. 比如Print的时候, 会调用它  (类似oc的description)
7. 方法和函数的区别: (通用函数. 类型方法)
	1. 参数列表不一样, 方法默认会带入自己的值或指针. 然后才是实参
	2. 调用方式不一样: 方法使用实例调用, 函数直接调用
	3. 两者的实参是什么类型, 都必须传入什么类型. (不可像方法前面的实例一样, 带入值和指针都可以)
8. 构建私有的结构体可通过工厂模式构建. 私有的结构体变量, 可使用结构体方法访问(自己构建getMethod setMethod)
9. 同一个包下面的私有属性, 可以相互访问 如: student可以访问student2
10. 结构体字段名可以是基本数据类型, 比如int float64, 单个时不需要指定字段名, 多个时需要指定字段名
11. 结构体声明为引用类型时, 必须先初始化, 否则无法使用. 如new() &structName{}
*/
type Cat struct {
	name string
	age int
	color string
	size *int
	//此处内部的结构体在Cat实例声明时就已经创建好了, 不需要像map slice需要额外的make
	extra struct {
		aaa string
	}
}

type AA struct {
	num int
	value string
}

type BB struct {
	num int
	value string
}

type BBB struct {
	num int
}

type Person struct {
	Name string `json:"name"`
	Age int `json:"age"`
	Height int `json:"我艹"`
}

//此处调用时, p是值传递
func (p Person) getBirthYear(age int) int {
	p.Age = 55
	date := time.Now()
	return date.Year() - age
}
//此处调用时, p是引用传递
func (p *Person) getBirthYearForPtr() int  {
	p.Age = 66
	date := time.Now()
	return date.Year() - p.Age
}

func (p Person) String() string {
 	return fmt.Sprintf("Namg:%v, Age:%v", p.Name, p.Age)
}

func (p Person) getBirthYearArgs(pp Person) int {
	return time.Now().Year() - pp.Age
}
func (p Person) getBirthYearArgsForPtr(pp *Person) int {
	return time.Now().Year() - pp.Age
}

func PrintPerson (p Person) {
	fmt.Println("Person信息:%v", p)
}
func PrintPersonPtr(p *Person) {
	fmt.Println("Person信息:%v", *p)
}

type integer int
func (i integer) fbn() []int {
	value := int(i)
	arr := make([]int, value)
	if value >= 1 {
		arr[0] = 1
	}
	if value >= 2 {
		arr[1] = 1
	}
	if value >= 3 {
		for i := 2; i < value; i++ {
			arr[i] = arr[i - 1] + arr[i - 2]
		}
	}
	return arr
}

//integer类型得说明
func (i integer) String() string {
	return fmt.Sprintf("这是integer类型得说明:%d", i)
}

func main() {
	//结构体初始化
	//声明变量 直接使用
	var cat1 Cat
	cat1.name = "锅锅"
	fmt.Printf("cat1.extra = %p\n", &(cat1.extra))
	cat1.extra = struct{ aaa string }{aaa: "11"}
	fmt.Printf("cat1.extra = %p\n", &(cat1.extra))
	cat1.extra.aaa = "我 是 锅锅"
	fmt.Printf("cat1.extra = %p\n", &(cat1.extra))
	fmt.Printf("cat1 = %v\n", cat1)

	//直接声明赋值厨初始值
	cat3 := Cat{} //要么不填参数 要么就全部填 要么像map一样, 指定属性填写
	cat2 := Cat{"嘿嘿", 11, "", &cat1.age, struct{ aaa string }{aaa: "嘿嘿傻傻的猫"}}
	cat5 := Cat{
		name: "长猫咪",
		size: &cat1.age,
	}
	fmt.Printf("cat3 = %v\n", cat3)
	fmt.Printf("cat2 = %v\n", cat2)
	fmt.Printf("cat5 = %v\n", cat5)

	//使用new  此处返回的是一个结构体指针. 怎么区分? 看输出是否带&
	cat := new(Cat)
	cat.name = "花花"
	cat.age = 1
	cat.color = "yellow"
	cat.extra = struct{aaa string}{aaa: "我是一只小花猫"}
	fmt.Printf("cat = %v\n", cat)

	//使用取地址符进行初始化
	var cat4 *Cat = &Cat{} //同样返回的是结构体指针
	cat4.age = 33
	cat4.name = "花小花"  //理论上cat4是一个指针, 所以访问name应该使用下一行的方式.  但是这里golang的作者为了使用方便. 不带*访问也是可以的.(会在底层自动带上*)
	(*cat4).name = "花大花"
	cat4.size = &cat4.age //此处也要注意. 如果某一个字段也是指针, 那么取值必须带上*. 否则是打印指针本身的值
	fmt.Printf(" cat4.size: %v\n", cat4.size)  //地址
	fmt.Printf(" *cat4.size: %v\n", *cat4.size) //值.  因为.的优先级更高, 所以这里取的是值
	fmt.Printf(" (*cat4).size: %v\n", (*cat4).size) //地址.
	fmt.Printf(" *(cat4.size): %v\n", *(cat4.size)) //值
	fmt.Printf("cat4 = %v %p\n", cat4)
	fmt.Printf("cat4 = %v\n", *cat4)



	aa := new(AA)
	aa.value = "AA"
	bb := new(BB)
	bb.num = 22
	bb.value = "BB"
	//bbb := &BBB{333}

	//此处加*强转是因为aa bb都是引用类型
	*aa = AA(*bb) //相同 (字段名称 字段类型 字段个数) 可以直接转
	//aa == AA(bb) //值类型强转
	fmt.Printf("aa = %v\nbb = %v\n", aa, bb)

	//两者 (字段名称 字段类型 字段个数)有一个条件不相同. 都无法强转
	//*aa = BBB(*bbb)
	//*bbb = BBB(*aa)

	var person Person = Person{"wang大爷", 23, 187}
	//结构体字段必须大写开头, 因为字段小写在json包中无法访问. 那么序列化为无效的bytes.(不会有error).
	//tag必须遵循`json:"字段名"`, 否则解析失败会以大写的字段名解析
	bytes, err := json.Marshal(person)
	if err != nil {
		fmt.Printf("person convert json err: %v\n", err)
	}
	strJson := string(bytes)
	fmt.Println(bytes)
	fmt.Printf("person = %v\njsonStr = %v\n", person, strJson)

	var birthYear = person.getBirthYear(person.Age)
	fmt.Println("出生年:", birthYear, "年龄:", person.Age)

	var birthYear2 = person.getBirthYearForPtr() //参数是指针, 也可以直接调用.  内部会带入person的指针
	var birthYear3 = (&person).getBirthYearForPtr()  //参数是指针, 使用传入指针类型调用
	fmt.Println("出生年:", birthYear2, "年龄:", person.Age)
	fmt.Println("出生年:", birthYear3, "年龄:", person.Age)


	//获取num的fnb值
	var num integer = 12
	fmt.Printf("num的斐波那契值:%v\n", num.fbn())
	fmt.Println(num)

	PrintPerson(person)
	PrintPersonPtr(&person)

	//实参是什么类型 就必须带入什么类型
	//person.getBirthYearArgsForPtr(person)
	person.getBirthYearArgsForPtr(&person)

	//使用工厂方法构建私有的struct
	var Stu model.Student = model.Student{
		Name: "公有Student",
		Score: 11,
	}
	var stu = model.StudentInstance("私有student", 22)
	fmt.Printf("Stu = %v, stu = %v\n", Stu, stu)
	fmt.Printf("通过方法访问变量:name = %v\n", stu.GetName())

}
