package main

import (
	"fmt"
	"go-Demos/LearnGoProject/demo21_interface/model"
)

/*
go  assertion 断言
1. 接口类型断言.  使用时, 必须确保左边必须是接口类型. 否则编译不通过
2. 类型断言使用时, 如只接收一个返回值, 当转换失败时, 会直接崩溃. 接收两个返回值时, 转换成功或失败都不会崩溃.
3. 类型断言内部的原理: 运行时自动选择调用哪个函数.  现在不懂, 后面再来搞 参考链接: https://zhuanlan.zhihu.com/p/136949285
	1. func assertI2I(inter *interfacetype, i iface) (r iface)
	2. func assertI2I2(inter *interfacetype, i iface) (r iface, b bool)
	3. 上面两个函数会调用getitab. 其中一个返回值时, gettitab的参数canfail为false, 会panic(运行时错误) 两个返回值时, 内部canfail为true. 不会panic
*/

func main() {
	var aObject model.AObject = model.AObject{"AObject"}
	var aInterface model.AInterface = &aObject

	aaObject := model.AAObject{"AAObject"}
	var aaInterface model.AAInterface = aaObject
	//aaInterface1 := aaObject
	//aaInterface1 = aaObject.(model.AAInterface)

	var aInterface1 = aInterface.(model.AInterface)
	fmt.Printf("aInterface1 = %v\n", aInterface1)

	//xxx.(type)  //判断变量实际类型, 固定写法, 只能配合switch语句使用
	//fmt.Printf("aInterface Type = %v\n", aInterface1.(type))

	//从接口转为原始类型. 此处因为aInterface指向的是aObject的地址. 所以转的时候需要自行转指针取值
	var aaObject1 model.AAObject = aaInterface.(model.AAObject)
	//编译不通过, 类型断言左边必须是显示的接口. 否则无法使用
	//var aaObject2 model.AAObject = aaInterface1.(model.AAObject)

	var aObject1 *model.AObject = aInterface.(*model.AObject)
	var aObject2 model.AObject = *(aInterface.(*model.AObject))

	fmt.Printf("aaObject1 = %v\n", aaObject1)
	fmt.Printf("aObject1 = %v\n", aObject1)
	fmt.Printf("aObject2 = %v\n", aObject2)


	//正常转换
	var aaObject2 model.AAObject = aaInterface.(model.AAObject)
	//非正常转换.
	var aObject3, convertToAObject3 = aaInterface.(model.AObject)
	//运行时错误, aaInterface无法转AObject. 内部代码实现时, 只有接收一个返回值直接崩溃 接口两个返回值时接着后面的代码运行
	//fmt.Println(aaInterface.(model.AObject))
	//var aObject4 = aaInterface.(model.AObject)

	var aaObject3, convertToAAObject3 = aaInterface.(model.AAObject)
	fmt.Printf("aaObject2 = %v\n", aaObject2)
	fmt.Printf("aObject3---- = %v\n", aObject3)
	fmt.Printf("aaObject3--- = %v\n", aaObject3)

	//aObject3 = model.AObject{}
	if convertToAObject3 {
		fmt.Printf("aObject3 success = %v\n", aObject3)
	}else {
		fmt.Printf("aObject3 error = %v\n", aObject3)
	}

	if convertToAAObject3 {
		fmt.Printf("aaObject3 success = %v\n", aaObject3)
	}else{
		fmt.Printf("aaObject3 fail = %v\n", aaObject3)
	}

	//直接两个返回值接收
	if aaObject4, ok := aaInterface.(model.AAInterface); ok {
		fmt.Printf("类型断言成功:aaObject4 = %v\n", aaObject4)
	}else {
		fmt.Printf("类型断言失败~~\n")
	}

	判断类型(33.33)
	判断类型([]int{})
	判断类型(aObject2)
	判断类型(aInterface)
}

func 判断类型(a interface{}) {
	switch a.(type) {
	case int:
		fmt.Printf("%v是一个int类型\n", a)
	case []int64:
		fmt.Printf("%v是一个切片类型\n", a)
	case string:
		fmt.Printf("%v是一个string类型\n", a)
	case int64, float64, float32:
		fmt.Printf("%v是一个数字类型\n", a)
	case model.AAObject:
		fmt.Printf("%v是一个model.AAObject类型\n", a)
	case *model.AAObject:
		fmt.Printf("%v是一个*model.AAObject类型\n", a)
	default:
		fmt.Printf("未判断到传入的类型%v\n", a)
	}
}