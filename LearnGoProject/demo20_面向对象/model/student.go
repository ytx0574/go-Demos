package model

import (
	"fmt"
)

type PseronHobby struct {
	TypeName string
}

type Student struct {
	Name string
	age int
}

//学生
func (stu *Student) Testing() {
	fmt.Printf("测验人名字:%v, 年龄:%v\n", stu.Name, stu.age)
}
func (stu *Student) setAge(age int) {
	stu.age = age
}

//大学生
type Graduate struct {
	Student  //匿名结构体
	Age int
	Hobby PseronHobby  //有名结构体
}
func (graduate *Graduate) SetAge(age int) {
	if age < 18 {
		graduate.Age = -1
	}else {
		graduate.Age = age
	}
}

//小学生
type Pupil struct {
	Books
	int
	//int int //编译不通过, 其实int int声明 = 上面的int声明
	//int string //同样编译不通过.  已经有一个变量名为int的字段
	*Student  //指针型匿名结构体  不可同时出现同一个匿名结构体 指针和值 指针型匿名结构体必须先初始化, 否则导致其他值类型无法正常调用(运行时错误)

	OtherBook *Books //有名指针型结构体不需要强制初始化.
	Extra *struct{
		Content string
	}
	Height *int
}
func (pupil *Pupil) SetAge(age int) {
	pupil.int = 12312
	if age < 6 || age > 12 {
		pupil.age = -1
	}else {
		pupil.age = age
	}
}

type Books struct {
	Name string
	page int
}