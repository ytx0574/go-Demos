package main

import (
	"C"
	"go-Demos/LearnGoProject/demo23_CustomerManager/model"
	"go-Demos/LearnGoProject/demo23_CustomerManager/view"
	"go-Demos/LearnGoProject/demo23_CustomerManager/viewcontroller"
)

/*
#cgo CFLAGS: -x -objective-c
*/

/*
客户管理系统
*/
func main() {

	var vc = viewcontroller.CustomerViewController{}
	vc.Add(model.NewCustomer("张三", "男", 11, "111", "qqq"))
	vc.Add(model.NewCustomer("李斯特", "男", 22, "222", "www"))
	vc.Add(model.NewCustomer("墨然", "女", 33, "333", "eee"))

	var view = view.CustomView{
		VC: &vc,
	}
	view.ShowMenu()
}