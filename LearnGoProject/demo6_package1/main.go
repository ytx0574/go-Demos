package main

import (
	"go-Demos/LearnGoProject/demo6_package1/j_utils"
	ioPackage "fmt"
)

/*
关于包的说明:
	1. go使用文件夹来区分包的, 比如utils文件夹下面有个utils.go文件
	2. 需要外部使用的方法和变量必须大写开头
	3. 同一个文件夹下面, 不能打两个包, 比如(package a, packag a1)
	4. 默认情况下, 文件夹名称和包名一样. 如不一样的情况下, 导入文件夹名称(包), 使用里面的包名调用函数或变量. 包里面的文件名可以任意取
	5. 引入的包可以取一个别名  在引入时, 前面加一个别名, 如(ioPackake "fmt"), 使用"_"则不加载此包. 如引入时使用了别名, 则使用时也必须用别名
*/
func main () {
	utils_demo6.Print_string_replaced()
	ioPackage.Printf("%v", utils_demo6.ConstString2)

	//utils.Print_string()
	//fmt.Printf("j_utils.ConstString = %v\n", utils.ConstString)
}