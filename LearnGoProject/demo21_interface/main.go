package main

import (
	"fmt"
	"go-Demos/LearnGoProject/demo21_interface/model"
	"sort"
	"sync"
	"time"
)

/*
go  interface 引用类型 实现程序的高内聚 低耦合 (多态的体现)
1. 接口的方法是没有方法体的, 接口的方法声明, 不需要func关键字
2. go的接口不需要显示实现, 只要有一个数据结构, 包含接口的(所有)方法, 那么它就实现了接口.
3. 接口只能声明方法, 不能有变量
4. 接口本身不能创建实例, 但是可以指向一个实现了该接口的变量
5. 一个数据结构可以同时实现多个接口
6. 数据结构在实现接口时, 内部实现的方法只要有一个是带入的自己是指针 那么在使用该数据结构作为该接口变量时, 也必须使用指针类型.
   如AObject, 只有一个MethodA使用了带入自己使用了指针, 那么使用var aInterface model.AInterface = &aObject和aObjectPerform.Perform(&aObject)也必须为指针类型
7. 接口也能继承其他接口, 但是继承体系中, 如果有父接口中有相同的方法, 则会编译不通过 ⚠️⚠️⚠️⚠️此条规则在go1.14版本以后无效 参考链接:https://learnku.com/go/t/40984
   此后的版本允许两个多个接口有相同的方法(试了很多个版本, 1.14以下的版本都有此规则, 后面的取消了)
8. 任何数据结构都可以实现接口, 和给数据结构添加方法一个原理. 系统数据结构的添加, 需要先自己定义一个新的类型
9. 接口中不能有变量
10. 如果一个接口继承自其他接口, 那么实现接口也必须实现父接口的方法
11. interface{} 空接口, 没有任何方法, 所以所有的数据结构都实现了空接口

接口和继承的关系:
1. 接口是对继承的一种补充. 继承者完整拥有父类的特征的同时, 还可以使用接口对自身进行拓展
2. 继承是is - a的关系.  而接口只需要满足like - a的关系
3. 继承的主要价值在于代码的复用性和可维护性. 只需要在父类操作, 子类即可拥有相同的特征
3. 接口的主要价值在于设计, 使用更好的规范让自定义的各种类型实现各自(相同)的功能
*/

func main() {
	var phone *model.Phone = &model.Phone{}
	var camera *model.Camera = &model.Camera{}

	var phone1 model.Phone = model.Phone{}
	var camera1 model.Camera = model.Camera{}

	var computer *model.Computer = &model.Computer{}
	computer.UsedUsb(phone)
	computer.UsedUsb(camera)
	fmt.Println("-------")
	computer.UsedUsb(&phone1)  // 此处必须传入结构体地址. 接口为引用类型
	computer.UsedUsb(&camera1)


	//aobject实现了 AInterface. 但是Ainterface又继承自接口AAAInterface和AAInterface. 所以它三个结构体的方法都能调用
	var aObject model.AObject = model.AObject{}
	var aObjectPtr *model.AObject = &aObject
	var aaaInterface model.AAAInterface = aObject
	var aaInterface model.AAInterface = model.AAInterface(aObject)
	var aInterface model.AInterface = &aObject
	//直接使用*aInterface, 编译不通过

	var _aInterface model.AInterface  //interface默认为引用类型, 没有初始化, 输出nil
	fmt.Printf("aaaInterface = %v, aaInterface = %v, aInterface = %v\n", aaaInterface, aaInterface, aInterface)
	fmt.Printf("aInterface = %v, aobject = %v\n", aInterface, aObject)
	fmt.Printf("_aInterface = %v\n", _aInterface)

	aaaInterface.MethodAAA()
	fmt.Println()
	aaInterface.MethodAA()
	fmt.Println()
	aInterface.MethodA()
	aInterface.MethodAA()
	aInterface.MethodAAA()


	var aObjectPerform model.AObjectPerform = model.AObjectPerform{}
	aObjectPerform.Perform(&aObject)
	aObjectPerform.Perform(aObjectPtr)

	var st interface{} = 11
	var stt interface{} = aObjectPtr
	var sttt interface{} = aaaInterface
	fmt.Printf("st = %v, stt = %v, sttt = %v\n", st, stt, sttt)

	//sort.Ints 内部使用的是sort.Sort方法. 该方法传入一个接口, 接口定制好排序规则
	var slice []int = []int{1, 23, 2, 3232, 23, 44, 4}
	fmt.Printf("slice sort before:%v\n", slice)
	sort.Ints(slice)
	fmt.Printf("slice sort after:%v\n", slice)


	var slicePersons model.SlicePerson = model.SlicePerson{
		model.Person{"张三", 22, 170},
		model.Person{"李四", 20, 160},
		model.Person{"王五", 27, 180},
		model.Person{"陈六", 19, 165},
	}


	//sort.Sort 带入的一个实现Interface接口的集合 如sort.IntSlice sort.Float64Slice sort.StringSlice. 内部自行实现Interface接口 此接口名字就叫Interface(大写)
	fmt.Printf("slicePersons sort before = %v\n", slicePersons)
	sort.Sort(slicePersons)  //默认按年龄升序
	fmt.Printf("slicePersons sort by age ascending = %v\n", slicePersons)

	model.SortPersonTypeValue = model.KSortPersonType_Height
	sort.Sort(slicePersons)
	fmt.Printf("slicePersons sort by height ascending = %v\n", slicePersons)

	model.SortPersonTypeValue = model.KSortPersonType_Age
	model.SortPersonRuleValue = model.KSortPersonRule_Descending
	sort.Sort(slicePersons)
	fmt.Printf("slicePersons sort by age descending = %v\n", slicePersons)



	obj := model.AAAObject{AAA: "111"}
	var i interface{} = obj

	//todo  obj = 222, i和obj2=111
	obj.AAA =" 222"
	obj2 := i.(model.AAAObject)
	fmt.Println(obj, &obj, &obj2)
	fmt.Println(i)

}


type AInterface interface {
	test01()
	test02()
}
type BInterface interface {
	test01()
	test03()
}
type CInterface interface {
	AInterface  //1.14以后的版本允许同时父接口存在相同的方法
	BInterface
}
