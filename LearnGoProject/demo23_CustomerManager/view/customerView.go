package view

import (
	"fmt"
	"go-Demos/LearnGoProject/demo23_CustomerManager/model"
	"go-Demos/LearnGoProject/demo23_CustomerManager/viewcontroller"
	"os"
	"strings"
)

type CustomView struct {
	VC *viewcontroller.CustomerViewController
	funcType int
}

func (self *CustomView) ShowMenu() {

	fmt.Printf("--------客户信息管理软件--------\n")
	fmt.Printf("	         1添加客户\n")
	fmt.Printf("	         2修改客户\n")
	fmt.Printf("	         3删除客户\n")
	fmt.Printf("	         4客户列表\n")
	fmt.Printf("	         5退出程序\n")
	fmt.Println("请选择功能(1-5):")
	status, err := fmt.Scanf("%d", &self.funcType)
	fmt.Printf("status = %v, err = %v, funcType = %v\n", status, err, self.funcType)

	switch self.funcType {
	case 1:
		self.add()
	case 2:
		self.update()
	case 3:
		self.delete()
	case 4:
		self.showList()
	case 5:
		self.exit()
	default:
		fmt.Printf("选择的功能无效, 请重试!!!")
		self.ShowMenu()
	}
}

func (self *CustomView)add() {
	var Name string
	var Gender string
	var Age int
	var Phone string
	var Email string
	fmt.Println("请输入姓名:")
	status, err := fmt.Scanf("%v\n", &Name)
	fmt.Printf("status = %v, err = %v, Name = %v\n", status, err, Name)

	for  {
		if Gender == "" {
			fmt.Println("请输入性别:")
			fmt.Scanf("%v\n", &Gender)
		} else if Gender != "男" && Gender != "女" {
			fmt.Println("性别输入错误, 性别只能为男或女")
			fmt.Scanf("%v\n", &Gender)
		} else {
			break
		}
	}
	fmt.Println("请输入年龄:")
	fmt.Scanf("%v\n", &Age)

	fmt.Println("请输入手机号:")
	fmt.Scanf("%v\n", &Phone)

	fmt.Println("请输入邮箱:")
	fmt.Scanf("%v\n", &Email)

	customer := model.NewCustomer(Name, Gender, Age, Phone, Email)
	if self.VC.Add(customer) {
		fmt.Printf("用户添加成功:%v\n", customer)
		self.ShowMenu()
	}else {
		fmt.Printf("用户失败成功:%v\n", customer)
	}
}

func (self *CustomView)delete() {
	var id int
	fmt.Println("请输入要删除的用户编号:")
	fmt.Scanf("%v\n", &id)

	flag, err := self.VC.DeleteById(id);
	if flag && err == nil {
		fmt.Printf("删除成功\n")
	}else {
		fmt.Printf("删除失败%v\n", err)
	}
	self.ShowMenu()
}

func (self *CustomView)update() {
	var id int
	for  {
		fmt.Println("请输入要修改的用户编号:")
		fmt.Scanf("%v\n", &id)

		flag, _ := self.VC.FindById(id)
		if flag == false {
			fmt.Printf("您输入的用户id不存在, 请重新输入\n")
		}else {
			break
		}
	}

	customer := self.InputNewCustomer()
	self.VC.UpdateById(id, customer)

	self.ShowMenu()
}

func (self *CustomView)showList() {
	fmt.Println("编号\t名字\t性别\t年龄\t手机\t邮箱\t")
	for _, v := range self.VC.List() {
		fmt.Printf("%v\t%v\t%v\t%v\t%v\t%v\t\n", v.Id, v.Name, v.Gender, v.Age, v.Phone, v.Email)
	}

	self.ShowMenu()
}

func (self *CustomView)exit() {
	var value string

	for  {
		fmt.Println("是否退出程序? (Y/N)")
		fmt.Scanf("%v\n", &value)
		value = strings.ToLower(value)
		fmt.Printf("value = %q\n", value)


		if value != "y" && value != "n" {
			fmt.Printf("输入不正确, 请重新输入\n")
		}else {
			break
		}
	}

	if value == "y" {
		os.Exit(0)
	}else {
		self.ShowMenu()
	}
}

func (self *CustomView) InputNewCustomer() model.Customer {
	var Name string
	var Gender string
	var Age int
	var Phone string
	var Email string
	fmt.Println("请输入姓名:")
	status, err := fmt.Scanf("%v\n", &Name)
	fmt.Printf("status = %v, err = %v, Name = %v\n", status, err, Name)

	for  {
		if Gender == "" {
			fmt.Println("请输入性别:")
			fmt.Scanf("%v\n", &Gender)
		} else if Gender != "男" && Gender != "女" {
			fmt.Println("性别输入错误, 性别只能为男或女")
			fmt.Scanf("%v\n", &Gender)
		} else {
			break
		}
	}
	fmt.Println("请输入年龄:")
	fmt.Scanf("%v\n", &Age)

	fmt.Println("请输入手机号:")
	fmt.Scanf("%v\n", &Phone)

	fmt.Println("请输入邮箱:")
	fmt.Scanf("%v\n", &Email)

	return model.NewCustomer(Name, Gender, Age, Phone, Email)
}
