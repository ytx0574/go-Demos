package model

import "fmt"

type Customer struct {
	Id int
	Name string
	Gender string
	Age int
	Phone string
	Email string
}

func (self *Customer)String() string {
	return fmt.Sprint("编号:%v, 姓名:%v, 性别:%v, 年龄:%v, 手机号码:$v, 邮箱:%v\n", self.Id, self.Name, self.Gender, self.Age, self.Phone, self.Email)
}

func NewCustomer(name string, gender string, age int, phone string, email string) Customer {
	return Customer{
		Name : name,
		Gender : gender,
		Age : age,
		Phone : phone,
		Email : email,
	}
}

