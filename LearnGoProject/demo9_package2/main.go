package main

import (
	ioPackage "fmt"
	"go-Demos/LearnGoProject/demo6_package1/j_utils"
	"go-Demos/LearnGoProject/demo9_package2/j_utils"
)
/*
go 给导入的包起别名, 使用时只能用别名调用
*/
func init () {
	ioPackage.Printf("demo9_package2 main init()\n")
}

func main () {
	utils_demo6.Print_string_replaced()
	ioPackage.Printf("%v\n", utils_demo6.ConstString2)

	utils_demo9.Print_string()
}