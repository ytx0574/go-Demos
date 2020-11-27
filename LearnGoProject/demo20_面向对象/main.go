package main

import (
	"fmt"
	"go-Demos/LearnGoProject/demo20_面向对象/model"
)

/*
go 封装 继承 多态
1. 封装
	1. 使用小写的名称 字段对不需要暴露的数据进行隐藏
	2. 提供一个工厂模式, 用于创建结构体. 类似其他语言的构造函数
	3. 提供方法对字段进行包装, 可添加业务逻辑, 规避不合理的取值 赋值
2. 继承:
	1. 使用嵌套的匿名结构体来实现继承, 可以直接使用匿名结构体大写/小写的字段和方法
	2. 匿名结构体的字段可以简化, 如果本访问本结构体没有的方法或字段, 可以直接点语法访问匿名结构体的方法和字段
	3. 如果本结构体有字段/方法, 点语法访问的时候, 优先访问自己的字段/方法. 没有得情况下才会访问匿名结构体的字段/方法(就近原则访问)
	4. 如果本结构体没有的方法, 点语法访问得是匿名结构体的方法. 此时如果方法中访问了两者同样的字段, 那么访问的是匿名结构体里面的字段
	5. 嵌入多个匿名结构体, 且他们有相同的字段和方法, 那么要通过指定匿名结构体访问. 否则编译不通过. 他们中不同的字段/方法, 还是采用就近原则访问
	6. 如果结构体嵌套了一个有名结构体, 相当于创建了一个结构体属性. 那么必须使用属性名访问
	7. 结构体字段名可以是基本数据类型, 比如int, 那么使用时直接使用基本数据类型访问(int 声明 = int int声明). 如果有多个相同的相同的基本数据类型, 则另外的必须指定一个字段名
	8. go可以使用多个匿名结构体实现多继承, 但是不建议这么使用
	9. 特殊注意点:
		1. 结构体中如果有指针型匿名结构体字段, 那么必须先初始化其指针型匿名结构体, 才能使用该结构体其他字段.而有名指针型结构体字段是否初始化, 不影响其他字段的使用.
		   比如下面的model.Pupil.Books如果不先初始化, 那么model.Pupil.age也无法使用. 运行时错误
		2. 默认的fmt.Print打印的内容会包含指针型匿名结构体的内容和其他基本数据结构内容. 其他有名指针型数据结构不会打印
		   比如下面的model.Pupil.Books会被直接打印出来, 而Extra和OtherBook和Height只会打印地址
*/

func main() {
	var account *model.Account = model.CreateAccount("张三", "999999")

	depositStatus, error := account.Deposit(1, "999999")
	fmt.Printf("存款状态:%t, 错误信息:%v\n", depositStatus, error)

	depositStatus, error = account.Deposit(101, "999999")
	fmt.Printf("存款状态:%t, 错误信息:%v\n", depositStatus, error)


	checkStatus, balance, error := account.Check("111")
	fmt.Printf("检查银行卡余额信息: 状态:%v, 余额:%v, 错误信息:%q\n", checkStatus, balance, error)

	//此处编译不通过, 会提示三个字段中没有新字段出现
	//checkStatus, balance, error := account.Check("999999")
	checkStatus, balance, error = account.Check("999999")
	fmt.Printf("检查银行卡余额信息: 状态:%t, 余额:%v, 错误信息:%v\n", checkStatus, balance, error)


	withdrawStatus, error := account.Withdraw(11, "999999")
	fmt.Printf("取现状态:%t, 错误信息:%v\n", withdrawStatus, error)

	checkStatus, balance, error = account.Check("999999")
	fmt.Printf("检查银行卡余额信息: 状态:%t, 余额:%v, 错误信息:%v\n", checkStatus, balance, error)


	//此处的age指向的是各自的值
	graduate := &model.Graduate{
		Age: 22,
		//Student : model.Student{"大学生", age:1}
		Student : model.Student{
			Name: "大学生",
		},
	}
	graduate.Testing()
	fmt.Printf("graduate = %v\n", graduate)

	var graduate1 *model.Graduate = new(model.Graduate)
	graduate1.Name = "大学生1"
	graduate.Age = 22
	graduate1.SetAge(18) //此处设置的值为Graduate自己的Age. Student的age还是0
	fmt.Printf("graduate1 = %v\n", graduate1)

	graduate1.Hobby.TypeName = "羽毛球"
	fmt.Printf("graduate1 = %v\n", graduate1)
	
	var pupil model.Pupil
	pupil.Student = &model.Student{
		Name: "我是小学生",
	}
	fmt.Printf("pupil.Student = %v\n", pupil.Student)
	//pupil.Name = "我是小学生" //此处编译不通过, 因为内部两个匿名函数都包含Name
	pupil.Student.Name = "我是小学生" //此处使用时, 必须保证Student初始化了. pupil.Student是指针
	(*pupil.Student).Name = "我是小学生"    //此处为上一行的标准写法. 内部匿名Student是指针类型, 所以先去除Student, 再设置其Name
	pupil.SetAge(1) //无效年龄. 内部设置为-1
	fmt.Printf("pupil = %v\n", pupil)

	pupil.SetAge(9)
	pupil.Books.Name = "语文书"
	fmt.Printf("~pupil= %v\n", pupil)

	pupil.OtherBook = new(model.Books)
	pupil.OtherBook.Name = "小学数学书"
	fmt.Printf("~~pupil= %v\n", pupil)

	pupil.Extra = &struct{ Content string }{Content: "我是一个extra"}
	pupil.Extra.Content = "我是重新赋值的Extra"
	fmt.Printf("~~~pupil= %v\n", pupil)

	pupil.Height = new(int)
	*pupil.Height = 22
	fmt.Printf("~~~~pupil= %v\n", pupil)

	fmt.Printf("pupil.OtherBook:%v\n", pupil.OtherBook)
	fmt.Printf("pupil.Extra:%v\n", pupil.Extra)
	fmt.Printf("pupil.Height:%v\n", *pupil.Height)
}